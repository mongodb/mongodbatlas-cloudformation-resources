package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/tidwall/pretty"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

//https://github.com/aws-cloudformation/cloudformation-cli/blob/master/src/rpdk/core/data/schema/provider.definition.schema.v1.json

const (
	url                = "https://github.com/aws-cloudformation/aws-cloudformation-rpdk.git"
	MongoDBAtlasPrefix = "MongoDB::Atlas::"
	Unique             = "Unique"
)

var optionalInputParams = []string{"envelope", "pretty", "apikeys", "app"}
var optionalReqParams = []string{"app"}

func main() {
	//dir := "/autogen/schema-gen"
	dir := "/schema-gen"

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
	dir = path + dir
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/mapping.json", dir))
	if err != nil {
		fmt.Printf("Path %s", path)
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	doc, err := openapi3.NewLoader().LoadFromFile(fmt.Sprintf("%s/swagger.json", dir))
	if err != nil {
		fmt.Printf("Path %s", path)
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	if doc == nil {
		os.Exit(1)
	}
	// validate the swagger.yaml matches Openapi spec
	_ = doc.Validate(context.Background())

	data := OpenApiMapping{}
	_ = json.Unmarshal(file, &data)

	var h Handlers
	json.Unmarshal([]byte(handler), &h)

	chn := make(chan CfnSchema, len(doc.Components.Schemas))
	reqFieldsChan := make(chan RequiredParams)

	done := make(chan bool)
	reqDone := make(chan bool)

	go generateSchemas(chn, done)
	go generateReqFields(reqFieldsChan, reqDone)

	definitions := make(map[string]Definitions, 0)
	var ids, readOnly, idsDef, readOnlyDef []string
	var cfn CfnSchema
	key := "params"
	var typeName string
	var description string
	requiredParams := RequiredParams{}
	var createReqParams []string

	for _, res := range data.Resources {

		queryParams := make(map[string]Property, 0)
		allMethodProps := make(map[string]map[string]Property, 0)
		bodySchema := make(map[string]map[string]Property, 0)

		typeName = res.TypeName

		for _, path := range res.OpenApiPaths {
			pathItem := doc.Paths.Find(path)
			if pathItem == nil {
				continue
			}
			if method := pathItem.Post; method != nil {

				// Read from Req params
				reqSchemaKeys, reqSchema, reqDefinitions, reqParams := readRequestBody(method, doc)
				createReqParams = reqParams

				// Read from Response params
				resSchemaKey, resSchema, resDefinitions := readResponseBody(method, doc)

				// Read from query params
				queryParams := readQueryParams(method)
				for _, reqSchemaKey := range reqSchemaKeys {
					bodySchema[key] = mergePropertyMaps(bodySchema[key], mergePropertyMaps(reqSchema[reqSchemaKey], resSchema[resSchemaKey]))
				}

				//Merge all params
				allMethodProps[key] = mergePropertyMaps(allMethodProps[key], mergePropertyMaps(bodySchema[key], queryParams))

				definitions = mergeDefinitionMaps(definitions, mergeDefinitionMaps(reqDefinitions, resDefinitions))

				definitions = defaultDefinition(definitions)
				readOnly, ids = readOnlyAndUniqueProperties(bodySchema)
				readOnlyDef, idsDef = readOnlyAndUniqueDefinitions(definitions)
				readOnly = append(readOnly, readOnlyDef...)
				ids = append(ids, idsDef...)

				//Required Props Query Params
				requiredParams.CreateFields.RequiredParams = requiredOnlyProperties(method.Parameters)
				//Required Props from Body
				requiredParams.CreateFields.RequiredParams = append(requiredParams.CreateFields.RequiredParams, capitalizeArray(reqParams)...)
				//All Props from Body
				requiredParams.CreateFields.InputParams = inputOnlyProperties(bodySchema[key])
			}

			if method := pathItem.Patch; method != nil {
				// Read from Req params
				reqSchemaKeys, reqSchema, reqDefinitions, reqParams := readRequestBody(method, doc)
				createReqParams = reqParams

				// Read from Response params
				resSchemaKey, resSchema, resDefinitions := readResponseBody(method, doc)

				// Read from query params
				queryParams := readQueryParams(method)
				for _, reqSchemaKey := range reqSchemaKeys {
					bodySchema[key] = mergePropertyMaps(bodySchema[key], mergePropertyMaps(reqSchema[reqSchemaKey], resSchema[resSchemaKey]))
				}

				//Merge all params
				allMethodProps[key] = mergePropertyMaps(allMethodProps[key], mergePropertyMaps(bodySchema[key], queryParams))

				definitions = mergeDefinitionMaps(definitions, mergeDefinitionMaps(reqDefinitions, resDefinitions))

				definitions = defaultDefinition(definitions)
				readOnly, ids = readOnlyAndUniqueProperties(bodySchema)
				readOnlyDef, idsDef = readOnlyAndUniqueDefinitions(definitions)
				readOnly = append(readOnly, readOnlyDef...)
				ids = append(ids, idsDef...)

				//Merge all params
				// Required Props from Body
				requiredParams.CreateFields.RequiredParams = append(requiredParams.CreateFields.RequiredParams, capitalizeArray(reqParams)...)
				// All Props from Body
				requiredParams.CreateFields.InputParams = inputOnlyProperties(bodySchema[key])
			}

			if method := pathItem.Patch; method != nil {
				// Read from Req params
				reqSchemaKeys, reqSchema, reqDefinitions, reqParams := readRequestBody(method, doc)
				createReqParams = reqParams

				// Read from Response params
				resSchemaKey, resSchema, resDefinitions := readResponseBody(method, doc)

				// Read from query params
				queryParams := readQueryParams(method)
				for _, reqSchemaKey := range reqSchemaKeys {
					bodySchema[key] = mergePropertyMaps(bodySchema[key], mergePropertyMaps(reqSchema[reqSchemaKey], resSchema[resSchemaKey]))
				}

				//Merge all params
				allMethodProps[key] = mergePropertyMaps(allMethodProps[key], mergePropertyMaps(bodySchema[key], queryParams))

				definitions = mergeDefinitionMaps(definitions, mergeDefinitionMaps(reqDefinitions, resDefinitions))

				definitions = defaultDefinition(definitions)
				readOnly, ids = readOnlyAndUniqueProperties(bodySchema)
				readOnlyDef, idsDef = readOnlyAndUniqueDefinitions(definitions)
				readOnly = append(readOnly, readOnlyDef...)
				ids = append(ids, idsDef...)

				//Required Props Query Params
				requiredParams.UpdateFields.RequiredParams = requiredOnlyProperties(method.Parameters)
				//Required Props from Body
				requiredParams.UpdateFields.RequiredParams = append(requiredParams.UpdateFields.RequiredParams, capitalizeArray(reqParams)...)
				//All Props from Body
				requiredParams.UpdateFields.InputParams = inputOnlyProperties(bodySchema[key])

			}
			if method := pathItem.Put; method != nil {
				// Read from Req params
				reqSchemaKeys, reqSchema, reqDefinitions, reqParams := readRequestBody(method, doc)
				createReqParams = reqParams

				// Read from Response params
				resSchemaKey, resSchema, resDefinitions := readResponseBody(method, doc)

				// Read from query params
				queryParams := readQueryParams(method)
				for _, reqSchemaKey := range reqSchemaKeys {
					bodySchema[key] = mergePropertyMaps(bodySchema[key], mergePropertyMaps(reqSchema[reqSchemaKey], resSchema[resSchemaKey]))
				}

				// Merge all params
				allMethodProps[key] = mergePropertyMaps(allMethodProps[key], mergePropertyMaps(bodySchema[key], queryParams))

				definitions = mergeDefinitionMaps(definitions, mergeDefinitionMaps(reqDefinitions, resDefinitions))

				definitions = defaultDefinition(definitions)
				readOnly, ids = readOnlyAndUniqueProperties(bodySchema)
				readOnlyDef, idsDef = readOnlyAndUniqueDefinitions(definitions)
				readOnly = append(readOnly, readOnlyDef...)
				ids = append(ids, idsDef...)

				// Required Props Query Params
				requiredParams.UpdateFields.RequiredParams = requiredOnlyProperties(method.Parameters)
				// Required Props from Body
				requiredParams.UpdateFields.RequiredParams = append(requiredParams.UpdateFields.RequiredParams, capitalizeArray(reqParams)...)
				// All Props from Body
				requiredParams.UpdateFields.InputParams = inputOnlyProperties(bodySchema[key])

			}

			if method := pathItem.Get; method != nil {
				if description == "" {
					description = readDescription(method.Tags[0], doc)
				}

				// Read from Req params
				_, _, _, reqParams := readRequestBody(method, doc)
				createReqParams = reqParams
				// Read from Response params
				resSchemaKey, resSchema, resDefinitions := readResponseBody(method, doc)

				// Read from query params
				queryParams := readQueryParams(method)
				bodySchema[key] = mergePropertyMaps(bodySchema[key], resSchema[resSchemaKey])

				// Merge all params
				allMethodProps[key] = mergePropertyMaps(allMethodProps[key], mergePropertyMaps(bodySchema[key], queryParams))

				definitions = mergeDefinitionMaps(definitions, resDefinitions)

				definitions = defaultDefinition(definitions)
				readOnly, ids = readOnlyAndUniqueProperties(bodySchema)
				readOnlyDef, idsDef = readOnlyAndUniqueDefinitions(definitions)
				readOnly = append(readOnly, readOnlyDef...)
				ids = append(ids, idsDef...)

				//Required Props Query Params
				requiredParams.ReadFields.RequiredParams = requiredOnlyProperties(method.Parameters)
				//Required Props from Body
				requiredParams.ReadFields.RequiredParams = append(requiredParams.ReadFields.RequiredParams, capitalizeArray(reqParams)...)
				//All Props from Body
				requiredParams.ReadFields.InputParams = inputOnlyProperties(bodySchema[key])

			}
			if method := pathItem.Delete; method != nil {
				if description == "" {
					description = readDescription(method.Tags[0], doc)
				}
				//Read from query params
				for _, pRef := range method.Parameters {
					newProps := readProperty(pRef.Value)
					for k, property := range newProps {
						if contains(k, optionalInputParams) {
							continue
						}
						queryParams[capitalize(k)] = property
					}
				}
				//Merge body params and query params
				allMethodProps[key] = mergePropertyMaps(allMethodProps[key], queryParams)

				//Required Props Query Params
				requiredParams.DeleteFields.RequiredParams = requiredOnlyProperties(method.Parameters)

			}
		}
		if allMethodProps[key] != nil {
			allMethodProps[key] = defaultProperty(allMethodProps[key])

			requiredParams.FileName = res.TypeName
			cfn = CfnSchema{
				AdditionalProperties: false,
				Definitions:          definitions,
				Description:          description,
				Handlers:             h,
				PrimaryIdentifier:    ids,
				Properties:           allMethodProps[key],
				ReadOnlyProperties:   readOnly,
				Required:             createReqParams,
				TypeName:             MongoDBAtlasPrefix + typeName,
				SourceUrl:            url,
				FileName:             typeName,
			}
			chn <- cfn
			reqFieldsChan <- requiredParams
		}
	}

	close(chn)
	<-done

	close(reqFieldsChan)
	<-reqDone

}

func readRequestBody(method *openapi3.Operation, doc *openapi3.T) (schemaKeys []string, reqSchema map[string]map[string]Property, definitions map[string]Definitions, requiredParams []string) {
	reqBody := method.RequestBody
	if reqBody != nil && reqBody.Value != nil && reqBody.Value.Content["application/json"] != nil &&
		reqBody.Value.Content["application/json"].Schema != nil {
		reqSchemaKey := filepath.Base(reqBody.Value.Content["application/json"].Schema.Ref)
		schemaKeys = append(schemaKeys, capitalize(reqSchemaKey))
		//Read from Request body
		if doc.Components.Schemas[filepath.Base(reqSchemaKey)] != nil {
			value := *doc.Components.Schemas[filepath.Base(reqSchemaKey)]
			requiredParams = value.Value.Required
			reqSchema, definitions = processSchema(reqSchemaKey, &value, doc.Components.Schemas)

			// Read Discriminator params
			if value.Value.Discriminator != nil {
				for key, def := range value.Value.Discriminator.Mapping {
					fmt.Println(def, key)
					schemaKey := def[strings.LastIndex(def, "/")+1:]

					if doc.Components.Schemas[filepath.Base(schemaKey)].Value != nil && doc.Components.Schemas[filepath.Base(schemaKey)].Value.AllOf != nil {
						allOf := doc.Components.Schemas[filepath.Base(schemaKey)].Value.AllOf
						for i := range allOf {

							defs := processDefinitionsSchema(schemaKey, allOf[i], doc.Components.Schemas)
							definitions[schemaKey] = defs[schemaKey]
							if contains(schemaKey, schemaKeys) {
								continue
							}
							schemaKeys = append(schemaKeys, capitalize(schemaKey))

						}
					}
				}
			}
		}
	}
	return schemaKeys, reqSchema, definitions, requiredParams
}

func readResponseBody(method *openapi3.Operation, doc *openapi3.T) (schemaKey string, resSchema map[string]map[string]Property, definitions map[string]Definitions) {
	var resSchemaKey string

	if method.Responses["200"] != nil && method.Responses["200"].Value != nil && method.Responses["200"].Value.Content["application/json"] != nil &&
		method.Responses["200"].Value.Content["application/json"].Schema != nil {

		resSchemaKey = filepath.Base(method.Responses["200"].Value.Content["application/json"].Schema.Ref)
		// Read from Request body
		if doc.Components.Schemas[filepath.Base(resSchemaKey)] != nil {
			value := *doc.Components.Schemas[filepath.Base(resSchemaKey)]
			resSchema, definitions = processSchema(resSchemaKey, &value, doc.Components.Schemas)

		}

	}
	return capitalize(resSchemaKey), resSchema, definitions
}

func readQueryParams(method *openapi3.Operation) map[string]Property {
	queryParams := map[string]Property{}
	for _, pRef := range method.Parameters {
		newProps := readProperty(pRef.Value)
		for k, property := range newProps {
			if contains(k, optionalInputParams) {
				continue
			}
			queryParams[capitalize(k)] = property
		}
	}
	return queryParams
}

func defaultProperty(defaultProperty map[string]Property) map[string]Property {
	defaultProperty["ApiKeys"] = Property{Ref: "#/definitions/ApiKeyDefinition"}
	return defaultProperty
}

func primaryIdentifier() Property {
	return Property{
		Type: "string",
	}
}

func defaultDefinition(definitions map[string]Definitions) map[string]Definitions {
	defaultDef := make(map[string]Property)
	defaultDef["PublicKey"] = Property{
		Type: "string",
	}
	defaultDef["PrivateKey"] = Property{
		Type: "string",
	}
	definitions["ApiKeyDefinition"] = Definitions{
		Type:                 "object",
		Properties:           defaultDef,
		AdditionalProperties: false,
	}
	return definitions
}

func readOnlyAndUniqueDefinitions(def map[string]Definitions) ([]string, []string) {
	var readOnly []string
	var ids []string
	for id, p := range def {
		for rid, ro := range p.Properties {
			if ro.ReadOnly {
				readOnly = append(readOnly, fmt.Sprintf("/%s/%s/%s", "definitions", id, rid))
			}
			if strings.HasPrefix(ro.Description, Unique) {
				ids = append(ids, fmt.Sprintf("/%s/%s/%s", "definitions", capitalize(id), capitalize(rid)))
			}
		}
	}
	return readOnly, ids
}

func readOnlyAndUniqueProperties(properties map[string]map[string]Property) ([]string, []string) {
	var readOnly []string
	var ids []string

	for _, p := range properties {
		for rid, ro := range p {
			if ro.ReadOnly {
				readOnly = append(readOnly, fmt.Sprintf("/%s/%s", "properties", rid))
			}
			if strings.HasPrefix(ro.Description, Unique) {
				ids = append(ids, fmt.Sprintf("/%s/%s", "properties", rid))
			}
		}
	}
	return readOnly, ids
}

func requiredOnlyProperties(properties openapi3.Parameters) []string {
	var requiredParams []string

	for _, p := range properties {
		if p.Value.Required && !contains(p.Value.Name, optionalReqParams) && !contains(p.Value.Name, requiredParams) {
			requiredParams = append(requiredParams, capitalize(p.Value.Name))
		}
	}
	return requiredParams
}
func inputOnlyProperties(bodyParams map[string]Property) []string {
	var inputParams []string

	for name, val := range bodyParams {
		if val.Required != nil && !contains(name, optionalInputParams) && !contains(name, inputParams) {
			inputParams = append(inputParams, capitalizeArray(val.Required)...)
		}
		inputParams = append(inputParams, capitalize(name))
	}

	return inputParams
}

func generateSchemas(chn chan CfnSchema, done chan bool) {
	for cfn := range chn {
		rankingsJson, _ := json.Marshal(cfn)
		result := pretty.Pretty(rankingsJson)

		dir := "configs"
		if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(dir, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		}

		fileName := fmt.Sprintf("%s/mongodb-atlas-%s.json", dir, strings.ToLower(cfn.FileName))

		err := ioutil.WriteFile(fileName, result, 0644)
		if err != nil {
			print(err)
		}
	}
	done <- true
}

func generateReqFields(reqChan chan RequiredParams, reqDone chan bool) {
	for reqFlds := range reqChan {
		spew.Dump(reqFlds)
		fieldsJson, err := json.Marshal(reqFlds)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := pretty.Pretty(fieldsJson)

		dir := "configs"
		if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(dir, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		}

		fileName := fmt.Sprintf("%s/mongodb-atlas-%s-req.json", dir, strings.ToLower(reqFlds.FileName))

		err = ioutil.WriteFile(fileName, result, 0644)
		if err != nil {
			print(err)
		}
	}
	reqDone <- true
}

var pDefinitions = make(map[string]Definitions, 0)

func processSchema(id string, v *openapi3.SchemaRef, schemas openapi3.Schemas) (map[string]map[string]Property,
	map[string]Definitions) {

	definitions := make(map[string]*openapi3.SchemaRef, 0)
	properties := make(map[string]map[string]Property, 0)

	pMap := make(map[string]Property, 0)

	for k, p := range v.Value.Properties {
		// capture only those properties that are required by Cloudformation from openAPI spec
		pMap[capitalize(k)] = property(p)
		if p.Ref != "" {
			//v.Value.Properties[k].Ref = "" // update ref
			// definition already processed
			if _, ok := pDefinitions[strings.ReplaceAll(filepath.Base(p.Ref), "_", "")]; !ok {
				if val, ok1 := schemas[filepath.Base(p.Ref)]; ok1 {
					// collect definitions
					definitions[strings.ReplaceAll(filepath.Base(p.Ref), "_", "")] = val
				}
			}
		}
		if p.Value.Items != nil && len(p.Value.Items.Ref) != 0 {
			//p.Value.Items.Ref = "" // update ref

			ref := filepath.Base(p.Value.Items.Ref)
			ref1 := strings.ReplaceAll(filepath.Base(ref), "_", "")
			if val, ok := schemas[ref]; ok {
				// collect definitions
				definitions[ref1] = val
			}
		}
	}
	properties[id] = defaultProperty(pMap)
	for k, def := range definitions {
		p, _ := processSchema(k, def, schemas)
		for _, v1 := range p {
			pDefinitions[capitalize(k)] = Definitions{
				Type:                 def.Value.Type,
				Properties:           v1,
				AdditionalProperties: false,
			}
		}
	}

	return properties, pDefinitions
}

func processDefinitionsSchema(id string, ref *openapi3.SchemaRef, schemas openapi3.Schemas) (definitions map[string]Definitions) {
	var pDefinitions = make(map[string]Definitions, 0)
	p, _ := processSchema(id, ref, schemas)
	for _, v1 := range p {
		pDefinitions[capitalize(id)] = Definitions{
			Type:                 ref.Value.Type,
			Properties:           v1,
			AdditionalProperties: false,
		}
	}

	return pDefinitions
}

func property(val *openapi3.SchemaRef) Property {
	return Property{
		Type:           val.Value.Type,
		Description:    val.Value.Description,
		MaxLength:      val.Value.MaxLength,
		MinLength:      val.Value.MinLength,
		InsertionOrder: false,
		Ref: func() string {
			if val.Ref != "" {
				return "#/definitions/" + capitalize(strings.ReplaceAll(filepath.Base(val.Ref), "_", ""))
			}
			return ""
		}(),
		AdditionalProperties: false,
		Enum:                 val.Value.Enum,
		Pattern:              val.Value.Pattern,
		Items: func() *Items {
			var ref string
			if val.Value.Items != nil {
				if val.Value.Items.Ref != "" {
					ref = "#/definitions/" + capitalize(strings.ReplaceAll(filepath.Base(val.Value.Items.Ref), "_", ""))
				}
				return &Items{
					Ref:  ref,
					Type: val.Value.Items.Value.Type,
					Enum: val.Value.Items.Value.Enum,
				}
			}

			return nil
		}(),
		ReadOnly: val.Value.ReadOnly,
		Required: val.Value.Required,
	}
}

func readProperty(parameter *openapi3.Parameter) map[string]Property {
	val := parameter.Schema
	return map[string]Property{
		parameter.Name: {
			Type:           parameter.Schema.Value.Type,
			Description:    parameter.Description,
			MaxLength:      val.Value.MaxLength,
			MinLength:      val.Value.MinLength,
			InsertionOrder: false,
			Ref: func() string {
				if val.Ref != "" {
					return "#/definitions/" + capitalize(strings.ReplaceAll(filepath.Base(val.Ref), "_", ""))
				}
				return ""
			}(),
			AdditionalProperties: false,
			Enum:                 val.Value.Enum,
			Pattern:              val.Value.Pattern,
			Items: func() *Items {
				var ref string
				if val.Value.Items != nil {
					if val.Value.Items.Ref != "" {
						ref = "#/definitions/" + strings.ReplaceAll(filepath.Base(val.Value.Items.Ref), "_", "")
					}
					return &Items{
						Ref:  ref,
						Type: val.Value.Items.Value.Type,
						Enum: val.Value.Items.Value.Enum,
					}
				}

				return nil
			}(),
			ReadOnly: val.Value.ReadOnly,
			Required: val.Value.Required,
		},
	}
}

func capitalize(key string) string {
	if len(key) <= 0 {
		return key
	}
	r := []rune(key)
	return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
}

func capitalizeArray(keys []string) []string {
	var ks []string
	for _, elem := range keys {
		ks = append(ks, capitalize(elem))
	}
	return ks
}
func contains(key string, values []string) bool {
	for _, elem := range values {
		if strings.ToLower(key) == strings.ToLower(elem) {
			return true
		}
	}
	return false
}

func mergePropertyMaps(map1 map[string]Property, map2 map[string]Property) map[string]Property {
	if map1 == nil {
		return map2
	}
	if map2 == nil {
		return map1
	}
	for k, v := range map2 {
		map1[k] = v
	}
	return map1
}

func mergeDefinitionMaps(map1 map[string]Definitions, map2 map[string]Definitions) map[string]Definitions {
	if map1 == nil {
		return map2
	}
	if map2 == nil {
		return map1
	}
	for k, v := range map2 {
		map1[k] = v
	}
	return map1
}

func readDescription(tagName string, doc *openapi3.T) string {
	for i := range doc.Tags {
		if tagName == doc.Tags[i].Name {
			return doc.Tags[i].Description
		}
	}
	return tagName
}
