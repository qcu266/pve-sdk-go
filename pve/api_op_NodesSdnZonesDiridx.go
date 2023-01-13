// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesSdnZonesDiridxRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
	Zone string `query:"zone,omitempty"` // The SDN zone object identifier.
}

type NodesSdnZonesDiridxResponseItem struct {
	Subdir string `json:"subdir,omitempty"` //
}

type NodesSdnZonesDiridxResponse []NodesSdnZonesDiridxResponseItem

// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/sdn/zones/{zone}
func (c *Client) NodesSdnZonesDiridx(ctx context.Context, request *NodesSdnZonesDiridxRequest) (*NodesSdnZonesDiridxResponse, error) {

	method := "GET"
	path := "/nodes/{node}/sdn/zones/{zone}"

	var response NodesSdnZonesDiridxResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
