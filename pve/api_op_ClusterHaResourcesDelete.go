// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterHaResourcesDeleteRequest struct {
	Sid string `json:"sid,omitempty"` // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
}

type ClusterHaResourcesDeleteResponse struct{}

// Delete resource configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/ha/resources/{sid}
func (c *Client) ClusterHaResourcesDelete(ctx context.Context, request *ClusterHaResourcesDeleteRequest) (*ClusterHaResourcesDeleteResponse, error) {

	method := "DELETE"
	path := "/cluster/ha/resources/{sid}"

	var response ClusterHaResourcesDeleteResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}