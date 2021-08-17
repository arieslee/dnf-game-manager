// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// passwdModEntryDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type passwdModEntryDao struct {
	internal.PasswdModEntryDao
}

var (
	// PasswdModEntry is globally public accessible object for table passwd_mod_entry operations.
	PasswdModEntry = passwdModEntryDao{
		internal.PasswdModEntry,
	}
)

// Fill with you ideas below.

func (d *passwdModEntryDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
