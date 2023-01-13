// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesSdnZonesContentIndexRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
	Zone string `query:"zone,omitempty"` // The SDN zone object identifier.
}

type NodesSdnZonesContentIndexResponseItem struct {
	Status    *string `json:"status,omitempty"`    // Status.
	Statusmsg *string `json:"statusmsg,omitempty"` // Status details
	Vnet      string  `json:"vnet,omitempty"`      // Vnet identifier.
}

type NodesSdnZonesContentIndexResponse []NodesSdnZonesContentIndexResponseItem

// List zone content.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/sdn/zones/{zone}/content
func (c *Client) NodesSdnZonesContentIndex(ctx context.Context, request *NodesSdnZonesContentIndexRequest) (*NodesSdnZonesContentIndexResponse, error) {

	method := "GET"
	path := "/nodes/{node}/sdn/zones/{zone}/content"

	var response NodesSdnZonesContentIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
