// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterLogLogRequest struct {
	Max *int64 `query:"max,omitempty"` // Maximum number of entries.
}

type ClusterLogLogResponse []ClusterLogLogResponseItem

type ClusterLogLogResponseItem struct {
}

// Read cluster log
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/log
func (c *Client) ClusterLogLog(ctx context.Context, request *ClusterLogLogRequest) (*ClusterLogLogResponse, error) {

	method := "GET"
	path := "/cluster/log"

	var response ClusterLogLogResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
