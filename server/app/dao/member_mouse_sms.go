// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberMouseSmsDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberMouseSmsDao struct {
	internal.MemberMouseSmsDao
}

var (
	// MemberMouseSms is globally public accessible object for table member_mouse_sms operations.
	MemberMouseSms = memberMouseSmsDao{
		internal.MemberMouseSms,
	}
)

// Fill with you ideas below.

func (d *memberMouseSmsDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
