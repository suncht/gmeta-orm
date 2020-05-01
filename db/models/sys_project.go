package models

import "time"

type SysProject struct {
	ID          uint64    `gorm:"primary_key" json:" - "` //主键
	ProjectName string    `json:"project_name"`           //项目名
	ProjectDesc string    `json:"project_desc"`           //项目描述
	CreateTime  time.Time `json:"create_time"`            //创建时间
	CreateUser  string    `json:"create_user"`            //创建人
	UpdateTime  time.Time `json:"update_time"`            //更新时间
	UpdateUser  string    `json:"update_user"`            //更新人
	Ts          time.Time `json:"ts"`                     //时间戳
	Yn          uint8     `json:"yn"`                     //是否有效， 1-正常  0-删除
}
