// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesStartallStartallRequest struct {
	Force *bool   `json:"force,omitempty"` // Issue start command even if virtual guest have 'onboot' not set or set to off.
	Node  string  `json:"node,omitempty"`  // The cluster node name.
	Vms   *string `json:"vms,omitempty"`   // Only consider guests from this comma separated list of VMIDs.
}

type NodesStartallStartallResponse string

// Start all VMs and containers located on this node (by default only those with onboot=1).
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/startall
func (c *Client) NodesStartallStartall(ctx context.Context, request *NodesStartallStartallRequest) (*NodesStartallStartallResponse, error) {

	method := "POST"
	path := "/nodes/{node}/startall"

	var response NodesStartallStartallResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}