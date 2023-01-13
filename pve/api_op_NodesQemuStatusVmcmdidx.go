// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuStatusVmcmdidxRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
	Vmid int64  `query:"vmid,omitempty"` // The (unique) ID of the VM.
}

type NodesQemuStatusVmcmdidxResponseItem struct {
	Subdir string `json:"subdir,omitempty"` //
}

type NodesQemuStatusVmcmdidxResponse []NodesQemuStatusVmcmdidxResponseItem

// Directory index
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/status
func (c *Client) NodesQemuStatusVmcmdidx(ctx context.Context, request *NodesQemuStatusVmcmdidxRequest) (*NodesQemuStatusVmcmdidxResponse, error) {

	method := "GET"
	path := "/nodes/{node}/qemu/{vmid}/status"

	var response NodesQemuStatusVmcmdidxResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}