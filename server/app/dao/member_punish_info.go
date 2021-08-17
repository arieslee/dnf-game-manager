// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberPunishInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberPunishInfoDao struct {
	internal.MemberPunishInfoDao
}

var (
	// MemberPunishInfo is globally public accessible object for table member_punish_info operations.
	MemberPunishInfo = memberPunishInfoDao{
		internal.MemberPunishInfo,
	}
)

// Fill with you ideas below.

func (d *memberPunishInfoDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
