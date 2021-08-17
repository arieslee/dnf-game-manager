// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// chStatusDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type chStatusDao struct {
	internal.ChStatusDao
}

var (
	// ChStatus is globally public accessible object for table ch_status operations.
	ChStatus = chStatusDao{
		internal.ChStatus,
	}
)

// Fill with you ideas below.

func (d *chStatusDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
