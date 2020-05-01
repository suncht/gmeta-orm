package def

import (
	"encoding/json"
	"errors"
	"github.com/mohae/deepcopy"
	"sync"
)

type TableDefLoader struct {
	lock           sync.Mutex
	TableDefCache  map[string]*TableDef
	ColumnDefCache map[string][]*ColumnDef

	tableDefService  TableDefService
	columnDefService ColumnDefService
}

func NewLoader() *TableDefLoader {
	return &TableDefLoader{
		TableDefCache:    make(map[string]*TableDef, 0),
		ColumnDefCache:   make(map[string][]*ColumnDef, 0),
		tableDefService:  nil,
		columnDefService: nil,
	}
}

func (loader *TableDefLoader) SetTableDefService(service TableDefService) {
	loader.tableDefService = service
}

func (loader *TableDefLoader) SetColumnDefService(service ColumnDefService) {
	loader.columnDefService = service
}

func (loader *TableDefLoader) Load(tableName string) error {
	loader.lock.Lock()
	defer loader.lock.Unlock()

	if _, ok := loader.TableDefCache[tableName]; ok {
		return nil
	}

	if loader.tableDefService == nil {
		return errors.New("haven't set tableDefService")
	}

	tableDef, err := loader.tableDefService.FindTableDef(tableName)
	if err != nil {
		return err
	}

	loader.TableDefCache[tableDef.TableName] = tableDef
	return nil
}

func (loader *TableDefLoader) Get(tableName string) (*TableDef, error) {
	if tableDef, ok := loader.TableDefCache[tableName]; ok {
		newTableDef := deepcopy.Copy(*tableDef).(TableDef)
		return &newTableDef, nil
	}

	err := loader.Load(tableName)
	if err != nil {
		return nil, err
	}

	return loader.TableDefCache[tableName], nil
}

func (loader *TableDefLoader) deepCopy(src *TableDef) (*TableDef, error) {
	var dst TableDef
	b, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, dst)
	return &dst, err
}
