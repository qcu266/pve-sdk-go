// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesCephPoolsCreatepoolRequest struct {
	AddStorages     *bool    `json:"add_storages,omitempty"`      // Configure VM and CT storage using the new pool.
	Application     *string  `json:"application,omitempty"`       // The application of the pool.
	CrushRule       *string  `json:"crush_rule,omitempty"`        // The rule to use for mapping object placement in the cluster.
	ErasureCoding   *string  `json:"erasure-coding,omitempty"`    // Create an erasure coded pool for RBD with an accompaning replicated pool for metadata storage. With EC, the common ceph options 'size', 'min_size' and 'crush_rule' parameters will be applied to the metadata pool.
	MinSize         *int64   `json:"min_size,omitempty"`          // Minimum number of replicas per object
	Name            string   `json:"name,omitempty"`              // The name of the pool. It must be unique.
	Node            string   `json:"node,omitempty"`              // The cluster node name.
	PgAutoscaleMode *string  `json:"pg_autoscale_mode,omitempty"` // The automatic PG scaling mode of the pool.
	PgNum           *int64   `json:"pg_num,omitempty"`            // Number of placement groups.
	PgNumMin        *int64   `json:"pg_num_min,omitempty"`        // Minimal number of placement groups.
	Size            *int64   `json:"size,omitempty"`              // Number of replicas per object
	TargetSize      *string  `json:"target_size,omitempty"`       // The estimated target size of the pool for the PG autoscaler.
	TargetSizeRatio *float64 `json:"target_size_ratio,omitempty"` // The estimated target ratio of the pool for the PG autoscaler.
}

type NodesCephPoolsCreatepoolResponse string

// Create Ceph pool
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/ceph/pools
func (c *Client) NodesCephPoolsCreatepool(ctx context.Context, request *NodesCephPoolsCreatepoolRequest) (*NodesCephPoolsCreatepoolResponse, error) {

	method := "POST"
	path := "/nodes/{node}/ceph/pools"

	var response NodesCephPoolsCreatepoolResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
