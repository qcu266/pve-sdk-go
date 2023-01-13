// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcTemplateTemplateRequest struct {
	Node string `json:"node,omitempty"` // The cluster node name.
	Vmid int64  `json:"vmid,omitempty"` // The (unique) ID of the VM.
}

type NodesLxcTemplateTemplateResponse struct{}

// Create a Template.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/template
func (c *Client) NodesLxcTemplateTemplate(ctx context.Context, request *NodesLxcTemplateTemplateRequest) (*NodesLxcTemplateTemplateResponse, error) {

	method := "POST"
	path := "/nodes/{node}/lxc/{vmid}/template"

	var response NodesLxcTemplateTemplateResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
