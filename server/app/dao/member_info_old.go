// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberInfoOldDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberInfoOldDao struct {
	internal.MemberInfoOldDao
}

var (
	// MemberInfoOld is globally public accessible object for table member_info_old operations.
	MemberInfoOld = memberInfoOldDao{
		internal.MemberInfoOld,
	}
)

// Fill with you ideas below.

func (d *memberInfoOldDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
