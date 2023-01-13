package gen

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

var outputDir string

//go:embed templates/api_op.go.tpl
var apiFileTemplate string

var tpl *template.Template

func init() {
	var err error
	funcMap := template.FuncMap{
		"toCamel": strcase.ToCamel,
		"cleanLF": func(str string) string {
			return strings.ReplaceAll(str, "\n", " ")
		},
	}

	tpl, err = template.New("api_op").Funcs(funcMap).Parse(apiFileTemplate)
	if err != nil {
		panic(err)
	}
}

func GenApis(output string, apiSchemas []ApiSchema) error {
	outputDir = output
	return genApis(apiSchemas)
}

func genApis(apiSchemas []ApiSchema) error {

	for _, apiSchema := range apiSchemas {
		err := genApi(apiSchema)
		if err != nil {
			return err
		}
	}

	return nil
}

func genApi(apiSchema ApiSchema) error {

	for method, api := range apiSchema.Info {
		switch method {
		case ApiMethodGet:
		case ApiMethodPost:
		case ApiMethodPut:
		case ApiMethodDelete:
		default:
			return fmt.Errorf("unknown method: %s", method)
		}

		err := genApiFile(apiSchema.Path, api)
		if err != nil {
			return err
		}
	}

	err := genApis(apiSchema.Children)
	if err != nil {
		return err
	}

	return nil

}

func genApiFile(path string, api ApiInfo) error {
	apiDefine := NewApiDefine(path, api)

	var buf bytes.Buffer

	err := tpl.Execute(&buf, apiDefine)
	if err != nil {
		return err
	}
	apiFileContent := buf.Bytes()

	// format code
	apiFileContent, err = format.Source(apiFileContent)
	if err != nil {
		return err
	}

	os.WriteFile(fmt.Sprintf("%s/api_op_%s.go", outputDir, apiDefine.Name), apiFileContent, 0644)

	return nil
}

type ApiDefine struct {
	Name            string
	Description     string
	APILink         string
	Method          ApiMethod
	Path            string
	RequestPayload  []ApiPayload
	ResponsePayload []ApiPayload
}

func NewApiDefine(path string, api ApiInfo) *ApiDefine {
	name := genApiName(path, api.Name)
	description := api.Description
	apiLink := fmt.Sprintf("https://pve.proxmox.com/pve-docs/api-viewer/#%s", path)
	method := api.Method

	var requstTagName string
	switch method {
	case ApiMethodGet:
		requstTagName = "query"
	case ApiMethodPost, ApiMethodPut, ApiMethodDelete:
		requstTagName = "json"
	}

	var requestPayload, responsePayload []ApiPayload

	// for debug api
	// if path == "/nodes/{node}/apt/repositories" && method == ApiMethodGet {
	// 	fmt.Println("x")
	// }

	requestPayload = genApiPayload(name+"Request", requstTagName, api.Parameters)
	responsePayload = genApiPayload(name+"Response", "json", api.Returns)

	apiDefine := ApiDefine{
		Name:            name,
		Description:     description,
		APILink:         apiLink,
		Method:          method,
		Path:            path,
		RequestPayload:  requestPayload,
		ResponsePayload: responsePayload,
	}

	return &apiDefine

}

func genApiName(path, name string) string {

	var result string

	for _, p := range strings.Split(path, "/") {
		if len(p) > 0 && p[0] != '{' {
			result += strcase.ToCamel(p)
		}
	}

	result += strcase.ToCamel(name)

	return result
}

type ApiPayload struct {
	Name    string
	Comment string
	Type    string
	Fields  []ApiPayloadField
}

type ApiPayloadField struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

