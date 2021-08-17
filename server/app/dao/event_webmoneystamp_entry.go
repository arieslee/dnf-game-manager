// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventWebmoneystampEntryDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventWebmoneystampEntryDao struct {
	internal.EventWebmoneystampEntryDao
}

var (
	// EventWebmoneystampEntry is globally public accessible object for table event_webmoneystamp_entry operations.
	EventWebmoneystampEntry = eventWebmoneystampEntryDao{
		internal.EventWebmoneystampEntry,
	}
)

// Fill with you ideas below.

func (d *eventWebmoneystampEntryDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
