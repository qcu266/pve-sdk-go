// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesScanPbsPbsscanRequest struct {
	Fingerprint *string `query:"fingerprint,omitempty"` // Certificate SHA 256 fingerprint.
	Node        string  `query:"node,omitempty"`        // The cluster node name.
	Password    string  `query:"password,omitempty"`    // User password or API token secret.
	Port        *int64  `query:"port,omitempty"`        // Optional port.
	Server      string  `query:"server,omitempty"`      // The server address (name or IP).
	Username    string  `query:"username,omitempty"`    // User-name or API token-ID.
}

type NodesScanPbsPbsscanResponseItem struct {
	Comment *string `json:"comment,omitempty"` // Comment from server.
	Store   string  `json:"store,omitempty"`   // The datastore name.
}

type NodesScanPbsPbsscanResponse []NodesScanPbsPbsscanResponseItem

// Scan remote Proxmox Backup Server.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/scan/pbs
func (c *Client) NodesScanPbsPbsscan(ctx context.Context, request *NodesScanPbsPbsscanRequest) (*NodesScanPbsPbsscanResponse, error) {

	method := "GET"
	path := "/nodes/{node}/scan/pbs"

	var response NodesScanPbsPbsscanResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
