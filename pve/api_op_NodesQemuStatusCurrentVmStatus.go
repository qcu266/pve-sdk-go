// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuStatusCurrentVmStatusRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
	Vmid int64  `query:"vmid,omitempty"` // The (unique) ID of the VM.
}

// HA manager service status.
type NodesQemuStatusCurrentVmStatusResponseHa interface{}

type NodesQemuStatusCurrentVmStatusResponse struct {
	Agent          *bool                                    `json:"agent,omitempty"`           // Qemu GuestAgent enabled in config.
	Cpus           *float64                                 `json:"cpus,omitempty"`            // Maximum usable CPUs.
	Ha             NodesQemuStatusCurrentVmStatusResponseHa `json:"ha,omitempty"`              // HA manager service status.
	Lock           *string                                  `json:"lock,omitempty"`            // The current config lock, if any.
	Maxdisk        *int64                                   `json:"maxdisk,omitempty"`         // Root disk size in bytes.
	Maxmem         *int64                                   `json:"maxmem,omitempty"`          // Maximum memory in bytes.
	Name           *string                                  `json:"name,omitempty"`            // VM name.
	Pid            *int64                                   `json:"pid,omitempty"`             // PID of running qemu process.
	Qmpstatus      *string                                  `json:"qmpstatus,omitempty"`       // Qemu QMP agent status.
	RunningMachine *string                                  `json:"running-machine,omitempty"` // The currently running machine type (if running).
	RunningQemu    *string                                  `json:"running-qemu,omitempty"`    // The currently running QEMU version (if running).
	Spice          *bool                                    `json:"spice,omitempty"`           // Qemu VGA configuration supports spice.
	Status         string                                   `json:"status,omitempty"`          // Qemu process status.
	Tags           *string                                  `json:"tags,omitempty"`            // The current configured tags, if any
	Uptime         *int64                                   `json:"uptime,omitempty"`          // Uptime.
	Vmid           int64                                    `json:"vmid,omitempty"`            // The (unique) ID of the VM.
}

// Get virtual machine status.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/status/current
func (c *Client) NodesQemuStatusCurrentVmStatus(ctx context.Context, request *NodesQemuStatusCurrentVmStatusRequest) (*NodesQemuStatusCurrentVmStatusResponse, error) {

	method := "GET"
	path := "/nodes/{node}/qemu/{vmid}/status/current"

	var response NodesQemuStatusCurrentVmStatusResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}