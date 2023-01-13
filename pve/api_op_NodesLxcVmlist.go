// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcVmlistRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesLxcVmlistResponseItem struct {
	Cpus    *float64 `json:"cpus,omitempty"`    // Maximum usable CPUs.
	Lock    *string  `json:"lock,omitempty"`    // The current config lock, if any.
	Maxdisk *int64   `json:"maxdisk,omitempty"` // Root disk size in bytes.
	Maxmem  *int64   `json:"maxmem,omitempty"`  // Maximum memory in bytes.
	Maxswap *int64   `json:"maxswap,omitempty"` // Maximum SWAP memory in bytes.
	Name    *string  `json:"name,omitempty"`    // Container name.
	Status  string   `json:"status,omitempty"`  // LXC Container status.
	Tags    *string  `json:"tags,omitempty"`    // The current configured tags, if any.
	Uptime  *int64   `json:"uptime,omitempty"`  // Uptime.
	Vmid    int64    `json:"vmid,omitempty"`    // The (unique) ID of the VM.
}

type NodesLxcVmlistResponse []NodesLxcVmlistResponseItem

// LXC container index (per node).
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc
func (c *Client) NodesLxcVmlist(ctx context.Context, request *NodesLxcVmlistRequest) (*NodesLxcVmlistResponse, error) {

	method := "GET"
	path := "/nodes/{node}/lxc"

	var response NodesLxcVmlistResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
