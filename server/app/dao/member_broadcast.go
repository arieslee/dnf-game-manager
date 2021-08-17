// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberBroadcastDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberBroadcastDao struct {
	internal.MemberBroadcastDao
}

var (
	// MemberBroadcast is globally public accessible object for table member_broadcast operations.
	MemberBroadcast = memberBroadcastDao{
		internal.MemberBroadcast,
	}
)

// Fill with you ideas below.

func (d *memberBroadcastDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
