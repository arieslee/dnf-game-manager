// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventAradlotto0809EntryDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventAradlotto0809EntryDao struct {
	internal.EventAradlotto0809EntryDao
}

var (
	// EventAradlotto0809Entry is globally public accessible object for table event_aradlotto_0809_entry operations.
	EventAradlotto0809Entry = eventAradlotto0809EntryDao{
		internal.EventAradlotto0809Entry,
	}
)

// Fill with you ideas below.

func (d *eventAradlotto0809EntryDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
