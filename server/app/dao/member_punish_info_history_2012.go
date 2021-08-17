// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberPunishInfoHistory2012Dao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberPunishInfoHistory2012Dao struct {
	internal.MemberPunishInfoHistory2012Dao
}

var (
	// MemberPunishInfoHistory2012 is globally public accessible object for table member_punish_info_history_2012 operations.
	MemberPunishInfoHistory2012 = memberPunishInfoHistory2012Dao{
		internal.MemberPunishInfoHistory2012,
	}
)

// Fill with you ideas below.

func (d *memberPunishInfoHistory2012Dao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
