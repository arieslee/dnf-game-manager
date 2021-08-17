// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventWebmoneystampItemDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventWebmoneystampItemDao struct {
	internal.EventWebmoneystampItemDao
}

var (
	// EventWebmoneystampItem is globally public accessible object for table event_webmoneystamp_item operations.
	EventWebmoneystampItem = eventWebmoneystampItemDao{
		internal.EventWebmoneystampItem,
	}
)

// Fill with you ideas below.

func (d *eventWebmoneystampItemDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
