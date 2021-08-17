// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventCharacMov1ThDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventCharacMov1ThDao struct {
	internal.EventCharacMov1ThDao
}

var (
	// EventCharacMov1Th is globally public accessible object for table event_charac_mov_1th operations.
	EventCharacMov1Th = eventCharacMov1ThDao{
		internal.EventCharacMov1Th,
	}
)

// Fill with you ideas below.

func (d *eventCharacMov1ThDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
