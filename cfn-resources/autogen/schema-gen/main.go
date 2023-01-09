package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/tidwall/pretty"
)

// https://github.com/aws-cloudformation/cloudformation-cli/blob/master/src/rpdk/core/data/schema/provider.definition.schema.v1.json

const (
	url                = "https://github.com/aws-cloudformation/aws-cloudformation-rpdk.git"
	MongoDBAtlasPrefix = "MongoDB::Atlas::"
	Unique             = "Unique"
	OpenAPISpecPath    = "https://www.mongodb.com/8c07de79-53d6-41d8-8fc8-bacdf7f271fa"
	Dir                = "/schema-gen" // For debugging use 	"/autogen/schema-gen"
	SchemasDir         = "schemas"
	CurrentDir         = "schema-gen"
	LatestSwaggerFile  = "swagger.latest.json"
)

var optionalInputParams = []string{"envelope", "pretty", "apikeys", "app"}
var optionalReqParams = []string{"app"}

func main() {
	compare := false

	file, doc, err := readConfig(compare)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	data := OpenAPIMapping{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	var h Handlers
	err = json.Unmarshal([]byte(handler), &h)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	chn := make(chan CfnSchema, len(doc.Components.Schemas))
	reqFieldsChan := make(chan RequiredParams)

	done := make(chan bool)
	reqDone := make(chan bool)
	go generateSchemas(chn, done, compare)
	go generateReqFields(reqFieldsChan, reqDone, compare)

	for _, res := range data.Resources {
		definitions := make(map[string]Definitions, 0)
		var ids, readOnly, idsDef, readOnlyDef []string
		var cfn CfnSchema
		key := "params"
		var typeName string
		var description string
		requiredParams := RequiredParams{}
		var createReqParams []string

		queryParams := make(map[string]Property, 0)
		allMethodProps := make(map[string]map[string]Property, 0)
		bodySchema := make(map[string]map[string]Property, 0)
		typeName = capitalize(res.TypeName)

		for _, path := range res.OpenAPIPaths {
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

				// Merge all params
				allMethodProps[key] = mergePropertyMaps(allMethodProps[key], mergePropertyMaps(bodySchema[key], queryParams))
				definitions = mergeDefinitionMaps(definitions, mergeDefinitionMaps(reqDefinitions, resDefinitions))

				definitions = defaultDefinition(definitions)
				readOnly, ids = readOnlyAndUniqueProperties(bodySchema)
				readOnlyDef, idsDef = readOnlyAndUniqueDefinitions(definitions)
				readOnly = append(readOnly, readOnlyDef...)
				ids = append(ids, idsDef...)

				// Required Props from Body
				requiredParams.CreateFields.RequiredParams = append(requiredParams.CreateFields.RequiredParams, capitalizeArray(reqParams)...)
				// All Props from Body
				requiredParams.CreateFields.InputParams = inputOnlyProperties(bodySchema[key])
			}

			updateMethods := []*openapi3.Operation{pathItem.Put, pathItem.Patch}
			for i := range updateMethods {
				method := updateMethods[i]
				if method == nil {
					continue
				}
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

				// Required Props Query Params
				requiredParams.ReadFields.RequiredParams = requiredOnlyProperties(method.Parameters)
				// Required Props from Body
				requiredParams.ReadFields.RequiredParams = append(requiredParams.ReadFields.RequiredParams, capitalizeArray(reqParams)...)
				// All Props from Body
				requiredParams.ReadFields.InputParams = inputOnlyProperties(bodySchema[key])
			}
			if method := pathItem.Delete; method != nil {
				if description == "" {
					description = readDescription(method.Tags[0], doc)
				}
				// Read from query params
				for i := range method.Parameters {
					newProps := readProperty(method.Parameters[i].Value)
					for k := range newProps {
						if contains(k, optionalInputParams) {
							continue
						}
						queryParams[capitalize(k)] = newProps[k]
					}
				}
				// Merge body params and query params
				allMethodProps[key] = mergePropertyMaps(allMethodProps[key], queryParams)

				// Required Props Query Params
				requiredParams.DeleteFields.RequiredParams = requiredOnlyProperties(method.Parameters)
			}
		}
		if allMethodProps[key] != nil {
			allMethodProps[key] = defaultProperty(allMethodProps[key])

			sort.Strings(createReqParams)
			sort.Strings(readOnly)
			sort.Strings(ids)

			cfn = CfnSchema{
				AdditionalProperties: false,
				Definitions:          sortDefinitions(definitions),
				Description:          description,
				Handlers:             h,
				PrimaryIdentifier:    ids,
				Properties:           sortProperties(allMethodProps[key]),
				ReadOnlyProperties:   readOnly,
				Required:             createReqParams,
				TypeName:             MongoDBAtlasPrefix + typeName,
				SourceURL:            url,
				FileName:             res.TypeName,
			}
			chn <- cfn

			requiredParams.FileName = res.TypeName
			reqFieldsChan <- requiredParams
		}
	}
	close(chn)
	<-done

	close(reqFieldsChan)
	<-reqDone
}

