// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesDisksZfsDeleteRequest struct {
	CleanupConfig *bool  `json:"cleanup-config,omitempty"` // Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
	CleanupDisks  *bool  `json:"cleanup-disks,omitempty"`  // Also wipe disks so they can be repurposed afterwards.
	Name          string `json:"name,omitempty"`           // The storage identifier.
	Node          string `json:"node,omitempty"`           // The cluster node name.
}

type NodesDisksZfsDeleteResponse string

// Destroy a ZFS pool.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/disks/zfs/{name}
func (c *Client) NodesDisksZfsDelete(ctx context.Context, request *NodesDisksZfsDeleteRequest) (*NodesDisksZfsDeleteResponse, error) {

	method := "DELETE"
	path := "/nodes/{node}/disks/zfs/{name}"

	var response NodesDisksZfsDeleteResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}