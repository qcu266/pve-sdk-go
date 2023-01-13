// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcFirewallRefsRefsRequest struct {
	Node string  `query:"node,omitempty"` // The cluster node name.
	Type *string `query:"type,omitempty"` // Only list references of specified type.
	Vmid int64   `query:"vmid,omitempty"` // The (unique) ID of the VM.
}

type NodesLxcFirewallRefsRefsResponse []NodesLxcFirewallRefsRefsResponseItem

type NodesLxcFirewallRefsRefsResponseItem struct {
	Comment *string `json:"comment,omitempty"` //
	Name    string  `json:"name,omitempty"`    //
	Type    string  `json:"type,omitempty"`    //
}

// Lists possible IPSet/Alias reference which are allowed in source/dest properties.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/firewall/refs
func (c *Client) NodesLxcFirewallRefsRefs(ctx context.Context, request *NodesLxcFirewallRefsRefsRequest) (*NodesLxcFirewallRefsRefsResponse, error) {

	method := "GET"
	path := "/nodes/{node}/lxc/{vmid}/firewall/refs"

	var response NodesLxcFirewallRefsRefsResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
