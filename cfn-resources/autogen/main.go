package main

import (
	"encoding/json"
	"flag"
	"fmt"
	. "github.com/dave/jennifer/jen"
	"os"
	"unicode"
)

const (
	pubKey           = "ApiKeys.PublicKey"
	pvtKey           = "ApiKeys.PrivateKey"
	schemaDir        = "configs"
	validatorPackage = "validations"
	resourceFilePath = "/cmd/resource/resource.go"
	dir              = "/cmd/"
	validatorFile    = "/validator.go"
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

	//add validator.go
	resourceFileFullPath := fmt.Sprintf("%s%s", *resource, resourceFilePath)

	// add resource.go
	addResource(resourceFileFullPath, *resource, *schemaName)

}

func addResource(path, resourceName, schemaName string) {
	f := NewFile("resource")
	funcNames := []string{CreateMethod, ReadMethod, UpdateMethod, DeleteMethod, ListMethod}

	//constants := "github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	//validations := fmt.Sprintf("github.com/mongodb/mongodbatlas-cloudformation-resources/%s/cmd/validations", resourceName)

	//Imports
	f.ImportName(handler, "handler")
	f.ImportName(cloudformation, "cloudformation")
	f.ImportName("errors", "errors")
	f.ImportName(util, "util")
	f.ImportName(validator, "validator")
	f.ImportName(progressEvent, "progress_events")
	f.ImportAlias(log, "log")

	//Required Fields mapping
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

	//Add setup method
	f.Func().Id(setupMethod).Params().
		Block(
			Qual(util, "SetupLogger").Call(Lit("mongodb-atlas-" + schemaName)),
		)
	f.Line()

	//Add handlers
	for _, fn := range funcNames {
		inputParameters, modelObj := generateInputParams(fn, reqFields)
		if len(inputParameters) > 0 {
			inputParameters = "Considerable params from currentModel: \n" + inputParameters
		}

		atlasFunc := " res , resModel, err := client." + capitalize(schemaName) + "." + fn + "(context.Background(),&mongodbatlas." + capitalize(schemaName) + "{\n" + modelObj + "})"

		f.Func().Id(fn).Params(
			//Input Parameters
			Id("req").Qual(handler, "Request"),
			Id("prevModel").Id("*Model"),
			Id("currentModel").Id("*Model")).
			//Return types
			Id("(").Qual(handler, "ProgressEvent").Id(",").Id("error").Id(")").

			// Function starts
			Block(
				//log setup()
				Id("setup").Params(),
				Line(),
				//Debug log
				Qual(log, "Debugf").Params(Lit(fn+"() currentModel:%+v"), Id("currentModel")),
				Line(),

				//Validator
				Comment("Validation"),
				Id("modelValidation").Id(":=").Id(validationMethod).Call(Id(fn+"RequiredFields"), Id("currentModel")),
				If(Id("modelValidation").Op("!=").Id("nil")).Block(
					Return(
						Id("*modelValidation"),
						Id("nil"),
					),
				),
				Line(),

				//Atlas Client
				Comment("Create atlas client"),
				List(Id("client"), Err()).Op(":=").Qual(util, "CreateMongoDBClient").
					Call(Id("*currentModel.ApiKeys.PublicKey"),
						Id("*currentModel.ApiKeys.PrivateKey")),

				If(Id("err").Op("!=").Id("nil")).Block(
					Qual(log, "Debugf").
						Call(Lit(fn+" - error: %+v"), Id("err")),

					Return(
						Qual(handler, "ProgressEvent").
							Values(Dict{
								Id("OperationStatus"):  Qual(handler, "Failed"),
								Id("Message"):          Id("err.Error()"),
								Id("HandlerErrorCode"): Qual(cloudformation, "HandlerErrorCodeInvalidRequest"),
							}),
						Id("nil"),
					),
				),

				Id("var").Id(res).Id("*").Qual(atlas, "Response"),
				Line(),

				//Pseudocode
				Comment(inputParameters),
				Comment(" // Pseudocode:\n"+atlasFunc+"\n\n"),
				Line(),

				If(Id("err").Op("!=").Id("nil")).Block(
					Qual(log, "Debugf").
						Call(Lit(fn+" - error: %+v"), Id("err")),

					Return(
						Qual(progressEvent, "GetFailedEventByResponse").
							Params(Id("err.Error()"), Id("res.Response")),
						Id("nil"),
					),
				),

				//Log
				Qual(log, "Debugf").
					Params(Lit("Atlas Client %v"), Id("client")),

				Line(),
				//Progress Event Response
				Comment("Response"),
				//Return statement
				Return(List(Qual(handler, "ProgressEvent").
					Values(Dict{
						Id("OperationStatus"): Qual(handler, "Success"),
						Id("ResourceModel"):   Id("currentModel"),
					}), Id("nil"))),
			)
		f.Line()
	}
	fmt.Println("saving file at", path)
	err = f.Save(path)

	fmt.Println(err)
}

