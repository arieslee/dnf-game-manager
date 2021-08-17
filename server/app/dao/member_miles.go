// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberMilesDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberMilesDao struct {
	internal.MemberMilesDao
}

var (
	// MemberMiles is globally public accessible object for table member_miles operations.
	MemberMiles = memberMilesDao{
		internal.MemberMiles,
	}
)

// Fill with you ideas below.

func (d *memberMilesDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
