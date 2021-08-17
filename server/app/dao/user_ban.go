// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// userBanDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type userBanDao struct {
	internal.UserBanDao
}

var (
	// UserBan is globally public accessible object for table user_ban operations.
	UserBan = userBanDao{
		internal.UserBan,
	}
)

// Fill with you ideas below.

func (d *userBanDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
