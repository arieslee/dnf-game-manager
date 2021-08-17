// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberInfoUtf8Dao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberInfoUtf8Dao struct {
	internal.MemberInfoUtf8Dao
}

var (
	// MemberInfoUtf8 is globally public accessible object for table member_info_utf8 operations.
	MemberInfoUtf8 = memberInfoUtf8Dao{
		internal.MemberInfoUtf8,
	}
)

// Fill with you ideas below.

func (d *memberInfoUtf8Dao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
