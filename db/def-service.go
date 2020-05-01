package db

import (
	"gmeta-orm/db/models"
	"gmeta-orm/def"
)

type DefaultTableDefService struct {
}

func (server *DefaultTableDefService) QueryAllTableDef() []*def.TableDef {
	var sysTables []models.SysTable
	DB.Debug().
		Find(&sysTables)

	tableDefs := make([]*def.TableDef, 0, len(sysTables))

	for i := 0; i < len(sysTables); i++ {
		tableDefs = append(tableDefs, &def.TableDef{
			TableName: sysTables[i].TableName,
			SourceDef: def.TableSourceDef{
				TableComment: sysTables[i].TableComment,
			},
		})
	}

	return tableDefs
}

func (server *DefaultTableDefService) FindTableDef(tableName string) (*def.TableDef, error) {
	var sysTable models.SysTable
	err := DB.Debug().
		Where("table_name = ?", tableName).
		Table("sys_table").
		First(&sysTable).Error

	if err != nil {
		return nil, err
	}

	tableDef := &def.TableDef{
		TableName: sysTable.TableName,
		SourceDef: def.TableSourceDef{
			TableComment: sysTable.TableComment,
		},
	}

	return tableDef, nil
}
