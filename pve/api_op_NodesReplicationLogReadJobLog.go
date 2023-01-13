// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesReplicationLogReadJobLogRequest struct {
	Id    string `query:"id,omitempty"`    // Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
	Limit *int64 `query:"limit,omitempty"` //
	Node  string `query:"node,omitempty"`  // The cluster node name.
	Start *int64 `query:"start,omitempty"` //
}

type NodesReplicationLogReadJobLogResponse []NodesReplicationLogReadJobLogResponseItem

type NodesReplicationLogReadJobLogResponseItem struct {
	N int64  `json:"n,omitempty"` // Line number
	T string `json:"t,omitempty"` // Line text
}

// Read replication job log.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/replication/{id}/log
func (c *Client) NodesReplicationLogReadJobLog(ctx context.Context, request *NodesReplicationLogReadJobLogRequest) (*NodesReplicationLogReadJobLogResponse, error) {

	method := "GET"
	path := "/nodes/{node}/replication/{id}/log"

	var response NodesReplicationLogReadJobLogResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
