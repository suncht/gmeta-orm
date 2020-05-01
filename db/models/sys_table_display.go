package models

import "time"

// SysTableDisplay 表显示定义
type SysTableDisplay struct {
	ID             int64     `gorm:"primary_key;column:id;type:bigint(20);not null"` // 主键
	TableID        int64     `gorm:"column:table_id;type:bigint(20)"`                // 表ID
	TableName      string    `gorm:"column:table_name;type:varchar(50);not null"`    // 表名
	NeedTree       int8      `gorm:"column:need_tree;type:tinyint(1)"`               // 是否需要树表 1-是 2-否
	NeedAsync      int8      `gorm:"column:need_async;type:tinyint(1)"`              // 是否需要异步加载 1-是 2-否
	NeedPagination int8      `gorm:"column:need_pagination;type:tinyint(1)"`         // 是否需要分页
	NeedSearch     int8      `gorm:"column:need_search;type:tinyint(1)"`             // 是否分页
	CreateTime     time.Time `gorm:"column:create_time;type:datetime"`               // 创建时间
	CreateUser     string    `gorm:"column:create_user;type:varchar(50)"`            // 创建人
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime"`               // 更新时间
	UpdateUser     string    `gorm:"column:update_user;type:varchar(50)"`            // 更新人
	Ts             time.Time `gorm:"column:ts;type:timestamp"`                       // 时间戳
	Yn             int8      `gorm:"column:yn;type:tinyint(1)"`                      // 是否有效， 1-正常  0-删除
}
