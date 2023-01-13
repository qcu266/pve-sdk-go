// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesServicesIndexRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesServicesIndexResponseItem struct {
}

type NodesServicesIndexResponse []NodesServicesIndexResponseItem

// Service list.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/services
func (c *Client) NodesServicesIndex(ctx context.Context, request *NodesServicesIndexRequest) (*NodesServicesIndexResponse, error) {

	method := "GET"
	path := "/nodes/{node}/services"

	var response NodesServicesIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}