// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcRrddataRrddataRequest struct {
	Cf        *string `query:"cf,omitempty"`        // The RRD consolidation function
	Node      string  `query:"node,omitempty"`      // The cluster node name.
	Timeframe string  `query:"timeframe,omitempty"` // Specify the time frame you are interested in.
	Vmid      int64   `query:"vmid,omitempty"`      // The (unique) ID of the VM.
}

type NodesLxcRrddataRrddataResponse []NodesLxcRrddataRrddataResponseItem

type NodesLxcRrddataRrddataResponseItem struct {
}

// Read VM RRD statistics
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/rrddata
func (c *Client) NodesLxcRrddataRrddata(ctx context.Context, request *NodesLxcRrddataRrddataRequest) (*NodesLxcRrddataRrddataResponse, error) {

	method := "GET"
	path := "/nodes/{node}/lxc/{vmid}/rrddata"

	var response NodesLxcRrddataRrddataResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
