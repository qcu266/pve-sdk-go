// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessAclUpdateAclRequest struct {
	Delete    *bool   `json:"delete,omitempty"`    // Remove permissions (instead of adding it).
	Groups    *string `json:"groups,omitempty"`    // List of groups.
	Path      string  `json:"path,omitempty"`      // Access control path
	Propagate *bool   `json:"propagate,omitempty"` // Allow to propagate (inherit) permissions.
	Roles     string  `json:"roles,omitempty"`     // List of roles.
	Tokens    *string `json:"tokens,omitempty"`    // List of API tokens.
	Users     *string `json:"users,omitempty"`     // List of users.
}

type AccessAclUpdateAclResponse struct{}

// Update Access Control List (add or remove permissions).
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/acl
func (c *Client) AccessAclUpdateAcl(ctx context.Context, request *AccessAclUpdateAclRequest) (*AccessAclUpdateAclResponse, error) {

	method := "PUT"
	path := "/access/acl"

	var response AccessAclUpdateAclResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
