// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuMigrateMigrateVmPreconditionRequest struct {
	Node   string  `query:"node,omitempty"`   // The cluster node name.
	Target *string `query:"target,omitempty"` // Target node.
	Vmid   int64   `query:"vmid,omitempty"`   // The (unique) ID of the VM.
}

// List not allowed nodes with additional informations, only passed if VM is offline
type NodesQemuMigrateMigrateVmPreconditionResponseNotAllowedNodes interface{}

type NodesQemuMigrateMigrateVmPreconditionResponse struct {
	AllowedNodes    *[]interface{}                                                `json:"allowed_nodes,omitempty"`     // List nodes allowed for offline migration, only passed if VM is offline
	LocalDisks      []interface{}                                                 `json:"local_disks,omitempty"`       // List local disks including CD-Rom, unsused and not referenced disks
	LocalResources  []interface{}                                                 `json:"local_resources,omitempty"`   // List local resources e.g. pci, usb
	NotAllowedNodes *NodesQemuMigrateMigrateVmPreconditionResponseNotAllowedNodes `json:"not_allowed_nodes,omitempty"` // List not allowed nodes with additional informations, only passed if VM is offline
	Running         bool                                                          `json:"running,omitempty"`           //
}

// Get preconditions for migration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/migrate
func (c *Client) NodesQemuMigrateMigrateVmPrecondition(ctx context.Context, request *NodesQemuMigrateMigrateVmPreconditionRequest) (*NodesQemuMigrateMigrateVmPreconditionResponse, error) {

	method := "GET"
	path := "/nodes/{node}/qemu/{vmid}/migrate"

	var response NodesQemuMigrateMigrateVmPreconditionResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
