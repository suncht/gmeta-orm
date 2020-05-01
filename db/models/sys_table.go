package models

import "time"

//
type SysTable struct {
	ID              uint64    `gorm:"primary_key" json:" - "` //主键
	ProjectId       uint64    `json:"project_id"`             //项目表ID
	TableName       string    `json:"table_name"`             //表名
	TableComment    string    `json:"table_comment"`          //表注释
	IsSystem        uint8     `json:"is_system"`              //是否系统表
	TableCategory   string    `json:"table_category"`         //表分类
	TableMode       string    `json:"table_mode"`             //表模式：single主表、sub主子表、ext扩展表
	ParentTableName string    `json:"parent_table_name"`      //从属于哪张表
	Status          uint8     `json:"status"`                 //表状态： 1-新建未同步 2-修改未同步 3-删除未同步 10-已同步
	Enabled         uint8     `json:"enabled"`                //启用状态： 1-启用 2-禁用
	CreateTime      time.Time `json:"create_time"`            //创建时间
	CreateUser      string    `json:"create_user"`            //创建人
	UpdateTime      time.Time `json:"update_time"`            //更新时间
	UpdateUser      string    `json:"update_user"`            //更新人
	Ts              time.Time `json:"ts"`                     //时间戳
	Yn              uint8     `json:"yn"`                     //是否有效， 1-正常  0-删除
}
