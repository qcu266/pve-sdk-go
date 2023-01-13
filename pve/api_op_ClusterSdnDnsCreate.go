// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterSdnDnsCreateRequest struct {
	Dns           string `json:"dns,omitempty"`           // The SDN dns object identifier.
	Key           string `json:"key,omitempty"`           //
	Reversemaskv6 *int64 `json:"reversemaskv6,omitempty"` //
	Reversev6Mask *int64 `json:"reversev6mask,omitempty"` //
	Ttl           *int64 `json:"ttl,omitempty"`           //
	Type          string `json:"type,omitempty"`          // Plugin type.
	Url           string `json:"url,omitempty"`           //
}

type ClusterSdnDnsCreateResponse struct{}

// Create a new sdn dns object.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/sdn/dns
func (c *Client) ClusterSdnDnsCreate(ctx context.Context, request *ClusterSdnDnsCreateRequest) (*ClusterSdnDnsCreateResponse, error) {

	method := "POST"
	path := "/cluster/sdn/dns"

	var response ClusterSdnDnsCreateResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}