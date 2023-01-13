// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterStatusGetStatusRequest interface{}

type ClusterStatusGetStatusResponse []ClusterStatusGetStatusResponseItem

type ClusterStatusGetStatusResponseItem struct {
	Id      string  `json:"id,omitempty"`      //
	Ip      *string `json:"ip,omitempty"`      // [node] IP of the resolved nodename.
	Level   *string `json:"level,omitempty"`   // [node] Proxmox VE Subscription level, indicates if eligible for enterprise support as well as access to the stable Proxmox VE Enterprise Repository.
	Local   *bool   `json:"local,omitempty"`   // [node] Indicates if this is the responding node.
	Name    string  `json:"name,omitempty"`    //
	Nodeid  *int64  `json:"nodeid,omitempty"`  // [node] ID of the node from the corosync configuration.
	Nodes   *int64  `json:"nodes,omitempty"`   // [cluster] Nodes count, including offline nodes.
	Online  *bool   `json:"online,omitempty"`  // [node] Indicates if the node is online or offline.
	Quorate *bool   `json:"quorate,omitempty"` // [cluster] Indicates if there is a majority of nodes online to make decisions
	Type    string  `json:"type,omitempty"`    // Indicates the type, either cluster or node. The type defines the object properties e.g. quorate available for type cluster.
	Version *int64  `json:"version,omitempty"` // [cluster] Current version of the corosync configuration file.
}

// Get cluster status information.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/status
func (c *Client) ClusterStatusGetStatus(ctx context.Context, request *ClusterStatusGetStatusRequest) (*ClusterStatusGetStatusResponse, error) {

	method := "GET"
	path := "/cluster/status"

	var response ClusterStatusGetStatusResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
