package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"unicode"

	"github.com/dave/jennifer/jen"
)

const (
	pubKey           = "ApiKeys.PublicKey"
	pvtKey           = "ApiKeys.PrivateKey"
	schemaDir        = "configs"
	resourceFilePath = "/cmd/resource/resource.go"
	CreateMethod     = "Create"
	ReadMethod       = "Read"
	UpdateMethod     = "Update"
	ListMethod       = "List"
	DeleteMethod     = "Delete"
	validationMethod = "validateModel"
	setupMethod      = "setup"

	createReqFields = "CreateRequiredFields"
	readReqFields   = "ReadRequiredFields"
	updateReqFields = "UpdateRequiredFields"
	deleteReqFields = "DeleteRequiredFields"
	listReqFields   = "ListRequiredFields"
	handler         = "github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	cloudformation  = "github.com/aws/aws-sdk-go/service/cloudformation"
	util            = "github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	validator       = "github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	progressEvent   = "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progress_event"
	log             = "github.com/sirupsen/logrus"
	atlas           = "go.mongodb.org/atlas/mongodbatlas"
	res             = "res"
)

func main() {
	resource := flag.String("res", "search-indexes", "adds generated client to the resource.go")
	schemaName := flag.String("schemaName", "indexes", "adds generated client to the resource.go")
	flag.Parse()

	// add validator.go
	resourceFileFullPath := fmt.Sprintf("%s%s", *resource, resourceFilePath)

	// add resource.go
	addResource(resourceFileFullPath, *schemaName)
}