func sortProperties(properties map[string]Property) (props map[string]Property) {
	keys := make([]string, len(properties))
	for key := range properties {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	props = make(map[string]Property, len(properties))
	for _, key := range keys {
		props[key] = properties[key]
	}
	return props
}

func sortDefinitions(properties map[string]Definitions) (props map[string]Definitions) {
	keys := make([]string, len(properties))
	for key := range properties {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	props = make(map[string]Definitions, len(properties))
	for _, key := range keys {
		props[key] = properties[key]
	}
	return props
}

func downloadOpenAPISpec(url, fileName string) (err error) {
	spaceClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, OpenAPISpecPath, http.NoBody)
	if err != nil {
		return err
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		return getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}

	err = os.WriteFile(fileName, body, 0600)
	return err
}

func getCurrentDir() (path string, err error) {
	path, err = os.Getwd()
	if err != nil {
		fmt.Printf("Error while fetching current directory: %+v", err)
		return path, err
	}
	dir := path + Dir
	return dir, err
}

func readConfig(compare bool) ([]byte, *openapi3.T, error) {
	dir, err := getCurrentDir()
	if err != nil {
		return nil, nil, err
	}
	file, err := os.ReadFile(fmt.Sprintf("%s/mapping.json", dir))
	if err != nil {
		return nil, nil, err
	}

	openAPISpecFile := fmt.Sprintf("%s/swagger.json", dir)
	// For comparison download the latest openAPIspec file
	if compare {
		openAPISpecFile = fmt.Sprintf("%s/%s", dir, LatestSwaggerFile)
		if err := downloadOpenAPISpec(OpenAPISpecPath, openAPISpecFile); err != nil {
			return []byte{}, nil, err
		}
	}
	doc, err := openapi3.NewLoader().LoadFromFile(openAPISpecFile)
	if err != nil {
		return nil, nil, err
	}

	if doc == nil {
		fmt.Println("empty document found")
		os.Exit(1)
	}

	return file, doc, err
}

func readRequestBody(method *openapi3.Operation, doc *openapi3.T) (schemaKeys []string, reqSchema map[string]map[string]Property, definitions map[string]Definitions, requiredParams []string) {
	reqBody := method.RequestBody
	if reqBody != nil && reqBody.Value != nil && reqBody.Value.Content["application/json"] != nil &&
		reqBody.Value.Content["application/json"].Schema != nil {
		reqSchemaKey := filepath.Base(reqBody.Value.Content["application/json"].Schema.Ref)
		schemaKeys = append(schemaKeys, capitalize(reqSchemaKey))
		// Read from Request body
		if doc.Components.Schemas[filepath.Base(reqSchemaKey)] != nil {
			value := *doc.Components.Schemas[filepath.Base(reqSchemaKey)]
			requiredParams = value.Value.Required
			reqSchema, definitions = processSchema(reqSchemaKey, &value, doc.Components.Schemas)

			// Read Discriminator params
			if value.Value.Discriminator != nil {
				for _, def := range value.Value.Discriminator.Mapping {
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
	for i := range method.Parameters {
		newProps := readProperty(method.Parameters[i].Value)
		for k := range newProps {
			if contains(k, optionalInputParams) {
				continue
			}
			queryParams[capitalize(k)] = newProps[k]
		}
	}
	return queryParams
}

func defaultProperty(defaultProperty map[string]Property) map[string]Property {
	defaultProperty["ApiKeys"] = Property{Ref: "#/definitions/ApiKeyDefinition"}
	return defaultProperty
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

func readOnlyAndUniqueDefinitions(def map[string]Definitions) (readOnly []string, ids []string) {
	for id := range def {
		for k := range def[id].Properties {
			if def[id].Properties[k].ReadOnly {
				readOnly = append(readOnly, fmt.Sprintf("/%s/%s/%s", "definitions", id, k))
			}
			if strings.HasPrefix(def[id].Properties[k].Description, Unique) {
				ids = append(ids, fmt.Sprintf("/%s/%s/%s", "definitions", capitalize(id), capitalize(k)))
			}
		}
	}
	return readOnly, ids
}

func readOnlyAndUniqueProperties(properties map[string]map[string]Property) (readOnly []string, ids []string) {
	for _, p := range properties {
		for k := range p {
			if p[k].ReadOnly {
				readOnly = append(readOnly, fmt.Sprintf("/%s/%s", "properties", k))
			}
			if strings.HasPrefix(p[k].Description, Unique) {
				ids = append(ids, fmt.Sprintf("/%s/%s", "properties", k))
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

	for name := range bodyParams {
		if bodyParams[name].Required != nil && !contains(name, optionalInputParams) && !contains(name, inputParams) {
			inputParams = append(inputParams, capitalizeArray(bodyParams[name].Required)...)
		}
		inputParams = append(inputParams, capitalize(name))
	}

	return inputParams
}

func generateSchemas(chn chan CfnSchema, done chan bool, compare bool) {
	for cfn := range chn {
		rankingsJSON, err := json.Marshal(cfn)
		if err != nil {
			fmt.Printf("error in generateSchemas : %+v ", err)
			return
		}
		result := pretty.Pretty(rankingsJSON)

		dir, err := getCurrentDir()
		if err != nil {
			return
		}

		schemaDir := strings.Replace(dir, CurrentDir, SchemasDir, 1)

		if _, err := os.Stat(schemaDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(schemaDir, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		}

		schemaFilePath := fmt.Sprintf("%s/mongodb-atlas-%s.json", schemaDir, strings.ToLower(cfn.FileName))
		latestSchemaFilePath := ""

		// create required schema file
		if compare {
			latestSchemaDir := strings.Replace(dir, CurrentDir, SchemasDir, 1)
			if _, err := os.Stat(latestSchemaDir); errors.Is(err, os.ErrNotExist) {
				err := os.Mkdir(latestSchemaDir, os.ModePerm)
				if err != nil {
					print(err)
					done <- true
				}
			}
			latestSchemaFilePath = fmt.Sprintf("%s/mongodb-atlas-%s-latest.json", latestSchemaDir, strings.ToLower(cfn.FileName))

			// Write schema into the latest file
			err = os.WriteFile(latestSchemaFilePath, result, 0600)
			if err != nil {
				print(err)
				done <- true
				return
			}
		}

		if compare && latestSchemaFilePath != "" {
			_, err = CompareJSONFiles(cfn.FileName, schemaFilePath, latestSchemaFilePath)
			if err != nil {
				print(err)
			}
			// Delete the latest file created for comparison
			err = os.RemoveAll(latestSchemaFilePath)
			if err != nil {
				print(err)
			}
		}

		// Update with the schema file with the latest schema
		err = os.WriteFile(schemaFilePath, result, 0600)
		if err != nil {
			print(err)
			done <- true
			return
		}
	}
	done <- true
}

func generateReqFields(reqChan chan RequiredParams, reqDone chan bool, compare bool) {
	for reqFlds := range reqChan {
		fieldsJSON, err := json.Marshal(reqFlds)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := pretty.Pretty(fieldsJSON)

		if _, err := os.Stat(SchemasDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(SchemasDir, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		}

		fileName := fmt.Sprintf("%s/mongodb-atlas-%s-req.json", SchemasDir, strings.ToLower(reqFlds.FileName))
		if compare {
			fileName = fmt.Sprintf("%s/mongodb-atlas-%s-req-latest.json", SchemasDir, strings.ToLower(reqFlds.FileName))
		}
		err = os.WriteFile(fileName, result, 0600)
		if err != nil {
			print(err)
		}
	}
	reqDone <- true
}

func processSchema(id string, v *openapi3.SchemaRef, schemas openapi3.Schemas) (properties map[string]map[string]Property,
	pDefinitions map[string]Definitions) {
	definitions := make(map[string]*openapi3.SchemaRef, 0)

	properties = make(map[string]map[string]Property, 0)
	pDefinitions = make(map[string]Definitions, 0)

	pMap := make(map[string]Property, 0)

	for k, p := range v.Value.Properties {
		// capture only those properties that are required by Cloudformation from openAPI spec
		pMap[capitalize(k)] = property(p)
		if p.Ref != "" {
			// definition already processed
			if _, ok := pDefinitions[strings.ReplaceAll(filepath.Base(p.Ref), "_", "")]; !ok {
				if val, ok1 := schemas[filepath.Base(p.Ref)]; ok1 {
					// collect definitions
					definitions[strings.ReplaceAll(filepath.Base(p.Ref), "_", "")] = val
				}
			}
		}
		if p.Value.Items != nil && p.Value.Items.Ref != "" {
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
	if key == "" {
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
		if strings.EqualFold(key, elem) {
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
	for k := range map2 {
		map1[k] = map2[k]
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
