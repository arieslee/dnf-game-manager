// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// memberCaptchaInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type memberCaptchaInfoDao struct {
	internal.MemberCaptchaInfoDao
}

var (
	// MemberCaptchaInfo is globally public accessible object for table member_captcha_info operations.
	MemberCaptchaInfo = memberCaptchaInfoDao{
		internal.MemberCaptchaInfo,
	}
)

// Fill with you ideas below.

func (d *memberCaptchaInfoDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
