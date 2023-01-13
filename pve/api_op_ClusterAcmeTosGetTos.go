// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterAcmeTosGetTosRequest struct {
	Directory *string `query:"directory,omitempty"` // URL of ACME CA directory endpoint.
}

// ACME TermsOfService URL.
type ClusterAcmeTosGetTosResponse string

// Retrieve ACME TermsOfService URL from CA.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/acme/tos
func (c *Client) ClusterAcmeTosGetTos(ctx context.Context, request *ClusterAcmeTosGetTosRequest) (*ClusterAcmeTosGetTosResponse, error) {

	method := "GET"
	path := "/cluster/acme/tos"

	var response ClusterAcmeTosGetTosResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
