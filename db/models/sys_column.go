package models

import "time"

type SysColumn struct {
	ID                 uint8     `gorm:"primary_key" json:" - "` //主键
	TableId            uint8     `json:"table_id"`               //表ID
	TableName          string    `json:"table_name"`             //表名
	ColumnName         string    `json:"column_name"`            //列名
	ColumnComment      string    `json:"column_comment"`         //列注释
	ColumnDefault      string    `json:"column_default"`         //默认值
	ColumnType         string    `json:"column_type"`            //列类型
	PrimaryKey         string    `json:"primary_key"`            //是否主键 1-是 2-否
	ColumnProperty     uint8     `json:"column_property"`        //列属性，1-固定字段，2-扩展字段
	ColumnLength       int       `json:"column_length"`          //列长度
	ColumnPrecision    int       `json:"column_precision"`       //列精度
	Nullable           uint8     `json:"nullable"`               //是否为空
	RefTableId         uint8     `json:"ref_table_id"`           //引用的表ID
	RefTableName       string    `json:"ref_table_name"`         //引用的表名
	RefColumnId        uint8     `json:"ref_column_id"`          //引用表的字段ID
	RefColumnName      string    `json:"ref_column_name"`        //引用表的字段名
	RefColumnValueId   string    `json:"ref_column_value_id"`    //引用表的值的字段Id，多个
	RefColumnValueName string    `json:"ref_column_value_name"`  //引用表的值的字段名，多个
	RefDictCategory    string    `json:"ref_dict_category"`      //引用字典的类型
	RefDictCode        string    `json:"ref_dict_code"`          //引用字典的编码
	CreateTime         time.Time `json:"create_time"`            //创建时间
	CreateUser         string    `json:"create_user"`            //创建人
	UpdateTime         time.Time `json:"update_time"`            //更新时间
	UpdateUser         string    `json:"update_user"`            //更新人
	Ts                 time.Time `json:"ts"`                     //时间戳
	Yn                 uint8     `json:"yn"`                     //是否有效， 1-正常  0-删除
}
