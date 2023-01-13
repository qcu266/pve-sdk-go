// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessRolesReadRoleRequest struct {
	Roleid string `query:"roleid,omitempty"` //
}

type AccessRolesReadRoleResponse struct {
	DatastoreAllocate         *bool `json:"Datastore.Allocate,omitempty"`         //
	DatastoreAllocateSpace    *bool `json:"Datastore.AllocateSpace,omitempty"`    //
	DatastoreAllocateTemplate *bool `json:"Datastore.AllocateTemplate,omitempty"` //
	DatastoreAudit            *bool `json:"Datastore.Audit,omitempty"`            //
	GroupAllocate             *bool `json:"Group.Allocate,omitempty"`             //
	PermissionsModify         *bool `json:"Permissions.Modify,omitempty"`         //
	PoolAllocate              *bool `json:"Pool.Allocate,omitempty"`              //
	PoolAudit                 *bool `json:"Pool.Audit,omitempty"`                 //
	RealmAllocate             *bool `json:"Realm.Allocate,omitempty"`             //
	RealmAllocateUser         *bool `json:"Realm.AllocateUser,omitempty"`         //
	SDNAllocate               *bool `json:"SDN.Allocate,omitempty"`               //
	SDNAudit                  *bool `json:"SDN.Audit,omitempty"`                  //
	SysAudit                  *bool `json:"Sys.Audit,omitempty"`                  //
	SysConsole                *bool `json:"Sys.Console,omitempty"`                //
	SysIncoming               *bool `json:"Sys.Incoming,omitempty"`               //
	SysModify                 *bool `json:"Sys.Modify,omitempty"`                 //
	SysPowerMgmt              *bool `json:"Sys.PowerMgmt,omitempty"`              //
	SysSyslog                 *bool `json:"Sys.Syslog,omitempty"`                 //
	UserModify                *bool `json:"User.Modify,omitempty"`                //
	VMAllocate                *bool `json:"VM.Allocate,omitempty"`                //
	VMAudit                   *bool `json:"VM.Audit,omitempty"`                   //
	VMBackup                  *bool `json:"VM.Backup,omitempty"`                  //
	VMClone                   *bool `json:"VM.Clone,omitempty"`                   //
	VMConfigCDROM             *bool `json:"VM.Config.CDROM,omitempty"`            //
	VMConfigCPU               *bool `json:"VM.Config.CPU,omitempty"`              //
	VMConfigCloudinit         *bool `json:"VM.Config.Cloudinit,omitempty"`        //
	VMConfigDisk              *bool `json:"VM.Config.Disk,omitempty"`             //
	VMConfigHWType            *bool `json:"VM.Config.HWType,omitempty"`           //
	VMConfigMemory            *bool `json:"VM.Config.Memory,omitempty"`           //
	VMConfigNetwork           *bool `json:"VM.Config.Network,omitempty"`          //
	VMConfigOptions           *bool `json:"VM.Config.Options,omitempty"`          //
	VMConsole                 *bool `json:"VM.Console,omitempty"`                 //
	VMMigrate                 *bool `json:"VM.Migrate,omitempty"`                 //
	VMMonitor                 *bool `json:"VM.Monitor,omitempty"`                 //
	VMPowerMgmt               *bool `json:"VM.PowerMgmt,omitempty"`               //
	VMSnapshot                *bool `json:"VM.Snapshot,omitempty"`                //
	VMSnapshotRollback        *bool `json:"VM.Snapshot.Rollback,omitempty"`       //
}

// Get role configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/roles/{roleid}
func (c *Client) AccessRolesReadRole(ctx context.Context, request *AccessRolesReadRoleRequest) (*AccessRolesReadRoleResponse, error) {

	method := "GET"
	path := "/access/roles/{roleid}"

	var response AccessRolesReadRoleResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
