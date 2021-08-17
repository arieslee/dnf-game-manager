// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// pswdQstionDirectDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type pswdQstionDirectDao struct {
	internal.PswdQstionDirectDao
}

var (
	// PswdQstionDirect is globally public accessible object for table pswd_qstion_direct operations.
	PswdQstionDirect = pswdQstionDirectDao{
		internal.PswdQstionDirect,
	}
)

// Fill with you ideas below.

func (d *pswdQstionDirectDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