func generateInputParams(fn string, params *RequiredParams) (string, string) {
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
	return "", ""
}

func generateParamString(fields []string) (string, string) {
	params := ""
	obj := ""
	if len(fields) <= 0 {
		return params, obj
	}
	for _, param := range fields {
		params += param + ", "
	}
	params = params + "..."

	obj = prepareModel(fields)
	return params, obj
}

func addValidator(f *File, reqFields RequiredParams) (*File, error) {

	//Required Fields
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
	//Add validation Method
	f.Func().Id(validationMethod).Params(
		//Input params
		Id("fields").Id("[]string"), Id("model").Id("*Model")).Id("*handler.ProgressEvent").
		Block(
			Return(Qual(validator, "ValidateModel").Call(Id("fields"), Id("model"))),
		)
	f.Line()
	return f, nil
}

func readResourceSchema(f *File, schemaName string) (error, *File) {
	schemaFilePath := fmt.Sprintf("%s/mongodb-atlas-%s.json", schemaDir, schemaName)
	//Schema mapping
	fmt.Println("Schema path:", schemaFilePath)
	file, err := os.ReadFile(schemaFilePath)
	if err != nil {
		fmt.Println("SchemaFile read error:", err)
		return err, f
	}
	schema := ResourceSchema{}
	_ = json.Unmarshal(file, &schema)
	return err, f
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

func getLiterals(arr []string) []Code {
	var arrayLit []Code
	for _, it := range arr {
		arrayLit = append(arrayLit, Lit(it))
	}
	return arrayLit
}
func getLiteralsWithKeys(arr []string) []Code {
	arrayLit := getLiterals(arr)
	arrayLit = append(arrayLit, Lit(pvtKey))
	arrayLit = append(arrayLit, Lit(pubKey))
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

type ResourceSchema struct {
	AdditionalProperties bool `json:"additionalProperties"`
	Definitions          struct {
		ApiKeyDefinition struct {
			Type       string `json:"type"`
			Properties struct {
				PrivateKey struct {
					Type string `json:"type"`
				} `json:"PrivateKey"`
				PublicKey struct {
					Type string `json:"type"`
				} `json:"PublicKey"`
			} `json:"properties"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"apiKeyDefinition"`
	} `json:"definitions"`
	Description string `json:"description"`
	Handlers    struct {
		Create struct {
			Permissions []string `json:"permissions"`
		} `json:"create"`
		Read struct {
			Permissions []string `json:"permissions"`
		} `json:"read"`
		Update struct {
			Permissions []string `json:"permissions"`
		} `json:"update"`
		Delete struct {
			Permissions []string `json:"permissions"`
		} `json:"delete"`
	} `json:"handlers"`
	PrimaryIdentifier []string `json:"primaryIdentifier"`
	Properties        struct {
		App struct {
			Type        string   `json:"type"`
			Description string   `json:"description"`
			Enum        []string `json:"enum"`
		} `json:"app"`
		Envelope struct {
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"envelope"`
		GroupId struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			MaxLength   int    `json:"maxLength"`
			MinLength   int    `json:"minLength"`
			Pattern     string `json:"pattern"`
		} `json:"groupId"`
		IncludeCount struct {
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"includeCount"`
		ItemsPerPage struct {
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"itemsPerPage"`
		PageNum struct {
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"pageNum"`
		Pretty struct {
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"pretty"`
		Username struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Pattern     string `json:"pattern"`
		} `json:"username"`
	} `json:"properties"`
	TypeName  string `json:"typeName"`
	SourceUrl string `json:"sourceUrl"`
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
