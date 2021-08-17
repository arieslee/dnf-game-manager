// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberInfoEuckrDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberInfoEuckrDao struct {
	internal.MemberInfoEuckrDao
}

var (
	// MemberInfoEuckr is globally public accessible object for table member_info_euckr operations.
	MemberInfoEuckr = memberInfoEuckrDao{
		internal.MemberInfoEuckr,
	}
)

// Fill with you ideas below.

func (d *memberInfoEuckrDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
