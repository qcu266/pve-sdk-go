// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuSendkeyVmSendkeyRequest struct {
	Key      string `json:"key,omitempty"`      // The key (qemu monitor encoding).
	Node     string `json:"node,omitempty"`     // The cluster node name.
	Skiplock *bool  `json:"skiplock,omitempty"` // Ignore locks - only root is allowed to use this option.
	Vmid     int64  `json:"vmid,omitempty"`     // The (unique) ID of the VM.
}

type NodesQemuSendkeyVmSendkeyResponse struct{}

// Send key event to virtual machine.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/sendkey
func (c *Client) NodesQemuSendkeyVmSendkey(ctx context.Context, request *NodesQemuSendkeyVmSendkeyRequest) (*NodesQemuSendkeyVmSendkeyResponse, error) {

	method := "PUT"
	path := "/nodes/{node}/qemu/{vmid}/sendkey"

	var response NodesQemuSendkeyVmSendkeyResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
