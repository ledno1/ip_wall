// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRole is the golang structure for table admin_role.
type AdminRole struct {
	Id         int64       `json:"id"         description:"角色ID"`
	Name       string      `json:"name"       description:"角色名称"`
	Key        string      `json:"key"        description:"角色权限字符串"`
	DataScope  int         `json:"dataScope"  description:"数据范围"`
	CustomDept *gjson.Json `json:"customDept" description:"自定义部门权限"`
	Pid        int64       `json:"pid"        description:"上级角色ID"`
	Level      int         `json:"level"      description:"关系树等级"`
	Tree       string      `json:"tree"       description:"关系树"`
	Remark     string      `json:"remark"     description:"备注"`
	Sort       int         `json:"sort"       description:"排序"`
	Status     int         `json:"status"     description:"角色状态"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"更新时间"`
}