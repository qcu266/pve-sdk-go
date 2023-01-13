// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessGroupsIndexRequest interface{}

type AccessGroupsIndexResponse []AccessGroupsIndexResponseItem

type AccessGroupsIndexResponseItem struct {
	Comment *string `json:"comment,omitempty"` //
	Groupid string  `json:"groupid,omitempty"` //
	Users   *string `json:"users,omitempty"`   // list of users which form this group
}

// Group index.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/groups
func (c *Client) AccessGroupsIndex(ctx context.Context, request *AccessGroupsIndexRequest) (*AccessGroupsIndexResponse, error) {

	method := "GET"
	path := "/access/groups"

	var response AccessGroupsIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
