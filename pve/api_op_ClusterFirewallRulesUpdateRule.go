// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterFirewallRulesUpdateRuleRequest struct {
	Action   *string `json:"action,omitempty"`    // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Comment  *string `json:"comment,omitempty"`   // Descriptive comment.
	Delete   *string `json:"delete,omitempty"`    // A list of settings you want to delete.
	Dest     *string `json:"dest,omitempty"`      // Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Digest   *string `json:"digest,omitempty"`    // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Dport    *string `json:"dport,omitempty"`     // Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Enable   *int64  `json:"enable,omitempty"`    // Flag to enable/disable a rule.
	IcmpType *string `json:"icmp-type,omitempty"` // Specify icmp-type. Only valid if proto equals 'icmp'.
	Iface    *string `json:"iface,omitempty"`     // Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Log      *string `json:"log,omitempty"`       // Log level for firewall rule.
	Macro    *string `json:"macro,omitempty"`     // Use predefined standard macro.
	Moveto   *int64  `json:"moveto,omitempty"`    // Move rule to new position <moveto>. Other arguments are ignored.
	Pos      *int64  `json:"pos,omitempty"`       // Update rule at position <pos>.
	Proto    *string `json:"proto,omitempty"`     // IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Source   *string `json:"source,omitempty"`    // Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Sport    *string `json:"sport,omitempty"`     // Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Type     *string `json:"type,omitempty"`      // Rule type.
}

type ClusterFirewallRulesUpdateRuleResponse struct{}

// Modify rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/firewall/rules/{pos}
func (c *Client) ClusterFirewallRulesUpdateRule(ctx context.Context, request *ClusterFirewallRulesUpdateRuleRequest) (*ClusterFirewallRulesUpdateRuleResponse, error) {

	method := "PUT"
	path := "/cluster/firewall/rules/{pos}"

	var response ClusterFirewallRulesUpdateRuleResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
