// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessRolesCreateRoleRequest struct {
	Privs  *string `json:"privs,omitempty"`  //
	Roleid string  `json:"roleid,omitempty"` //
}

type AccessRolesCreateRoleResponse struct{}

// Create new role.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/roles
func (c *Client) AccessRolesCreateRole(ctx context.Context, request *AccessRolesCreateRoleRequest) (*AccessRolesCreateRoleResponse, error) {

	method := "POST"
	path := "/access/roles"

	var response AccessRolesCreateRoleResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
