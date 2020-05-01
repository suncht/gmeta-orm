package def

type TableDef struct {
	TableName  string
	SourceDef  TableSourceDef
	DisplayDef TableDisplayDef
}

type TableSourceDef struct {
	TableComment string
}

type TableBizDef struct {
}

type TableDisplayDef struct {
}

type TableExtDef struct {
}

type ColumnDef struct {
	SourceDef  ColumnSourceDef
	DisplayDef ColumnDisplayDef
}

type ColumnSourceDef struct {
}

type ColumnBizDef struct {
}

type ColumnDisplayDef struct {
}

type ColumnExtDef struct {
}
