// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesAptRepositoriesAddRepositoryRequest struct {
	Digest *string `json:"digest,omitempty"` // Digest to detect modifications.
	Handle string  `json:"handle,omitempty"` // Handle that identifies a repository.
	Node   string  `json:"node,omitempty"`   // The cluster node name.
}

type NodesAptRepositoriesAddRepositoryResponse struct{}

// Add a standard repository to the configuration
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/apt/repositories
func (c *Client) NodesAptRepositoriesAddRepository(ctx context.Context, request *NodesAptRepositoriesAddRepositoryRequest) (*NodesAptRepositoriesAddRepositoryResponse, error) {

	method := "PUT"
	path := "/nodes/{node}/apt/repositories"

	var response NodesAptRepositoriesAddRepositoryResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}