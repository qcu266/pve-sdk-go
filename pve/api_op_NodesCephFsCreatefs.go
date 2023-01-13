// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesCephFsCreatefsRequest struct {
	AddStorage *bool   `json:"add-storage,omitempty"` // Configure the created CephFS as storage for this cluster.
	Name       *string `json:"name,omitempty"`        // The ceph filesystem name.
	Node       string  `json:"node,omitempty"`        // The cluster node name.
	PgNum      *int64  `json:"pg_num,omitempty"`      // Number of placement groups for the backing data pool. The metadata pool will use a quarter of this.
}

type NodesCephFsCreatefsResponse string

// Create a Ceph filesystem
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/ceph/fs/{name}
func (c *Client) NodesCephFsCreatefs(ctx context.Context, request *NodesCephFsCreatefsRequest) (*NodesCephFsCreatefsResponse, error) {

	method := "POST"
	path := "/nodes/{node}/ceph/fs/{name}"

	var response NodesCephFsCreatefsResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
