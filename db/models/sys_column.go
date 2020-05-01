package models

import "time"

// SysColumn 列定义
type SysColumn struct {
	ID                 int64     `gorm:"primary_key;column:id;type:bigint(20);not null"` // 主键
	TableID            int64     `gorm:"column:table_id;type:bigint(20)"`                // 表ID
	TableName          string    `gorm:"column:table_name;type:varchar(50)"`             // 表名
	ColumnName         string    `gorm:"column:column_name;type:varchar(255)"`           // 列名
	ColumnComment      string    `gorm:"column:column_comment;type:varchar(255)"`        // 列注释
	ColumnDefault      string    `gorm:"column:column_default;type:varchar(255)"`        // 默认值
	ColumnType         string    `gorm:"column:column_type;type:varchar(255)"`           // 列类型
	PrimaryKey         string    `gorm:"column:primary_key;type:varchar(255)"`           // 是否主键 1-是 2-否
	ColumnProperty     int8      `gorm:"column:column_property;type:tinyint(1)"`         // 列属性，1-固定字段，2-扩展字段
	ColumnLength       int       `gorm:"column:column_length;type:int(11)"`              // 列长度
	ColumnPrecision    int       `gorm:"column:column_precision;type:int(11)"`           // 列精度
	Nullable           int8      `gorm:"column:nullable;type:tinyint(1)"`                // 是否为空
	RefTableID         int64     `gorm:"column:ref_table_id;type:bigint(20)"`            // 引用的表ID
	RefTableName       string    `gorm:"column:ref_table_name;type:varchar(100)"`        // 引用的表名
	RefColumnID        int64     `gorm:"column:ref_column_id;type:bigint(20)"`           // 引用表的字段ID
	RefColumnName      string    `gorm:"column:ref_column_name;type:varchar(50)"`        // 引用表的字段名
	RefColumnValueID   string    `gorm:"column:ref_column_value_id;type:varchar(50)"`    // 引用表的值的字段Id，多个
	RefColumnValueName string    `gorm:"column:ref_column_value_name;type:varchar(50)"`  // 引用表的值的字段名，多个
	RefDictCategory    string    `gorm:"column:ref_dict_category;type:varchar(50)"`      // 引用字典的类型
	RefDictCode        string    `gorm:"column:ref_dict_code;type:varchar(50)"`          // 引用字典的编码
	CreateTime         time.Time `gorm:"column:create_time;type:datetime"`               // 创建时间
	CreateUser         string    `gorm:"column:create_user;type:varchar(50)"`            // 创建人
	UpdateTime         time.Time `gorm:"column:update_time;type:datetime"`               // 更新时间
	UpdateUser         string    `gorm:"column:update_user;type:varchar(50)"`            // 更新人
	Ts                 time.Time `gorm:"column:ts;type:timestamp"`                       // 时间戳
	Yn                 int8      `gorm:"column:yn;type:tinyint(1)"`                      // 是否有效， 1-正常  0-删除
}
