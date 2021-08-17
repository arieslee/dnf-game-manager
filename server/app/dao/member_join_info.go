// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberJoinInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberJoinInfoDao struct {
	internal.MemberJoinInfoDao
}

var (
	// MemberJoinInfo is globally public accessible object for table member_join_info operations.
	MemberJoinInfo = memberJoinInfoDao{
		internal.MemberJoinInfo,
	}
)

// Fill with you ideas below.

func (d *memberJoinInfoDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
