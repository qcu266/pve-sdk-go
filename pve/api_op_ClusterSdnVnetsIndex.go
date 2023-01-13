// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterSdnVnetsIndexRequest struct {
	Pending *bool `query:"pending,omitempty"` // Display pending config.
	Running *bool `query:"running,omitempty"` // Display running config.
}

type ClusterSdnVnetsIndexResponse []ClusterSdnVnetsIndexResponseItem

type ClusterSdnVnetsIndexResponseItem struct {
}

// SDN vnets index.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/sdn/vnets
func (c *Client) ClusterSdnVnetsIndex(ctx context.Context, request *ClusterSdnVnetsIndexRequest) (*ClusterSdnVnetsIndexResponse, error) {

	method := "GET"
	path := "/cluster/sdn/vnets"

	var response ClusterSdnVnetsIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
