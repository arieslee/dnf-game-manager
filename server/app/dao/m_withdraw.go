// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// mWithdrawDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type mWithdrawDao struct {
	internal.MWithdrawDao
}

var (
	// MWithdraw is globally public accessible object for table m_withdraw operations.
	MWithdraw = mWithdrawDao{
		internal.MWithdraw,
	}
)

// Fill with you ideas below.

func (d *mWithdrawDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
