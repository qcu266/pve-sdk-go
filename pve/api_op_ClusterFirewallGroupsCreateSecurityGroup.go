// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterFirewallGroupsCreateSecurityGroupRequest struct {
	Comment *string `json:"comment,omitempty"` //
	Digest  *string `json:"digest,omitempty"`  // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Group   string  `json:"group,omitempty"`   // Security Group name.
	Rename  *string `json:"rename,omitempty"`  // Rename/update an existing security group. You can set 'rename' to the same value as 'name' to update the 'comment' of an existing group.
}

type ClusterFirewallGroupsCreateSecurityGroupResponse struct{}

// Create new security group.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/firewall/groups
func (c *Client) ClusterFirewallGroupsCreateSecurityGroup(ctx context.Context, request *ClusterFirewallGroupsCreateSecurityGroupRequest) (*ClusterFirewallGroupsCreateSecurityGroupResponse, error) {

	method := "POST"
	path := "/cluster/firewall/groups"

	var response ClusterFirewallGroupsCreateSecurityGroupResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
