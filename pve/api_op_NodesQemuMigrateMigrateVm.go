// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuMigrateMigrateVmRequest struct {
	Bwlimit          *int64  `json:"bwlimit,omitempty"`           // Override I/O bandwidth limit (in KiB/s).
	Force            *bool   `json:"force,omitempty"`             // Allow to migrate VMs which use local devices. Only root may use this option.
	MigrationNetwork *string `json:"migration_network,omitempty"` // CIDR of the (sub) network that is used for migration.
	MigrationType    *string `json:"migration_type,omitempty"`    // Migration traffic is encrypted using an SSH tunnel by default. On secure, completely private networks this can be disabled to increase performance.
	Node             string  `json:"node,omitempty"`              // The cluster node name.
	Online           *bool   `json:"online,omitempty"`            // Use online/live migration if VM is running. Ignored if VM is stopped.
	Target           string  `json:"target,omitempty"`            // Target node.
	Targetstorage    *string `json:"targetstorage,omitempty"`     // Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage. Providing the special value '1' will map each source storage to itself.
	Vmid             int64   `json:"vmid,omitempty"`              // The (unique) ID of the VM.
	WithLocalDisks   *bool   `json:"with-local-disks,omitempty"`  // Enable live storage migration for local disk
}

// the task ID.
type NodesQemuMigrateMigrateVmResponse string

// Migrate virtual machine. Creates a new migration task.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/migrate
func (c *Client) NodesQemuMigrateMigrateVm(ctx context.Context, request *NodesQemuMigrateMigrateVmRequest) (*NodesQemuMigrateMigrateVmResponse, error) {

	method := "POST"
	path := "/nodes/{node}/qemu/{vmid}/migrate"

	var response NodesQemuMigrateMigrateVmResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
