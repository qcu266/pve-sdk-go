package gen

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type ApiMethod string

const (
	ApiMethodGet    = "GET"
	ApiMethodPost   = "POST"
	ApiMethodPut    = "PUT"
	ApiMethodDelete = "DELETE"
)

type JsonSchemaType string

const (
	JsonSchemaTypeArrary  = "array"
	JsonSchemaTypeObject  = "object"
	JsonSchemaTypeBoolean = "boolean"
	JsonSchemaTypeNumber  = "number"
	JsonSchemaTypeInteger = "integer"
	JsonSchemaTypeString  = "string"
	JsonSchemaTypeAny     = "any"
	JsonSchemaTypeNull    = "null"
)

type ApiSchema struct {
	Children []ApiSchema
	Info     map[ApiMethod]ApiInfo
	Leaf     int
	Path     string
	Text     string
}

type ApiInfo struct {
	Allowtoken  ZeroOneBool
	Description string
	Method      ApiMethod
	Name        string
	Parameters  *JsonSchema
	Returns     *JsonSchema
	Permissions interface{}
}

type JsonSchema struct {
	Type     JsonSchemaType
	Optional ZeroOneBool

	TypeText    string
	Description string
	Format      PropertiesFormat

	Properties           map[string]JsonSchema
	AdditionalProperties AdditionalProperties

	Items *JsonSchema
	Links []*Link `json:"links"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type PropertiesFormat struct {
	Object map[string]PropertiesFormatObject
	String string
}

type PropertiesFormatObject struct {
	Default     interface{}
	DefaultKey  interface{}
	Description string
	Type        JsonSchemaType
	Enum        []interface{}
	Optional    ZeroOneBool
}

func (p *PropertiesFormat) UnmarshalJSON(data []byte) error {

	err := json.Unmarshal(data, &p.String)
	if err != nil {
		return json.Unmarshal(data, &p.Object)
	}

	return nil
}

type ZeroOneBool bool

func (s *ZeroOneBool) UnmarshalJSON(data []byte) error {

	datax := data[:]
	if len(datax) > 0 && datax[0] == '"' {
		datax = datax[1:]
	}
	if len(datax) > 0 && datax[len(datax)-1] == '"' {
		datax = datax[:len(datax)-1]
	}

	result, err := strconv.ParseInt(string(datax), 10, 64)
	if err != nil {
		return fmt.Errorf("parse data \"%s\" to Int failed: %s", string(datax), err)
	}

	switch result {
	case 0:
		*s = ZeroOneBool(false)
	case 1:
		*s = ZeroOneBool(true)
	default:
		return errors.New("Invalid ZeroOne Bool")
	}

	return nil
}

type AdditionalProperties struct {
	Schema *JsonSchema
	Bool   ZeroOneBool
}

func (s *AdditionalProperties) UnmarshalJSON(data []byte) error {

	err := json.Unmarshal(data, &s.Bool)
	if err != nil {
		s.Schema = &JsonSchema{}
		return json.Unmarshal(data, s.Schema)
	}

	return nil
}
