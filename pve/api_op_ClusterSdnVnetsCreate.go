// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterSdnVnetsCreateRequest struct {
	Alias     *string `json:"alias,omitempty"`     // alias name of the vnet
	Tag       *int64  `json:"tag,omitempty"`       // vlan or vxlan id
	Type      *string `json:"type,omitempty"`      // Type
	Vlanaware *bool   `json:"vlanaware,omitempty"` // Allow vm VLANs to pass through this vnet.
	Vnet      string  `json:"vnet,omitempty"`      // The SDN vnet object identifier.
	Zone      string  `json:"zone,omitempty"`      // zone id
}

type ClusterSdnVnetsCreateResponse struct{}

// Create a new sdn vnet object.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/sdn/vnets
func (c *Client) ClusterSdnVnetsCreate(ctx context.Context, request *ClusterSdnVnetsCreateRequest) (*ClusterSdnVnetsCreateResponse, error) {

	method := "POST"
	path := "/cluster/sdn/vnets"

	var response ClusterSdnVnetsCreateResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
