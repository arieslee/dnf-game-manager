// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventMage2YearsDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventMage2YearsDao struct {
	internal.EventMage2YearsDao
}

var (
	// EventMage2Years is globally public accessible object for table event_mage_2years operations.
	EventMage2Years = eventMage2YearsDao{
		internal.EventMage2Years,
	}
)

// Fill with you ideas below.

func (d *eventMage2YearsDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
