// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessTicketCreateTicketRequest struct {
	NewFormat    *bool   `json:"new-format,omitempty"`    // With webauthn the format of half-authenticated tickts changed. New clients should pass 1 here and not worry about the old format. The old format is deprecated and will be retired with PVE-8.0
	Otp          *string `json:"otp,omitempty"`           // One-time password for Two-factor authentication.
	Password     string  `json:"password,omitempty"`      // The secret password. This can also be a valid ticket.
	Path         *string `json:"path,omitempty"`          // Verify ticket, and check if user have access 'privs' on 'path'
	Privs        *string `json:"privs,omitempty"`         // Verify ticket, and check if user have access 'privs' on 'path'
	Realm        *string `json:"realm,omitempty"`         // You can optionally pass the realm using this parameter. Normally the realm is simply added to the username <username>@<relam>.
	TfaChallenge *string `json:"tfa-challenge,omitempty"` // The signed TFA challenge string the user wants to respond to.
	Username     string  `json:"username,omitempty"`      // User name
}

type AccessTicketCreateTicketResponse struct {
	CSRFPreventionToken *string `json:"CSRFPreventionToken,omitempty"` //
	Clustername         *string `json:"clustername,omitempty"`         //
	Ticket              *string `json:"ticket,omitempty"`              //
	Username            string  `json:"username,omitempty"`            //
}

// Create or verify authentication ticket.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/ticket
func (c *Client) AccessTicketCreateTicket(ctx context.Context, request *AccessTicketCreateTicketRequest) (*AccessTicketCreateTicketResponse, error) {

	method := "POST"
	path := "/access/ticket"

	var response AccessTicketCreateTicketResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}