// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesCephPoolsGetpoolRequest struct {
	Name    string `query:"name,omitempty"`    // The name of the pool. It must be unique.
	Node    string `query:"node,omitempty"`    // The cluster node name.
	Verbose *bool  `query:"verbose,omitempty"` // If enabled, will display additional data(eg. statistics).
}

type NodesCephPoolsGetpoolResponseStatistics interface{}

type NodesCephPoolsGetpoolResponseAutoscaleStatus interface{}

type NodesCephPoolsGetpoolResponse struct {
	Application          *string                                       `json:"application,omitempty"`            // The application of the pool.
	ApplicationList      *[]interface{}                                `json:"application_list,omitempty"`       //
	AutoscaleStatus      *NodesCephPoolsGetpoolResponseAutoscaleStatus `json:"autoscale_status,omitempty"`       //
	CrushRule            *string                                       `json:"crush_rule,omitempty"`             // The rule to use for mapping object placement in the cluster.
	FastRead             bool                                          `json:"fast_read,omitempty"`              //
	Hashpspool           bool                                          `json:"hashpspool,omitempty"`             //
	Id                   int64                                         `json:"id,omitempty"`                     //
	MinSize              *int64                                        `json:"min_size,omitempty"`               // Minimum number of replicas per object
	Name                 string                                        `json:"name,omitempty"`                   // The name of the pool. It must be unique.
	NodeepScrub          bool                                          `json:"nodeep-scrub,omitempty"`           //
	Nodelete             bool                                          `json:"nodelete,omitempty"`               //
	Nopgchange           bool                                          `json:"nopgchange,omitempty"`             //
	Noscrub              bool                                          `json:"noscrub,omitempty"`                //
	Nosizechange         bool                                          `json:"nosizechange,omitempty"`           //
	PgAutoscaleMode      *string                                       `json:"pg_autoscale_mode,omitempty"`      // The automatic PG scaling mode of the pool.
	PgNum                *int64                                        `json:"pg_num,omitempty"`                 // Number of placement groups.
	PgNumMin             *int64                                        `json:"pg_num_min,omitempty"`             // Minimal number of placement groups.
	PgpNum               int64                                         `json:"pgp_num,omitempty"`                //
	Size                 *int64                                        `json:"size,omitempty"`                   // Number of replicas per object
	Statistics           *NodesCephPoolsGetpoolResponseStatistics      `json:"statistics,omitempty"`             //
	TargetSize           *string                                       `json:"target_size,omitempty"`            // The estimated target size of the pool for the PG autoscaler.
	TargetSizeRatio      *float64                                      `json:"target_size_ratio,omitempty"`      // The estimated target ratio of the pool for the PG autoscaler.
	UseGmtHitset         bool                                          `json:"use_gmt_hitset,omitempty"`         //
	WriteFadviseDontneed bool                                          `json:"write_fadvise_dontneed,omitempty"` //
}

// List pool settings.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/ceph/pools/{name}
func (c *Client) NodesCephPoolsGetpool(ctx context.Context, request *NodesCephPoolsGetpoolRequest) (*NodesCephPoolsGetpoolResponse, error) {

	method := "GET"
	path := "/nodes/{node}/ceph/pools/{name}"

	var response NodesCephPoolsGetpoolResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}