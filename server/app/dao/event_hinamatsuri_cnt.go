// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventHinamatsuriCntDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventHinamatsuriCntDao struct {
	internal.EventHinamatsuriCntDao
}

var (
	// EventHinamatsuriCnt is globally public accessible object for table event_hinamatsuri_cnt operations.
	EventHinamatsuriCnt = eventHinamatsuriCntDao{
		internal.EventHinamatsuriCnt,
	}
)

// Fill with you ideas below.

func (d *eventHinamatsuriCntDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
