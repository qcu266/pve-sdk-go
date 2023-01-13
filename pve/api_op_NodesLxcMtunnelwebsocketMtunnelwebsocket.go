// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcMtunnelwebsocketMtunnelwebsocketRequest struct {
	Node   string `query:"node,omitempty"`   // The cluster node name.
	Socket string `query:"socket,omitempty"` // unix socket to forward to
	Ticket string `query:"ticket,omitempty"` // ticket return by initial 'mtunnel' API call, or retrieved via 'ticket' tunnel command
	Vmid   int64  `query:"vmid,omitempty"`   // The (unique) ID of the VM.
}

type NodesLxcMtunnelwebsocketMtunnelwebsocketResponse struct {
	Port   *string `json:"port,omitempty"`   //
	Socket *string `json:"socket,omitempty"` //
}

// Migration tunnel endpoint for websocket upgrade - only for internal use by VM migration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/mtunnelwebsocket
func (c *Client) NodesLxcMtunnelwebsocketMtunnelwebsocket(ctx context.Context, request *NodesLxcMtunnelwebsocketMtunnelwebsocketRequest) (*NodesLxcMtunnelwebsocketMtunnelwebsocketResponse, error) {

	method := "GET"
	path := "/nodes/{node}/lxc/{vmid}/mtunnelwebsocket"

	var response NodesLxcMtunnelwebsocketMtunnelwebsocketResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
