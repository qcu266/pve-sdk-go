// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesAptRepositoriesChangeRepositoryRequest struct {
	Digest  *string `json:"digest,omitempty"`  // Digest to detect modifications.
	Enabled *bool   `json:"enabled,omitempty"` // Whether the repository should be enabled or not.
	Index   int64   `json:"index,omitempty"`   // Index within the file (starting from 0).
	Node    string  `json:"node,omitempty"`    // The cluster node name.
	Path    string  `json:"path,omitempty"`    // Path to the containing file.
}

type NodesAptRepositoriesChangeRepositoryResponse struct{}

// Change the properties of a repository. Currently only allows enabling/disabling.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/apt/repositories
func (c *Client) NodesAptRepositoriesChangeRepository(ctx context.Context, request *NodesAptRepositoriesChangeRepositoryRequest) (*NodesAptRepositoriesChangeRepositoryResponse, error) {

	method := "POST"
	path := "/nodes/{node}/apt/repositories"

	var response NodesAptRepositoriesChangeRepositoryResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
