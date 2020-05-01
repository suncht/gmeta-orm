package models

import "time"

//
type SysTableDisplay struct {
	ID             uint64    `gorm:"primary_key" json:" - "` //主键
	TableId        uint64    `json:"table_id"`               //表ID
	TableName      string    `json:"table_name"`             //表名
	NeedTree       uint8     `json:"need_tree"`              //是否需要树表 1-是 2-否
	NeedAsync      uint8     `json:"need_async"`             //是否需要异步加载 1-是 2-否
	NeedPagination uint8     `json:"need_pagination"`        //是否需要分页
	NeedSearch     uint8     `json:"need_search"`            //是否分页
	CreateTime     time.Time `json:"create_time"`            //创建时间
	CreateUser     string    `json:"create_user"`            //创建人
	UpdateTime     time.Time `json:"update_time"`            //更新时间
	UpdateUser     string    `json:"update_user"`            //更新人
	Ts             time.Time `json:"ts"`                     //时间戳
	Yn             uint8     `json:"yn"`                     //是否有效， 1-正常  0-删除
}
