// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesCapabilitiesQemuQemuCapsIndexRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesCapabilitiesQemuQemuCapsIndexResponseItem struct {
}

type NodesCapabilitiesQemuQemuCapsIndexResponse []NodesCapabilitiesQemuQemuCapsIndexResponseItem

// QEMU capabilities index.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/capabilities/qemu
func (c *Client) NodesCapabilitiesQemuQemuCapsIndex(ctx context.Context, request *NodesCapabilitiesQemuQemuCapsIndexRequest) (*NodesCapabilitiesQemuQemuCapsIndexResponse, error) {

	method := "GET"
	path := "/nodes/{node}/capabilities/qemu"

	var response NodesCapabilitiesQemuQemuCapsIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
