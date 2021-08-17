// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberSafeEnsureDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberSafeEnsureDao struct {
	internal.MemberSafeEnsureDao
}

var (
	// MemberSafeEnsure is globally public accessible object for table member_safe_ensure operations.
	MemberSafeEnsure = memberSafeEnsureDao{
		internal.MemberSafeEnsure,
	}
)

// Fill with you ideas below.

func (d *memberSafeEnsureDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
