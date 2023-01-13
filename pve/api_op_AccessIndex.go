// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessIndexRequest interface{}

type AccessIndexResponse []AccessIndexResponseItem

type AccessIndexResponseItem struct {
	Subdir string `json:"subdir,omitempty"` //
}

// Directory index.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access
func (c *Client) AccessIndex(ctx context.Context, request *AccessIndexRequest) (*AccessIndexResponse, error) {

	method := "GET"
	path := "/access"

	var response AccessIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
