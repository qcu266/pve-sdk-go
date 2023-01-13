// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuFirewallIndexRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
	Vmid int64  `query:"vmid,omitempty"` // The (unique) ID of the VM.
}

type NodesQemuFirewallIndexResponseItem struct {
}

type NodesQemuFirewallIndexResponse []NodesQemuFirewallIndexResponseItem

// Directory index.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/firewall
func (c *Client) NodesQemuFirewallIndex(ctx context.Context, request *NodesQemuFirewallIndexRequest) (*NodesQemuFirewallIndexResponse, error) {

	method := "GET"
	path := "/nodes/{node}/qemu/{vmid}/firewall"

	var response NodesQemuFirewallIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
