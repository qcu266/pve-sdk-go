// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesCertificatesAcmeIndexRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesCertificatesAcmeIndexResponse []NodesCertificatesAcmeIndexResponseItem

type NodesCertificatesAcmeIndexResponseItem struct {
}

// ACME index.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/certificates/acme
func (c *Client) NodesCertificatesAcmeIndex(ctx context.Context, request *NodesCertificatesAcmeIndexRequest) (*NodesCertificatesAcmeIndexResponse, error) {

	method := "GET"
	path := "/nodes/{node}/certificates/acme"

	var response NodesCertificatesAcmeIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
