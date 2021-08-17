// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventGoldcardEntry1Dao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventGoldcardEntry1Dao struct {
	internal.EventGoldcardEntry1Dao
}

var (
	// EventGoldcardEntry1 is globally public accessible object for table event_goldcard_entry1 operations.
	EventGoldcardEntry1 = eventGoldcardEntry1Dao{
		internal.EventGoldcardEntry1,
	}
)

// Fill with you ideas below.

func (d *eventGoldcardEntry1Dao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
