// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessUsersUpdateUserRequest struct {
	Append    *bool   `json:"append,omitempty"`    //
	Comment   *string `json:"comment,omitempty"`   //
	Email     *string `json:"email,omitempty"`     //
	Enable    *bool   `json:"enable,omitempty"`    // Enable the account (default). You can set this to '0' to disable the account
	Expire    *int64  `json:"expire,omitempty"`    // Account expiration date (seconds since epoch). '0' means no expiration date.
	Firstname *string `json:"firstname,omitempty"` //
	Groups    *string `json:"groups,omitempty"`    //
	Keys      *string `json:"keys,omitempty"`      // Keys for two factor auth (yubico).
	Lastname  *string `json:"lastname,omitempty"`  //
	Userid    string  `json:"userid,omitempty"`    // User ID
}

type AccessUsersUpdateUserResponse struct{}

// Update user configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/users/{userid}
func (c *Client) AccessUsersUpdateUser(ctx context.Context, request *AccessUsersUpdateUserRequest) (*AccessUsersUpdateUserResponse, error) {

	method := "PUT"
	path := "/access/users/{userid}"

	var response AccessUsersUpdateUserResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
