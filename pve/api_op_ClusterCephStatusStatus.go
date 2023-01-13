// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterCephStatusStatusRequest interface{}

type ClusterCephStatusStatusResponse interface{}

// Get ceph status.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/ceph/status
func (c *Client) ClusterCephStatusStatus(ctx context.Context, request *ClusterCephStatusStatusRequest) (*ClusterCephStatusStatusResponse, error) {

	method := "GET"
	path := "/cluster/ceph/status"

	var response ClusterCephStatusStatusResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
