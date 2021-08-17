// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// dnfEventEntryNotuseDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type dnfEventEntryNotuseDao struct {
	internal.DnfEventEntryNotuseDao
}

var (
	// DnfEventEntryNotuse is globally public accessible object for table dnf_event_entry_notuse operations.
	DnfEventEntryNotuse = dnfEventEntryNotuseDao{
		internal.DnfEventEntryNotuse,
	}
)

// Fill with you ideas below.

func (d *dnfEventEntryNotuseDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
