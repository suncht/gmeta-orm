package models

import "time"

// SysProject 项目
type SysProject struct {
	ID          int64     `gorm:"primary_key;column:id;type:bigint(20);not null"` // 主键
	ProjectName string    `gorm:"column:project_name;type:varchar(50)"`           // 项目名
	ProjectDesc string    `gorm:"column:project_desc;type:varchar(100)"`          // 项目描述
	CreateTime  time.Time `gorm:"column:create_time;type:datetime"`               // 创建时间
	CreateUser  string    `gorm:"column:create_user;type:varchar(50)"`            // 创建人
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime"`               // 更新时间
	UpdateUser  string    `gorm:"column:update_user;type:varchar(50)"`            // 更新人
	Ts          time.Time `gorm:"column:ts;type:timestamp"`                       // 时间戳
	Yn          int8      `gorm:"column:yn;type:tinyint(1)"`                      // 是否有效， 1-正常  0-删除
}
