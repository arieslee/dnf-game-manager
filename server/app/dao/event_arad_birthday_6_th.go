// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventAradBirthday6ThDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventAradBirthday6ThDao struct {
	internal.EventAradBirthday6ThDao
}

var (
	// EventAradBirthday6Th is globally public accessible object for table event_arad_birthday_6th operations.
	EventAradBirthday6Th = eventAradBirthday6ThDao{
		internal.EventAradBirthday6Th,
	}
)

// Fill with you ideas below.

func (d *eventAradBirthday6ThDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
