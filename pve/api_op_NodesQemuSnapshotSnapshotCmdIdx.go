// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuSnapshotSnapshotCmdIdxRequest struct {
	Node     string `query:"node,omitempty"`     // The cluster node name.
	Snapname string `query:"snapname,omitempty"` // The name of the snapshot.
	Vmid     int64  `query:"vmid,omitempty"`     // The (unique) ID of the VM.
}

type NodesQemuSnapshotSnapshotCmdIdxResponse []NodesQemuSnapshotSnapshotCmdIdxResponseItem

type NodesQemuSnapshotSnapshotCmdIdxResponseItem struct {
}

// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/snapshot/{snapname}
func (c *Client) NodesQemuSnapshotSnapshotCmdIdx(ctx context.Context, request *NodesQemuSnapshotSnapshotCmdIdxRequest) (*NodesQemuSnapshotSnapshotCmdIdxResponse, error) {

	method := "GET"
	path := "/nodes/{node}/qemu/{vmid}/snapshot/{snapname}"

	var response NodesQemuSnapshotSnapshotCmdIdxResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
