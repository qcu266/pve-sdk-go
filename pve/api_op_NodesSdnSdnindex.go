// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesSdnSdnindexRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesSdnSdnindexResponseItem struct {
}

type NodesSdnSdnindexResponse []NodesSdnSdnindexResponseItem

// SDN index.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/sdn
func (c *Client) NodesSdnSdnindex(ctx context.Context, request *NodesSdnSdnindexRequest) (*NodesSdnSdnindexResponse, error) {

	method := "GET"
	path := "/nodes/{node}/sdn"

	var response NodesSdnSdnindexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}