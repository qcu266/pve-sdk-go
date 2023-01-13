// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesHardwarePciMdevMdevscanRequest struct {
	Node  string `query:"node,omitempty"`  // The cluster node name.
	Pciid string `query:"pciid,omitempty"` // The PCI ID to list the mdev types for.
}

type NodesHardwarePciMdevMdevscanResponse []NodesHardwarePciMdevMdevscanResponseItem

type NodesHardwarePciMdevMdevscanResponseItem struct {
	Available   int64  `json:"available,omitempty"`   // The number of still available instances of this type.
	Description string `json:"description,omitempty"` //
	Type        string `json:"type,omitempty"`        // The name of the mdev type.
}

// List mediated device types for given PCI device.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/hardware/pci/{pciid}/mdev
func (c *Client) NodesHardwarePciMdevMdevscan(ctx context.Context, request *NodesHardwarePciMdevMdevscanRequest) (*NodesHardwarePciMdevMdevscanResponse, error) {

	method := "GET"
	path := "/nodes/{node}/hardware/pci/{pciid}/mdev"

	var response NodesHardwarePciMdevMdevscanResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
