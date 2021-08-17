// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventNewmember0709EntryDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventNewmember0709EntryDao struct {
	internal.EventNewmember0709EntryDao
}

var (
	// EventNewmember0709Entry is globally public accessible object for table event_newmember0709_entry operations.
	EventNewmember0709Entry = eventNewmember0709EntryDao{
		internal.EventNewmember0709Entry,
	}
)

// Fill with you ideas below.

func (d *eventNewmember0709EntryDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
