// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// dnfCharacMovDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type dnfCharacMovDao struct {
	internal.DnfCharacMovDao
}

var (
	// DnfCharacMov is globally public accessible object for table dnf_charac_mov operations.
	DnfCharacMov = dnfCharacMovDao{
		internal.DnfCharacMov,
	}
)

// Fill with you ideas below.

func (d *dnfCharacMovDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
