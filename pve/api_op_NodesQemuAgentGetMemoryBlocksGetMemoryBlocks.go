// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuAgentGetMemoryBlocksGetMemoryBlocksRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
	Vmid int64  `query:"vmid,omitempty"` // The (unique) ID of the VM.
}

// Returns an object with a single `result` property.
type NodesQemuAgentGetMemoryBlocksGetMemoryBlocksResponse interface{}

// Execute get-memory-blocks.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/agent/get-memory-blocks
func (c *Client) NodesQemuAgentGetMemoryBlocksGetMemoryBlocks(ctx context.Context, request *NodesQemuAgentGetMemoryBlocksGetMemoryBlocksRequest) (*NodesQemuAgentGetMemoryBlocksGetMemoryBlocksResponse, error) {

	method := "GET"
	path := "/nodes/{node}/qemu/{vmid}/agent/get-memory-blocks"

	var response NodesQemuAgentGetMemoryBlocksGetMemoryBlocksResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}