func genApiPayload(name string, tagName string, params *JsonSchema) []ApiPayload {
	if params == nil {
		return nil
	}
	result := []ApiPayload{}

	payload := ApiPayload{
		Name:    name,
		Comment: params.Description,
	}

	// default type is object
	if params.Type == "" {
		params.Type = JsonSchemaTypeObject
	}

	switch params.Type {
	case JsonSchemaTypeArrary:

		if params.Items == nil {
			payload.Type = "[]interface{}"
		} else {
			switch params.Items.Type {
			case JsonSchemaTypeArrary:
				payload.Type = "[]" + name + "Item"
				subPayloads := genApiPayload(name+"Item"+"Item", tagName, params.Items)
				result = append(result, subPayloads...)
			case JsonSchemaTypeObject:
				payload.Type = "[]" + name + "Item"
				subPayloads := genApiPayload(name+"Item", tagName, params.Items)
				result = append(result, subPayloads...)
			case JsonSchemaTypeBoolean:
				payload.Type = "[]bool"

			case JsonSchemaTypeNumber:
				payload.Type = "[]float64"

			case JsonSchemaTypeInteger:
				payload.Type = "[]int64"

			case JsonSchemaTypeString:
				payload.Type = "[]string"

			case JsonSchemaTypeAny:
				payload.Type = "[]interface{}"

			case JsonSchemaTypeNull:
				payload.Type = "[]struct{}"
			}
		}

	case JsonSchemaTypeObject:

		if params.Properties != nil {
			payload.Type = "struct"
			fields, subPayloads := genFields(name, tagName, params.Properties)
			payload.Fields = append(payload.Fields, fields...)
			result = append(result, subPayloads...)
		}

		if params.AdditionalProperties.Schema != nil &&
			params.AdditionalProperties.Schema.Type == JsonSchemaTypeObject {
			payload.Type = "struct"
			fields, subPayloads := genFields(name, tagName, params.AdditionalProperties.Schema.Properties)
			payload.Fields = append(payload.Fields, fields...)
			result = append(result, subPayloads...)
		}

		if params.Properties == nil && params.AdditionalProperties.Schema == nil {
			payload.Type = "interface{}"
		}

	case JsonSchemaTypeBoolean:
		payload.Type = "bool"

	case JsonSchemaTypeNumber:
		payload.Type = "float64"

	case JsonSchemaTypeInteger:
		payload.Type = "int64"

	case JsonSchemaTypeString:
		payload.Type = "string"

	case JsonSchemaTypeAny:
		payload.Type = "interface{}"

	case JsonSchemaTypeNull:
		payload.Type = "struct{}"

	}

	// sort payload fields
	sort.Slice(payload.Fields, func(i, j int) bool {
		return strings.Compare(payload.Fields[i].Name, payload.Fields[j].Name) == -1
	})

	result = append(result, payload)

	// sort payload list
	sort.Slice(result, func(i, j int) bool {
		return strings.Compare(result[i].Name, result[j].Name) == -1
	})

	return result
}

func genFields(name, tagName string, properties map[string]JsonSchema) ([]ApiPayloadField, []ApiPayload) {
	fields := []ApiPayloadField{}
	subPayloads := []ApiPayload{}

	for fieldName, param := range properties {
		field := ApiPayloadField{
			Name:    fieldName,
			Tag:     tagName,
			Comment: param.Description,
		}
		// default field type
		if param.Type == "" {
			param.Type = JsonSchemaTypeString
		}
		switch param.Type {
		case JsonSchemaTypeArrary:
			if param.Items == nil {
				field.Type = "[]interface{}"
			} else {
				field.Type = "[]" + name + strcase.ToCamel(fieldName) + "Item"
				subP := genApiPayload(name+strcase.ToCamel(fieldName)+"Item", tagName, param.Items)
				subPayloads = append(subPayloads, subP...)
			}
		case JsonSchemaTypeObject:
			field.Type = name + strcase.ToCamel(fieldName)
			subP := genApiPayload(name+strcase.ToCamel(fieldName), tagName, &param)
			subPayloads = append(subPayloads, subP...)
		case JsonSchemaTypeBoolean:
			field.Type = "bool"

		case JsonSchemaTypeNumber:
			field.Type = "float64"

		case JsonSchemaTypeInteger:
			field.Type = "int64"

		case JsonSchemaTypeString:
			field.Type = "string"

		case JsonSchemaTypeAny:
			field.Type = "interface{}"

		case JsonSchemaTypeNull:
			field.Type = "struct"
		}
		if param.Optional {
			field.Type = "*" + field.Type
		}

		fields = append(fields, field)
	}

	return fields, subPayloads

}
