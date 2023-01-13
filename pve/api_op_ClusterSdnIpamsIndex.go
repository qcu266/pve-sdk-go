// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterSdnIpamsIndexRequest struct {
	Type *string `query:"type,omitempty"` // Only list sdn ipams of specific type
}

type ClusterSdnIpamsIndexResponse []ClusterSdnIpamsIndexResponseItem

type ClusterSdnIpamsIndexResponseItem struct {
	Ipam string `json:"ipam,omitempty"` //
	Type string `json:"type,omitempty"` //
}

// SDN ipams index.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/sdn/ipams
func (c *Client) ClusterSdnIpamsIndex(ctx context.Context, request *ClusterSdnIpamsIndexRequest) (*ClusterSdnIpamsIndexResponse, error) {

	method := "GET"
	path := "/cluster/sdn/ipams"

	var response ClusterSdnIpamsIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
