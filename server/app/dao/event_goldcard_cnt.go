// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventGoldcardCntDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventGoldcardCntDao struct {
	internal.EventGoldcardCntDao
}

var (
	// EventGoldcardCnt is globally public accessible object for table event_goldcard_cnt operations.
	EventGoldcardCnt = eventGoldcardCntDao{
		internal.EventGoldcardCnt,
	}
)

// Fill with you ideas below.

func (d *eventGoldcardCntDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
