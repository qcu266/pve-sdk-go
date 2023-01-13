// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterConfigNodesAddnodeRequest struct {
	Apiversion *int64  `json:"apiversion,omitempty"`  // The JOIN_API_VERSION of the new node.
	Force      *bool   `json:"force,omitempty"`       // Do not throw error if node already exists.
	Linkn      *string `json:"link[n],omitempty"`     // Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	NewNodeIp  *string `json:"new_node_ip,omitempty"` // IP Address of node to add. Used as fallback if no links are given.
	Node       string  `json:"node,omitempty"`        // The cluster node name.
	Nodeid     *int64  `json:"nodeid,omitempty"`      // Node id for this node.
	Votes      *int64  `json:"votes,omitempty"`       // Number of votes for this node
}

type ClusterConfigNodesAddnodeResponseWarningsItem string

type ClusterConfigNodesAddnodeResponse struct {
	CorosyncAuthkey string                                          `json:"corosync_authkey,omitempty"` //
	CorosyncConf    string                                          `json:"corosync_conf,omitempty"`    //
	Warnings        []ClusterConfigNodesAddnodeResponseWarningsItem `json:"warnings,omitempty"`         //
}

// Adds a node to the cluster configuration. This call is for internal use.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/config/nodes/{node}
func (c *Client) ClusterConfigNodesAddnode(ctx context.Context, request *ClusterConfigNodesAddnodeRequest) (*ClusterConfigNodesAddnodeResponse, error) {

	method := "POST"
	path := "/cluster/config/nodes/{node}"

	var response ClusterConfigNodesAddnodeResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}