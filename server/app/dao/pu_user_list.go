// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// puUserListDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type puUserListDao struct {
	internal.PuUserListDao
}

var (
	// PuUserList is globally public accessible object for table pu_user_list operations.
	PuUserList = puUserListDao{
		internal.PuUserList,
	}
)

// Fill with you ideas below.

func (d *puUserListDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
