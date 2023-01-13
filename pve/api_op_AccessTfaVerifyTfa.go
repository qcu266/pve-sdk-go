// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessTfaVerifyTfaRequest struct {
	Response string `json:"response,omitempty"` // The response to the current authentication challenge.
}

type AccessTfaVerifyTfaResponse struct {
	Ticket string `json:"ticket,omitempty"` //
}

// Finish a u2f challenge.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/tfa
func (c *Client) AccessTfaVerifyTfa(ctx context.Context, request *AccessTfaVerifyTfaRequest) (*AccessTfaVerifyTfaResponse, error) {

	method := "POST"
	path := "/access/tfa"

	var response AccessTfaVerifyTfaResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
