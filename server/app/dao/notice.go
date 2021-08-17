// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// noticeDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type noticeDao struct {
	internal.NoticeDao
}

var (
	// Notice is globally public accessible object for table notice operations.
	Notice = noticeDao{
		internal.Notice,
	}
)

// Fill with you ideas below.

func (d *noticeDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
