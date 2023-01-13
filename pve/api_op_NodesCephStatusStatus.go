// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesCephStatusStatusRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesCephStatusStatusResponse interface{}

// Get ceph status.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/ceph/status
func (c *Client) NodesCephStatusStatus(ctx context.Context, request *NodesCephStatusStatusRequest) (*NodesCephStatusStatusResponse, error) {

	method := "GET"
	path := "/nodes/{node}/ceph/status"

	var response NodesCephStatusStatusResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
