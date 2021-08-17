// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventSkill2025EntryDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventSkill2025EntryDao struct {
	internal.EventSkill2025EntryDao
}

var (
	// EventSkill2025Entry is globally public accessible object for table event_skill2025_entry operations.
	EventSkill2025Entry = eventSkill2025EntryDao{
		internal.EventSkill2025Entry,
	}
)

// Fill with you ideas below.

func (d *eventSkill2025EntryDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
