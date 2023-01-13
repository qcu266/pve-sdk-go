// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesTermproxyTermproxyRequest struct {
	Cmd     *string `json:"cmd,omitempty"`      // Run specific command or default to login.
	CmdOpts *string `json:"cmd-opts,omitempty"` // Add parameters to a command. Encoded as null terminated strings.
	Node    string  `json:"node,omitempty"`     // The cluster node name.
}

type NodesTermproxyTermproxyResponse struct {
	Port   int64  `json:"port,omitempty"`   //
	Ticket string `json:"ticket,omitempty"` //
	Upid   string `json:"upid,omitempty"`   //
	User   string `json:"user,omitempty"`   //
}

// Creates a VNC Shell proxy.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/termproxy
func (c *Client) NodesTermproxyTermproxy(ctx context.Context, request *NodesTermproxyTermproxyRequest) (*NodesTermproxyTermproxyResponse, error) {

	method := "POST"
	path := "/nodes/{node}/termproxy"

	var response NodesTermproxyTermproxyResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}