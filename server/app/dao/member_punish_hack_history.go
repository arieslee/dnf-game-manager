// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberPunishHackHistoryDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberPunishHackHistoryDao struct {
	internal.MemberPunishHackHistoryDao
}

var (
	// MemberPunishHackHistory is globally public accessible object for table member_punish_hack_history operations.
	MemberPunishHackHistory = memberPunishHackHistoryDao{
		internal.MemberPunishHackHistory,
	}
)

// Fill with you ideas below.

func (d *memberPunishHackHistoryDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
