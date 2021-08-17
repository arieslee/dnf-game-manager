// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// adminMemberDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type adminMemberDao struct {
	internal.AdminMemberDao
}

var (
	// AdminMember is globally public accessible object for table admin_member operations.
	AdminMember = adminMemberDao{
		internal.AdminMember,
	}
)

// Fill with you ideas below.

func (d *adminMemberDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
