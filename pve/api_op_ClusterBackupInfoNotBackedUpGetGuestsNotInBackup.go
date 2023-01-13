// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterBackupInfoNotBackedUpGetGuestsNotInBackupRequest interface{}

type ClusterBackupInfoNotBackedUpGetGuestsNotInBackupResponseItem struct {
	Name *string `json:"name,omitempty"` // Name of the guest
	Type string  `json:"type,omitempty"` // Type of the guest.
	Vmid int64   `json:"vmid,omitempty"` // VMID of the guest.
}

// Contains the guest objects.
type ClusterBackupInfoNotBackedUpGetGuestsNotInBackupResponse []ClusterBackupInfoNotBackedUpGetGuestsNotInBackupResponseItem

// Shows all guests which are not covered by any backup job.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/backup-info/not-backed-up
func (c *Client) ClusterBackupInfoNotBackedUpGetGuestsNotInBackup(ctx context.Context, request *ClusterBackupInfoNotBackedUpGetGuestsNotInBackupRequest) (*ClusterBackupInfoNotBackedUpGetGuestsNotInBackupResponse, error) {

	method := "GET"
	path := "/cluster/backup-info/not-backed-up"

	var response ClusterBackupInfoNotBackedUpGetGuestsNotInBackupResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
