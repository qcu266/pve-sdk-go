// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcStatusShutdownVmShutdownRequest struct {
	ForceStop *bool  `json:"forceStop,omitempty"` // Make sure the Container stops.
	Node      string `json:"node,omitempty"`      // The cluster node name.
	Timeout   *int64 `json:"timeout,omitempty"`   // Wait maximal timeout seconds.
	Vmid      int64  `json:"vmid,omitempty"`      // The (unique) ID of the VM.
}

type NodesLxcStatusShutdownVmShutdownResponse string

// Shutdown the container. This will trigger a clean shutdown of the container, see lxc-stop(1) for details.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/status/shutdown
func (c *Client) NodesLxcStatusShutdownVmShutdown(ctx context.Context, request *NodesLxcStatusShutdownVmShutdownRequest) (*NodesLxcStatusShutdownVmShutdownResponse, error) {

	method := "POST"
	path := "/nodes/{node}/lxc/{vmid}/status/shutdown"

	var response NodesLxcStatusShutdownVmShutdownResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
