// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// dbConnectDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type dbConnectDao struct {
	internal.DbConnectDao
}

var (
	// DbConnect is globally public accessible object for table db_connect operations.
	DbConnect = dbConnectDao{
		internal.DbConnect,
	}
)

// Fill with you ideas below.

func (d *dbConnectDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
