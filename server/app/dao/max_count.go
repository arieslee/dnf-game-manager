// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// maxCountDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type maxCountDao struct {
	internal.MaxCountDao
}

var (
	// MaxCount is globally public accessible object for table max_count operations.
	MaxCount = maxCountDao{
		internal.MaxCount,
	}
)

// Fill with you ideas below.

func (d *maxCountDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
