package models

import "time"

//
type SysColumnDisplay struct {
	ID            uint64    `gorm:"primary_key" json:" - "` //主键
	ColumnId      uint64    `json:"column_id"`              //列ID
	ColumnName    string    `json:"column_name"`            //列名
	ShowName      string    `json:"show_name"`              //显示名
	DisplayType   string    `json:"display_type"`           //显示组件类型：input、select、checkbox等
	DisplayLength int       `json:"display_length"`         //显示长度
	DisplayFormat string    `json:"display_format"`         //显示格式
	ShowSeq       int       `json:"show_seq"`               //显示排序，越小越靠前
	Showable      uint8     `json:"showable"`               //是否显示 1-是 2-否
	Editable      uint8     `json:"editable"`               //是否可编辑 1-是 2-否
	Orderable     uint8     `json:"orderable"`              //是否可排序 1-是 2-否
	OrderType     uint8     `json:"order_type"`             //排序类型 1-no 2-asc 3-desc
	CreateTime    time.Time `json:"create_time"`            //创建时间
	CreateUser    string    `json:"create_user"`            //创建人
	UpdateTime    time.Time `json:"update_time"`            //更新时间
	UpdateUser    string    `json:"update_user"`            //更新人
	Ts            time.Time `json:"ts"`                     //时间戳
	Yn            uint8     `json:"yn"`                     //是否有效， 1-正常  0-删除
}