func addResource(path, schemaName string) {
	f := jen.NewFile("resource")
	funcNames := []string{CreateMethod, ReadMethod, UpdateMethod, DeleteMethod, ListMethod}

	// Imports
	f.ImportName(handler, "handler")
	f.ImportName(cloudformation, "cloudformation")
	f.ImportName("errors", "errors")
	f.ImportName(util, "util")
	f.ImportName(validator, "validator")
	f.ImportName(progressEvent, "progress_events")
	f.ImportAlias(log, "log")

	// Required Fields mapping
	reqFields, err := readRequiredParameters(schemaName)
	if err != nil {
		fmt.Println(err)
		return
	}

	if reqFields != nil {
		f, err = addValidator(f, *reqFields)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Add setup method
	f.Func().Id(setupMethod).Params().
		Block(
			jen.Qual(util, "SetupLogger").Call(jen.Lit("mongodb-atlas-" + schemaName)),
		)
	f.Line()

	// Add handlers
	for _, fn := range funcNames {
		inputParameters, modelObj := generateInputParams(fn, reqFields)
		if len(inputParameters) > 0 {
			inputParameters = "Considerable params from currentModel: \n" + inputParameters
		}

		atlasFunc := " res , resModel, err := client." + capitalize(schemaName) + "." + fn + "(context.Background(),&mongodbatlas." + capitalize(schemaName) + "{\n" + modelObj + "})"

		f.Func().Id(fn).Params(
			// Input Parameters
			jen.Id("req").Qual(handler, "Request"),
			jen.Id("prevModel").Id("*Model"),
			jen.Id("currentModel").Id("*Model")).
			// Return types
			Id("(").Qual(handler, "ProgressEvent").Id(",").Id("error").Id(")").

			// Function starts
			Block(
				// log setup()
				jen.Id("setup").Params(),
				jen.Line(),
				// Debug log
				jen.Qual(log, "Debugf").Params(jen.Lit(fn+"() currentModel:%+v"), jen.Id("currentModel")),
				jen.Line(),

				// Validator
				jen.Comment("Validation"),
				jen.Id("modelValidation").Id(":=").Id(validationMethod).Call(jen.Id(fn+"RequiredFields"), jen.Id("currentModel")),
				jen.If(jen.Id("modelValidation").Op("!=").Id("nil")).Block(
					jen.Return(
						jen.Id("*modelValidation"),
						jen.Id("nil"),
					),
				),
				jen.Line(),

				// Atlas Client
				jen.Comment("Create atlas client"),
				jen.List(jen.Id("client"), jen.Err()).Op(":=").Qual(util, "CreateMongoDBClient").
					Call(jen.Id("*currentModel.ApiKeys.PublicKey"),
						jen.Id("*currentModel.ApiKeys.PrivateKey")),

				jen.If(jen.Id("err").Op("!=").Id("nil")).Block(
					jen.Qual(log, "Debugf").
						Call(jen.Lit(fn+" - error: %+v"), jen.Id("err")),

					jen.Return(
						jen.Qual(handler, "ProgressEvent").
							Values(jen.Dict{
								jen.Id("OperationStatus"):  jen.Qual(handler, "Failed"),
								jen.Id("Message"):          jen.Id("err.Error()"),
								jen.Id("HandlerErrorCode"): jen.Qual(cloudformation, "HandlerErrorCodeInvalidRequest"),
							}),
						jen.Id("nil"),
					),
				),

				jen.Id("var").Id(res).Id("*").Qual(atlas, "Response"),
				jen.Line(),

				// Pseudocode
				jen.Comment(inputParameters),
				jen.Comment(" // Pseudocode:\n"+atlasFunc+"\n\n"),
				jen.Line(),

				jen.If(jen.Id("err").Op("!=").Id("nil")).Block(
					jen.Qual(log, "Debugf").
						Call(jen.Lit(fn+" - error: %+v"), jen.Id("err")),

					jen.Return(
						jen.Qual(progressEvent, "GetFailedEventByResponse").
							Params(jen.Id("err.Error()"), jen.Id("res.Response")),
						jen.Id("nil"),
					),
				),

				// Log
				jen.Qual(log, "Debugf").
					Params(jen.Lit("Atlas Client %v"), jen.Id("client")),

				jen.Line(),
				// Progress Event Response
				jen.Comment("Response"),
				// Return statement
				jen.Return(jen.List(jen.Qual(handler, "ProgressEvent").
					Values(jen.Dict{
						jen.Id("OperationStatus"): jen.Qual(handler, "Success"),
						jen.Id("ResourceModel"):   jen.Id("currentModel"),
					}), jen.Id("nil"))),
			)
		f.Line()
	}
	fmt.Println("saving file at", path)
	err = f.Save(path)

	fmt.Println(err)
}

func generateInputParams(fn string, params *RequiredParams) (prms string, obj string) {
	switch fn {
	case CreateMethod:
		return generateParamString(params.CreateFields.InputParams)
	case ReadMethod:
		return generateParamString(params.ReadFields.RequiredParams)
	case UpdateMethod:
		return generateParamString(params.UpdateFields.InputParams)
	case DeleteMethod:
		return generateParamString(params.DeleteFields.RequiredParams)
	case ListMethod:
		return generateParamString(params.ListFields.InputParams)
	}
	return prms, obj
}

func generateParamString(fields []string) (params string, obj string) {
	if len(fields) == 0 {
		return params, obj
	}
	for _, param := range fields {
		params += param + ", "
	}
	params += "..."

	obj = prepareModel(fields)
	return params, obj
}

func addValidator(f *jen.File, reqFields RequiredParams) (*jen.File, error) {
	// Required Fields
	reqFieldVars := []string{createReqFields, readReqFields, updateReqFields, deleteReqFields, listReqFields}

	for _, varName := range reqFieldVars {
		switch varName {
		case createReqFields:
			f.Id("var ").Id(varName).Op("=").Index().String().Values(getLiteralsWithKeys(reqFields.CreateFields.RequiredParams)...)

		case readReqFields:
			f.Id("var ").Id(varName).Op("=").Index().String().Values(getLiteralsWithKeys(reqFields.ReadFields.RequiredParams)...)

		case updateReqFields:
			f.Id("var ").Id(varName).Op("=").Index().String().Values(getLiteralsWithKeys(reqFields.UpdateFields.RequiredParams)...)

		case deleteReqFields:
			f.Id("var ").Id(varName).Op("=").Index().String().Values(getLiteralsWithKeys(reqFields.DeleteFields.RequiredParams)...)

		case listReqFields:
			f.Id("var ").Id(varName).Op("=").Index().String().Values(getLiteralsWithKeys(reqFields.ListFields.RequiredParams)...)
		}
	}
	f.Line()
	// Add validation Method
	f.Func().Id(validationMethod).Params(
		// Input params
		jen.Id("fields").Id("[]string"), jen.Id("model").Id("*Model")).Id("*handler.ProgressEvent").
		Block(
			jen.Return(jen.Qual(validator, "ValidateModel").Call(jen.Id("fields"), jen.Id("model"))),
		)
	f.Line()
	return f, nil
}

func readRequiredParameters(schemaName string) (*RequiredParams, error) {
	reqFieldsSchemaPath := fmt.Sprintf("%s/mongodb-atlas-%s-req.json", schemaDir, schemaName)

	reqFile, err := os.ReadFile(reqFieldsSchemaPath)
	if err != nil {
		fmt.Println("ReqFields file read error:", err)
		return nil, err
	}
	var reqFields RequiredParams
	_ = json.Unmarshal(reqFile, &reqFields)
	return &reqFields, err
}

func getLiterals(arr []string) []jen.Code {
	var arrayLit []jen.Code
	for _, it := range arr {
		arrayLit = append(arrayLit, jen.Lit(it))
	}
	return arrayLit
}
func getLiteralsWithKeys(arr []string) []jen.Code {
	arrayLit := getLiterals(arr)
	arrayLit = append(arrayLit, jen.Lit(pubKey), jen.Lit(pvtKey))
	return arrayLit
}

func prepareModel(arr []string) string {
	str := ""
	for _, p := range arr {
		str += "\t" + p + ":" + "currentModel." + p + ",\n"
	}
	return str
}
func capitalize(key string) string {
	r := []rune(key)
	return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
}

type RequiredParams struct {
	CreateFields RequireParam `json:"CreateFields"`
	ReadFields   RequireParam `json:"ReadFields"`
	UpdateFields RequireParam `json:"UpdateFields"`
	DeleteFields RequireParam `json:"DeleteFields"`
	ListFields   RequireParam `json:"ListFields"`
	FileName     string       `json:"-"`
}

type RequireParam struct {
	InputParams    []string `json:",omitempty"`
	RequiredParams []string `json:",omitempty"`
}
