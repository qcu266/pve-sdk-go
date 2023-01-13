// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

{{- range .RequestPayload }}
// {{ .Comment | cleanLF }}
type {{ .Name | toCamel }} {{ .Type }} {{ if eq .Type "struct" }}{
{{- range .Fields  }}
	{{ .Name | toCamel }} {{ .Type }} `{{ .Tag }}:"{{ .Name }},omitempty"` // {{ .Comment | cleanLF}}
{{- end }}
}
{{- end }}


{{ end }}

{{- range .ResponsePayload }}
// {{ .Comment | cleanLF }}
type {{ .Name | toCamel }} {{ .Type }} {{ if eq .Type "struct" }}{
{{- range .Fields  }}
	{{ .Name | toCamel }} {{ .Type }} `{{ .Tag }}:"{{ .Name }},omitempty"` // {{ .Comment | cleanLF}}
{{- end }}
}
{{- end }}


{{ end }}

// {{ .Description | cleanLF }}
// {{ .APILink | cleanLF }}
func (c *Client) {{ .Name }}(ctx context.Context, request *{{ .Name }}Request) (*{{ .Name }}Response, error) {

	method := "{{ .Method }}"
	path := "{{ .Path }}"

	var response {{ .Name }}Response

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
