// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuAgentSuspendHybridSuspendHybridRequest struct {
	Node string `json:"node,omitempty"` // The cluster node name.
	Vmid int64  `json:"vmid,omitempty"` // The (unique) ID of the VM.
}

// Returns an object with a single `result` property.
type NodesQemuAgentSuspendHybridSuspendHybridResponse interface{}

// Execute suspend-hybrid.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/agent/suspend-hybrid
func (c *Client) NodesQemuAgentSuspendHybridSuspendHybrid(ctx context.Context, request *NodesQemuAgentSuspendHybridSuspendHybridRequest) (*NodesQemuAgentSuspendHybridSuspendHybridResponse, error) {

	method := "POST"
	path := "/nodes/{node}/qemu/{vmid}/agent/suspend-hybrid"

	var response NodesQemuAgentSuspendHybridSuspendHybridResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
