// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberPremiumNotuseDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberPremiumNotuseDao struct {
	internal.MemberPremiumNotuseDao
}

var (
	// MemberPremiumNotuse is globally public accessible object for table member_premium_notuse operations.
	MemberPremiumNotuse = memberPremiumNotuseDao{
		internal.MemberPremiumNotuse,
	}
)

// Fill with you ideas below.

func (d *memberPremiumNotuseDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}