// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesTasksNodeTasksRequest struct {
	Errors       *bool   `query:"errors,omitempty"`       // Only list tasks with a status of ERROR.
	Limit        *int64  `query:"limit,omitempty"`        // Only list this amount of tasks.
	Node         string  `query:"node,omitempty"`         // The cluster node name.
	Since        *int64  `query:"since,omitempty"`        // Only list tasks since this UNIX epoch.
	Source       *string `query:"source,omitempty"`       // List archived, active or all tasks.
	Start        *int64  `query:"start,omitempty"`        // List tasks beginning from this offset.
	Statusfilter *string `query:"statusfilter,omitempty"` // List of Task States that should be returned.
	Typefilter   *string `query:"typefilter,omitempty"`   // Only list tasks of this type (e.g., vzstart, vzdump).
	Until        *int64  `query:"until,omitempty"`        // Only list tasks until this UNIX epoch.
	Userfilter   *string `query:"userfilter,omitempty"`   // Only list tasks from this user.
	Vmid         *int64  `query:"vmid,omitempty"`         // Only list tasks for this VM.
}

type NodesTasksNodeTasksResponse []NodesTasksNodeTasksResponseItem

type NodesTasksNodeTasksResponseItem struct {
	Endtime   *int64  `json:"endtime,omitempty"`   //
	Id        string  `json:"id,omitempty"`        //
	Node      string  `json:"node,omitempty"`      //
	Pid       int64   `json:"pid,omitempty"`       //
	Pstart    int64   `json:"pstart,omitempty"`    //
	Starttime int64   `json:"starttime,omitempty"` //
	Status    *string `json:"status,omitempty"`    //
	Type      string  `json:"type,omitempty"`      //
	Upid      string  `json:"upid,omitempty"`      //
	User      string  `json:"user,omitempty"`      //
}

// Read task list for one node (finished tasks).
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/tasks
func (c *Client) NodesTasksNodeTasks(ctx context.Context, request *NodesTasksNodeTasksRequest) (*NodesTasksNodeTasksResponse, error) {

	method := "GET"
	path := "/nodes/{node}/tasks"

	var response NodesTasksNodeTasksResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
