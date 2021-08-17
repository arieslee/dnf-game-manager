// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventGoldcardEntry2Dao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventGoldcardEntry2Dao struct {
	internal.EventGoldcardEntry2Dao
}

var (
	// EventGoldcardEntry2 is globally public accessible object for table event_goldcard_entry2 operations.
	EventGoldcardEntry2 = eventGoldcardEntry2Dao{
		internal.EventGoldcardEntry2,
	}
)

// Fill with you ideas below.

func (d *eventGoldcardEntry2Dao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
