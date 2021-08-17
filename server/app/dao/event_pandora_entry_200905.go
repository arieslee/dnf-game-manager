// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventPandoraEntry200905Dao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventPandoraEntry200905Dao struct {
	internal.EventPandoraEntry200905Dao
}

var (
	// EventPandoraEntry200905 is globally public accessible object for table event_pandora_entry_200905 operations.
	EventPandoraEntry200905 = eventPandoraEntry200905Dao{
		internal.EventPandoraEntry200905,
	}
)

// Fill with you ideas below.

func (d *eventPandoraEntry200905Dao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
