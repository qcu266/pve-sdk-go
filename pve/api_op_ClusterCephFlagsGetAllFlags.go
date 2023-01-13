// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterCephFlagsGetAllFlagsRequest interface{}

type ClusterCephFlagsGetAllFlagsResponseItem struct {
	Name string `json:"name,omitempty"` // Flag name.
}

type ClusterCephFlagsGetAllFlagsResponse []ClusterCephFlagsGetAllFlagsResponseItem

// get the status of all ceph flags
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/ceph/flags
func (c *Client) ClusterCephFlagsGetAllFlags(ctx context.Context, request *ClusterCephFlagsGetAllFlagsRequest) (*ClusterCephFlagsGetAllFlagsResponse, error) {

	method := "GET"
	path := "/cluster/ceph/flags"

	var response ClusterCephFlagsGetAllFlagsResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
