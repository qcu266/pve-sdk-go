// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessGroupsDeleteGroupRequest struct {
	Groupid string `json:"groupid,omitempty"` //
}

type AccessGroupsDeleteGroupResponse struct{}

// Delete group.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/groups/{groupid}
func (c *Client) AccessGroupsDeleteGroup(ctx context.Context, request *AccessGroupsDeleteGroupRequest) (*AccessGroupsDeleteGroupResponse, error) {

	method := "DELETE"
	path := "/access/groups/{groupid}"

	var response AccessGroupsDeleteGroupResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}