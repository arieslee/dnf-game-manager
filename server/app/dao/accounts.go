// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// accountsDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type accountsDao struct {
	internal.AccountsDao
}

var (
	// Accounts is globally public accessible object for table accounts operations.
	Accounts = accountsDao{
		internal.Accounts,
	}
)

// Fill with you ideas below.

func (d *accountsDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
