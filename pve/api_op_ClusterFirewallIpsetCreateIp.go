// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterFirewallIpsetCreateIpRequest struct {
	Cidr    string  `json:"cidr,omitempty"`    // Network/IP specification in CIDR format.
	Comment *string `json:"comment,omitempty"` //
	Name    string  `json:"name,omitempty"`    // IP set name.
	Nomatch *bool   `json:"nomatch,omitempty"` //
}

type ClusterFirewallIpsetCreateIpResponse struct{}

// Add IP or Network to IPSet.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/firewall/ipset/{name}
func (c *Client) ClusterFirewallIpsetCreateIp(ctx context.Context, request *ClusterFirewallIpsetCreateIpRequest) (*ClusterFirewallIpsetCreateIpResponse, error) {

	method := "POST"
	path := "/cluster/firewall/ipset/{name}"

	var response ClusterFirewallIpsetCreateIpResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
