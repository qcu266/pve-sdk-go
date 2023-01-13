// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuRemoteMigrateRemoteMigrateVmRequest struct {
	Bwlimit        *int64 `json:"bwlimit,omitempty"`         // Override I/O bandwidth limit (in KiB/s).
	Delete         *bool  `json:"delete,omitempty"`          // Delete the original VM and related data after successful migration. By default the original VM is kept on the source cluster in a stopped state.
	Node           string `json:"node,omitempty"`            // The cluster node name.
	Online         *bool  `json:"online,omitempty"`          // Use online/live migration if VM is running. Ignored if VM is stopped.
	TargetBridge   string `json:"target-bridge,omitempty"`   // Mapping from source to target bridges. Providing only a single bridge ID maps all source bridges to that bridge. Providing the special value '1' will map each source bridge to itself.
	TargetEndpoint string `json:"target-endpoint,omitempty"` // Remote target endpoint
	TargetStorage  string `json:"target-storage,omitempty"`  // Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage. Providing the special value '1' will map each source storage to itself.
	TargetVmid     *int64 `json:"target-vmid,omitempty"`     // The (unique) ID of the VM.
	Vmid           int64  `json:"vmid,omitempty"`            // The (unique) ID of the VM.
}

// the task ID.
type NodesQemuRemoteMigrateRemoteMigrateVmResponse string

// Migrate virtual machine to a remote cluster. Creates a new migration task. EXPERIMENTAL feature!
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/remote_migrate
func (c *Client) NodesQemuRemoteMigrateRemoteMigrateVm(ctx context.Context, request *NodesQemuRemoteMigrateRemoteMigrateVmRequest) (*NodesQemuRemoteMigrateRemoteMigrateVmResponse, error) {

	method := "POST"
	path := "/nodes/{node}/qemu/{vmid}/remote_migrate"

	var response NodesQemuRemoteMigrateRemoteMigrateVmResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
