// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcFirewallIpsetIpsetIndexRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
	Vmid int64  `query:"vmid,omitempty"` // The (unique) ID of the VM.
}

type NodesLxcFirewallIpsetIpsetIndexResponseItem struct {
	Comment *string `json:"comment,omitempty"` //
	Digest  string  `json:"digest,omitempty"`  // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string  `json:"name,omitempty"`    // IP set name.
}

type NodesLxcFirewallIpsetIpsetIndexResponse []NodesLxcFirewallIpsetIpsetIndexResponseItem

// List IPSets
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/firewall/ipset
func (c *Client) NodesLxcFirewallIpsetIpsetIndex(ctx context.Context, request *NodesLxcFirewallIpsetIpsetIndexRequest) (*NodesLxcFirewallIpsetIpsetIndexResponse, error) {

	method := "GET"
	path := "/nodes/{node}/lxc/{vmid}/firewall/ipset"

	var response NodesLxcFirewallIpsetIpsetIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}