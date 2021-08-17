// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberWhiteAccountDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberWhiteAccountDao struct {
	internal.MemberWhiteAccountDao
}

var (
	// MemberWhiteAccount is globally public accessible object for table member_white_account operations.
	MemberWhiteAccount = memberWhiteAccountDao{
		internal.MemberWhiteAccount,
	}
)

// Fill with you ideas below.

func (d *memberWhiteAccountDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
