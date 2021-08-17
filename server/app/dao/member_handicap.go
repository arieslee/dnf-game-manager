// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberHandicapDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberHandicapDao struct {
	internal.MemberHandicapDao
}

var (
	// MemberHandicap is globally public accessible object for table member_handicap operations.
	MemberHandicap = memberHandicapDao{
		internal.MemberHandicap,
	}
)

// Fill with you ideas below.

func (d *memberHandicapDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
