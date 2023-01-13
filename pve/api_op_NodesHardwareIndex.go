// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesHardwareIndexRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesHardwareIndexResponse []NodesHardwareIndexResponseItem

type NodesHardwareIndexResponseItem struct {
	Type string `json:"type,omitempty"` //
}

// Index of hardware types
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/hardware
func (c *Client) NodesHardwareIndex(ctx context.Context, request *NodesHardwareIndexRequest) (*NodesHardwareIndexResponse, error) {

	method := "GET"
	path := "/nodes/{node}/hardware"

	var response NodesHardwareIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
