// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventCharacMov1ThEntryDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventCharacMov1ThEntryDao struct {
	internal.EventCharacMov1ThEntryDao
}

var (
	// EventCharacMov1ThEntry is globally public accessible object for table event_charac_mov_1th_entry operations.
	EventCharacMov1ThEntry = eventCharacMov1ThEntryDao{
		internal.EventCharacMov1ThEntry,
	}
)

// Fill with you ideas below.

func (d *eventCharacMov1ThEntryDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
