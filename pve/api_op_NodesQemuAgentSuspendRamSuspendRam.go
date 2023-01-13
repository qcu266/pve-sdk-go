// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuAgentSuspendRamSuspendRamRequest struct {
	Node string `json:"node,omitempty"` // The cluster node name.
	Vmid int64  `json:"vmid,omitempty"` // The (unique) ID of the VM.
}

// Returns an object with a single `result` property.
type NodesQemuAgentSuspendRamSuspendRamResponse interface{}

// Execute suspend-ram.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/agent/suspend-ram
func (c *Client) NodesQemuAgentSuspendRamSuspendRam(ctx context.Context, request *NodesQemuAgentSuspendRamSuspendRamRequest) (*NodesQemuAgentSuspendRamSuspendRamResponse, error) {

	method := "POST"
	path := "/nodes/{node}/qemu/{vmid}/agent/suspend-ram"

	var response NodesQemuAgentSuspendRamSuspendRamResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
