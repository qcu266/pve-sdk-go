// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesNetworkIndexRequest struct {
	Node string  `query:"node,omitempty"` // The cluster node name.
	Type *string `query:"type,omitempty"` // Only list specific interface types.
}

type NodesNetworkIndexResponseItem struct {
}

type NodesNetworkIndexResponse []NodesNetworkIndexResponseItem

// List available networks
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/network
func (c *Client) NodesNetworkIndex(ctx context.Context, request *NodesNetworkIndexRequest) (*NodesNetworkIndexResponse, error) {

	method := "GET"
	path := "/nodes/{node}/network"

	var response NodesNetworkIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
