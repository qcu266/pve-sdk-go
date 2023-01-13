// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuFirewallIpsetReadIpRequest struct {
	Cidr string `query:"cidr,omitempty"` // Network/IP specification in CIDR format.
	Name string `query:"name,omitempty"` // IP set name.
	Node string `query:"node,omitempty"` // The cluster node name.
	Vmid int64  `query:"vmid,omitempty"` // The (unique) ID of the VM.
}

type NodesQemuFirewallIpsetReadIpResponse interface{}

// Read IP or Network settings from IPSet.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}
func (c *Client) NodesQemuFirewallIpsetReadIp(ctx context.Context, request *NodesQemuFirewallIpsetReadIpRequest) (*NodesQemuFirewallIpsetReadIpResponse, error) {

	method := "GET"
	path := "/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}"

	var response NodesQemuFirewallIpsetReadIpResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
