// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterOptionsGetOptionsRequest interface{}

type ClusterOptionsGetOptionsResponse interface{}

// Get datacenter options. Without 'Sys.Audit' on '/' not all options are returned.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/options
func (c *Client) ClusterOptionsGetOptions(ctx context.Context, request *ClusterOptionsGetOptionsRequest) (*ClusterOptionsGetOptionsResponse, error) {

	method := "GET"
	path := "/cluster/options"

	var response ClusterOptionsGetOptionsResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}