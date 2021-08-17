// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventGoldcardInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventGoldcardInfoDao struct {
	internal.EventGoldcardInfoDao
}

var (
	// EventGoldcardInfo is globally public accessible object for table event_goldcard_info operations.
	EventGoldcardInfo = eventGoldcardInfoDao{
		internal.EventGoldcardInfo,
	}
)

// Fill with you ideas below.

func (d *eventGoldcardInfoDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
