// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// accountssDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type accountssDao struct {
	internal.AccountssDao
}

var (
	// Accountss is globally public accessible object for table accountss operations.
	Accountss = accountssDao{
		internal.Accountss,
	}
)

// Fill with you ideas below.

func (d *accountssDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
