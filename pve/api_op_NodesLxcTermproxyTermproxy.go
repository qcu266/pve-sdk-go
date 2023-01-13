// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcTermproxyTermproxyRequest struct {
	Node string `json:"node,omitempty"` // The cluster node name.
	Vmid int64  `json:"vmid,omitempty"` // The (unique) ID of the VM.
}

type NodesLxcTermproxyTermproxyResponse struct {
	Port   int64  `json:"port,omitempty"`   //
	Ticket string `json:"ticket,omitempty"` //
	Upid   string `json:"upid,omitempty"`   //
	User   string `json:"user,omitempty"`   //
}

// Creates a TCP proxy connection.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/termproxy
func (c *Client) NodesLxcTermproxyTermproxy(ctx context.Context, request *NodesLxcTermproxyTermproxyRequest) (*NodesLxcTermproxyTermproxyResponse, error) {

	method := "POST"
	path := "/nodes/{node}/lxc/{vmid}/termproxy"

	var response NodesLxcTermproxyTermproxyResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
