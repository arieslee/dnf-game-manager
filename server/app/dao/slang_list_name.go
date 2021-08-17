// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// slangListNameDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type slangListNameDao struct {
	internal.SlangListNameDao
}

var (
	// SlangListName is globally public accessible object for table slang_list_name operations.
	SlangListName = slangListNameDao{
		internal.SlangListName,
	}
)

// Fill with you ideas below.

func (d *slangListNameDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
