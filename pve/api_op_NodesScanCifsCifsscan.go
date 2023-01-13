// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesScanCifsCifsscanRequest struct {
	Domain   *string `query:"domain,omitempty"`   // SMB domain (Workgroup).
	Node     string  `query:"node,omitempty"`     // The cluster node name.
	Password *string `query:"password,omitempty"` // User password.
	Server   string  `query:"server,omitempty"`   // The server address (name or IP).
	Username *string `query:"username,omitempty"` // User name.
}

type NodesScanCifsCifsscanResponseItem struct {
	Description string `json:"description,omitempty"` // Descriptive text from server.
	Share       string `json:"share,omitempty"`       // The cifs share name.
}

type NodesScanCifsCifsscanResponse []NodesScanCifsCifsscanResponseItem

// Scan remote CIFS server.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/scan/cifs
func (c *Client) NodesScanCifsCifsscan(ctx context.Context, request *NodesScanCifsCifsscanRequest) (*NodesScanCifsCifsscanResponse, error) {

	method := "GET"
	path := "/nodes/{node}/scan/cifs"

	var response NodesScanCifsCifsscanResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
