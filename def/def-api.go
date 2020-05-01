package def

type TableDefService interface {
	QueryAllTableDef() []*TableDef
	FindTableDef(tableName string) (*TableDef, error)
}

type ColumnDefService interface {
	QueryAllColumnDef(tableName string) []ColumnDef
	FindColumnDef(tableName string, columnName string) (TableDef, error)
}
