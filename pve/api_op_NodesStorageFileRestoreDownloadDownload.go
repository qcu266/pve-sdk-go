// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesStorageFileRestoreDownloadDownloadRequest struct {
	Filepath string `query:"filepath,omitempty"` // base64-path to the directory or file to download.
	Node     string `query:"node,omitempty"`     // The cluster node name.
	Storage  string `query:"storage,omitempty"`  // The storage identifier.
	Volume   string `query:"volume,omitempty"`   // Backup volume ID or name. Currently only PBS snapshots are supported.
}

type NodesStorageFileRestoreDownloadDownloadResponse interface{}

// Extract a file or directory (as zip archive) from a PBS backup.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/file-restore/download
func (c *Client) NodesStorageFileRestoreDownloadDownload(ctx context.Context, request *NodesStorageFileRestoreDownloadDownloadRequest) (*NodesStorageFileRestoreDownloadDownloadResponse, error) {

	method := "GET"
	path := "/nodes/{node}/storage/{storage}/file-restore/download"

	var response NodesStorageFileRestoreDownloadDownloadResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}