// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventTowerEntryDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventTowerEntryDao struct {
	internal.EventTowerEntryDao
}

var (
	// EventTowerEntry is globally public accessible object for table event_tower_entry operations.
	EventTowerEntry = eventTowerEntryDao{
		internal.EventTowerEntry,
	}
)

// Fill with you ideas below.

func (d *eventTowerEntryDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
