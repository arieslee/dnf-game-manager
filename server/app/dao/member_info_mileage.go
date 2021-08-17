// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberInfoMileageDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberInfoMileageDao struct {
	internal.MemberInfoMileageDao
}

var (
	// MemberInfoMileage is globally public accessible object for table member_info_mileage operations.
	MemberInfoMileage = memberInfoMileageDao{
		internal.MemberInfoMileage,
	}
)

// Fill with you ideas below.

func (d *memberInfoMileageDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
