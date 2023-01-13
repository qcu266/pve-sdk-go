// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesNetstatNetstatRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesNetstatNetstatResponse []NodesNetstatNetstatResponseItem

type NodesNetstatNetstatResponseItem struct {
}

// Read tap/vm network device interface counters
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/netstat
func (c *Client) NodesNetstatNetstat(ctx context.Context, request *NodesNetstatNetstatRequest) (*NodesNetstatNetstatResponse, error) {

	method := "GET"
	path := "/nodes/{node}/netstat"

	var response NodesNetstatNetstatResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
