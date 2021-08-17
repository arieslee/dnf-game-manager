// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// dblabDbConnect130516Dao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type dblabDbConnect130516Dao struct {
	internal.DblabDbConnect130516Dao
}

var (
	// DblabDbConnect130516 is globally public accessible object for table dblab_db_connect_130516 operations.
	DblabDbConnect130516 = dblabDbConnect130516Dao{
		internal.DblabDbConnect130516,
	}
)

// Fill with you ideas below.

func (d *dblabDbConnect130516Dao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
