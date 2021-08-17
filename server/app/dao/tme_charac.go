// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// tmeCharacDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type tmeCharacDao struct {
	internal.TmeCharacDao
}

var (
	// TmeCharac is globally public accessible object for table tme_charac operations.
	TmeCharac = tmeCharacDao{
		internal.TmeCharac,
	}
)

// Fill with you ideas below.

func (d *tmeCharacDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
