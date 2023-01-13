// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterReplicationDeleteRequest struct {
	Force *bool  `json:"force,omitempty"` // Will remove the jobconfig entry, but will not cleanup.
	Id    string `json:"id,omitempty"`    // Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
	Keep  *bool  `json:"keep,omitempty"`  // Keep replicated data at target (do not remove).
}

type ClusterReplicationDeleteResponse struct{}

// Mark replication job for removal.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/replication/{id}
func (c *Client) ClusterReplicationDelete(ctx context.Context, request *ClusterReplicationDeleteRequest) (*ClusterReplicationDeleteResponse, error) {

	method := "DELETE"
	path := "/cluster/replication/{id}"

	var response ClusterReplicationDeleteResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}