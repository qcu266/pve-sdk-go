// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcFirewallRulesGetRulesRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
	Vmid int64  `query:"vmid,omitempty"` // The (unique) ID of the VM.
}

type NodesLxcFirewallRulesGetRulesResponseItem struct {
	Pos int64 `json:"pos,omitempty"` //
}

type NodesLxcFirewallRulesGetRulesResponse []NodesLxcFirewallRulesGetRulesResponseItem

// List rules.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/firewall/rules
func (c *Client) NodesLxcFirewallRulesGetRules(ctx context.Context, request *NodesLxcFirewallRulesGetRulesRequest) (*NodesLxcFirewallRulesGetRulesResponse, error) {

	method := "GET"
	path := "/nodes/{node}/lxc/{vmid}/firewall/rules"

	var response NodesLxcFirewallRulesGetRulesResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
