// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuFirewallIpsetUpdateIpRequest struct {
	Cidr    string  `json:"cidr,omitempty"`    // Network/IP specification in CIDR format.
	Comment *string `json:"comment,omitempty"` //
	Digest  *string `json:"digest,omitempty"`  // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string  `json:"name,omitempty"`    // IP set name.
	Node    string  `json:"node,omitempty"`    // The cluster node name.
	Nomatch *bool   `json:"nomatch,omitempty"` //
	Vmid    int64   `json:"vmid,omitempty"`    // The (unique) ID of the VM.
}

type NodesQemuFirewallIpsetUpdateIpResponse struct{}

// Update IP or Network settings
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}
func (c *Client) NodesQemuFirewallIpsetUpdateIp(ctx context.Context, request *NodesQemuFirewallIpsetUpdateIpRequest) (*NodesQemuFirewallIpsetUpdateIpResponse, error) {

	method := "PUT"
	path := "/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}"

	var response NodesQemuFirewallIpsetUpdateIpResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}