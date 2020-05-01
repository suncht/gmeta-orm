package models

import "time"

// SysColumnDisplay 列显示定义
type SysColumnDisplay struct {
	ID            int64     `gorm:"primary_key;column:id;type:bigint(20);not null"` // 主键
	ColumnID      int64     `gorm:"column:column_id;type:bigint(20)"`               // 列ID
	ColumnName    string    `gorm:"column:column_name;type:varchar(255)"`           // 列名
	ShowName      string    `gorm:"column:show_name;type:varchar(255)"`             // 显示名
	DisplayType   string    `gorm:"column:display_type;type:varchar(50)"`           // 显示组件类型：input、select、checkbox等
	DisplayLength int       `gorm:"column:display_length;type:int(11)"`             // 显示长度
	DisplayFormat string    `gorm:"column:display_format;type:varchar(255)"`        // 显示格式
	ShowSeq       int       `gorm:"column:show_seq;type:int(11)"`                   // 显示排序，越小越靠前
	Showable      int8      `gorm:"column:showable;type:tinyint(1)"`                // 是否显示 1-是 2-否
	Editable      int8      `gorm:"column:editable;type:tinyint(1)"`                // 是否可编辑 1-是 2-否
	Orderable     int8      `gorm:"column:orderable;type:tinyint(1)"`               // 是否可排序 1-是 2-否
	OrderType     int8      `gorm:"column:order_type;type:tinyint(1)"`              // 排序类型 1-no 2-asc 3-desc
	CreateTime    time.Time `gorm:"column:create_time;type:datetime"`               // 创建时间
	CreateUser    string    `gorm:"column:create_user;type:varchar(50)"`            // 创建人
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime"`               // 更新时间
	UpdateUser    string    `gorm:"column:update_user;type:varchar(50)"`            // 更新人
	Ts            time.Time `gorm:"column:ts;type:timestamp"`                       // 时间戳
	Yn            int8      `gorm:"column:yn;type:tinyint(1)"`                      // 是否有效， 1-正常  0-删除
}
