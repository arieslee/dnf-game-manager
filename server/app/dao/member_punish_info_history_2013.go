// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberPunishInfoHistory2013Dao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberPunishInfoHistory2013Dao struct {
	internal.MemberPunishInfoHistory2013Dao
}

var (
	// MemberPunishInfoHistory2013 is globally public accessible object for table member_punish_info_history_2013 operations.
	MemberPunishInfoHistory2013 = memberPunishInfoHistory2013Dao{
		internal.MemberPunishInfoHistory2013,
	}
)

// Fill with you ideas below.

func (d *memberPunishInfoHistory2013Dao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
