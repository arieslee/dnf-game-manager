// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberSecurityGradeDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberSecurityGradeDao struct {
	internal.MemberSecurityGradeDao
}

var (
	// MemberSecurityGrade is globally public accessible object for table member_security_grade operations.
	MemberSecurityGrade = memberSecurityGradeDao{
		internal.MemberSecurityGrade,
	}
)

// Fill with you ideas below.

func (d *memberSecurityGradeDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
