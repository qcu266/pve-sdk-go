// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterFirewallOptionsGetOptionsRequest interface{}

type ClusterFirewallOptionsGetOptionsResponse struct {
	Ebtables     *bool   `json:"ebtables,omitempty"`      // Enable ebtables rules cluster wide.
	Enable       *int64  `json:"enable,omitempty"`        // Enable or disable the firewall cluster wide.
	LogRatelimit *string `json:"log_ratelimit,omitempty"` // Log ratelimiting settings
	PolicyIn     *string `json:"policy_in,omitempty"`     // Input policy.
	PolicyOut    *string `json:"policy_out,omitempty"`    // Output policy.
}

// Get Firewall options.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/firewall/options
func (c *Client) ClusterFirewallOptionsGetOptions(ctx context.Context, request *ClusterFirewallOptionsGetOptionsRequest) (*ClusterFirewallOptionsGetOptionsResponse, error) {

	method := "GET"
	path := "/cluster/firewall/options"

	var response ClusterFirewallOptionsGetOptionsResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
