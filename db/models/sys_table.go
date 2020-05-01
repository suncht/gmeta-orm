package models

import "time"

// SysTable 表定义
type SysTable struct {
	ID              int64     `gorm:"primary_key;column:id;type:bigint(20);not null"` // 主键
	ProjectID       int64     `gorm:"column:project_id;type:bigint(20)"`              // 项目表ID
	TableName       string    `gorm:"column:table_name;type:varchar(50)"`             // 表名
	TableComment    string    `gorm:"column:table_comment;type:varchar(255)"`         // 表注释
	IsSystem        int8      `gorm:"column:is_system;type:tinyint(1)"`               // 是否系统表
	TableCategory   string    `gorm:"column:table_category;type:varchar(50)"`         // 表分类
	TableMode       int8      `gorm:"column:table_mode;type:tinyint(1)"`              // 表模式：1-主表、2子表、3-扩展表
	ParentTableID   int64     `gorm:"column:parent_table_id;type:bigint(20)"`         // 父表Id
	ParentTableName string    `gorm:"column:parent_table_name;type:varchar(100)"`     // 父表名称
	Status          int8      `gorm:"column:status;type:tinyint(1)"`                  // 表状态： 1-新建未同步 2-修改未同步 3-删除未同步 10-已同步
	Enabled         int8      `gorm:"column:enabled;type:tinyint(1)"`                 // 启用状态： 1-启用 2-禁用
	CreateTime      time.Time `gorm:"column:create_time;type:datetime"`               // 创建时间
	CreateUser      string    `gorm:"column:create_user;type:varchar(50)"`            // 创建人
	UpdateTime      time.Time `gorm:"column:update_time;type:datetime"`               // 更新时间
	UpdateUser      string    `gorm:"column:update_user;type:varchar(50)"`            // 更新人
	Ts              time.Time `gorm:"column:ts;type:timestamp"`                       // 时间戳
	Yn              int8      `gorm:"column:yn;type:tinyint(1)"`                      // 是否有效， 1-正常  0-删除
}
