// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessDomainsReadRequest struct {
	Realm string `query:"realm,omitempty"` // Authentication domain ID
}

type AccessDomainsReadResponse interface{}

// Get auth server configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/domains/{realm}
func (c *Client) AccessDomainsRead(ctx context.Context, request *AccessDomainsReadRequest) (*AccessDomainsReadResponse, error) {

	method := "GET"
	path := "/access/domains/{realm}"

	var response AccessDomainsReadResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
