// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// maxCountPvpDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type maxCountPvpDao struct {
	internal.MaxCountPvpDao
}

var (
	// MaxCountPvp is globally public accessible object for table max_count_pvp operations.
	MaxCountPvp = maxCountPvpDao{
		internal.MaxCountPvp,
	}
)

// Fill with you ideas below.

func (d *maxCountPvpDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
