// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterConfigNodesNodesRequest interface{}

type ClusterConfigNodesNodesResponse []ClusterConfigNodesNodesResponseItem

type ClusterConfigNodesNodesResponseItem struct {
	Node string `json:"node,omitempty"` //
}

// Corosync node list.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/config/nodes
func (c *Client) ClusterConfigNodesNodes(ctx context.Context, request *ClusterConfigNodesNodesRequest) (*ClusterConfigNodesNodesResponse, error) {

	method := "GET"
	path := "/cluster/config/nodes"

	var response ClusterConfigNodesNodesResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
