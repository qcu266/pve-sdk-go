// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesScanNfsNfsscanRequest struct {
	Node   string `query:"node,omitempty"`   // The cluster node name.
	Server string `query:"server,omitempty"` // The server address (name or IP).
}

type NodesScanNfsNfsscanResponseItem struct {
	Options string `json:"options,omitempty"` // NFS export options.
	Path    string `json:"path,omitempty"`    // The exported path.
}

type NodesScanNfsNfsscanResponse []NodesScanNfsNfsscanResponseItem

// Scan remote NFS server.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/scan/nfs
func (c *Client) NodesScanNfsNfsscan(ctx context.Context, request *NodesScanNfsNfsscanRequest) (*NodesScanNfsNfsscanResponse, error) {

	method := "GET"
	path := "/nodes/{node}/scan/nfs"

	var response NodesScanNfsNfsscanResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
