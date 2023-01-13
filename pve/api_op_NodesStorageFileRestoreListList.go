// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesStorageFileRestoreListListRequest struct {
	Filepath string `query:"filepath,omitempty"` // base64-path to the directory or file being listed, or "/".
	Node     string `query:"node,omitempty"`     // The cluster node name.
	Storage  string `query:"storage,omitempty"`  // The storage identifier.
	Volume   string `query:"volume,omitempty"`   // Backup volume ID or name. Currently only PBS snapshots are supported.
}

type NodesStorageFileRestoreListListResponseItem struct {
	Filepath string `json:"filepath,omitempty"` // base64 path of the current entry
	Leaf     bool   `json:"leaf,omitempty"`     // If this entry is a leaf in the directory graph.
	Mtime    *int64 `json:"mtime,omitempty"`    // Entry last-modified time (unix timestamp).
	Size     *int64 `json:"size,omitempty"`     // Entry file size.
	Text     string `json:"text,omitempty"`     // Entry display text.
	Type     string `json:"type,omitempty"`     // Entry type.
}

type NodesStorageFileRestoreListListResponse []NodesStorageFileRestoreListListResponseItem

// List files and directories for single file restore under the given path.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/file-restore/list
func (c *Client) NodesStorageFileRestoreListList(ctx context.Context, request *NodesStorageFileRestoreListListRequest) (*NodesStorageFileRestoreListListResponse, error) {

	method := "GET"
	path := "/nodes/{node}/storage/{storage}/file-restore/list"

	var response NodesStorageFileRestoreListListResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
