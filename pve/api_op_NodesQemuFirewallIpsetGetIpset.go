// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuFirewallIpsetGetIpsetRequest struct {
	Name string `query:"name,omitempty"` // IP set name.
	Node string `query:"node,omitempty"` // The cluster node name.
	Vmid int64  `query:"vmid,omitempty"` // The (unique) ID of the VM.
}

type NodesQemuFirewallIpsetGetIpsetResponseItem struct {
	Cidr    string  `json:"cidr,omitempty"`    //
	Comment *string `json:"comment,omitempty"` //
	Digest  string  `json:"digest,omitempty"`  // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Nomatch *bool   `json:"nomatch,omitempty"` //
}

type NodesQemuFirewallIpsetGetIpsetResponse []NodesQemuFirewallIpsetGetIpsetResponseItem

// List IPSet content
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}
func (c *Client) NodesQemuFirewallIpsetGetIpset(ctx context.Context, request *NodesQemuFirewallIpsetGetIpsetRequest) (*NodesQemuFirewallIpsetGetIpsetResponse, error) {

	method := "GET"
	path := "/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}"

	var response NodesQemuFirewallIpsetGetIpsetResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
