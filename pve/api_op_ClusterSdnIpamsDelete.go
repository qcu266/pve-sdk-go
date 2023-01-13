// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterSdnIpamsDeleteRequest struct {
	Ipam string `json:"ipam,omitempty"` // The SDN ipam object identifier.
}

type ClusterSdnIpamsDeleteResponse struct{}

// Delete sdn ipam object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/sdn/ipams/{ipam}
func (c *Client) ClusterSdnIpamsDelete(ctx context.Context, request *ClusterSdnIpamsDeleteRequest) (*ClusterSdnIpamsDeleteResponse, error) {

	method := "DELETE"
	path := "/cluster/sdn/ipams/{ipam}"

	var response ClusterSdnIpamsDeleteResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}