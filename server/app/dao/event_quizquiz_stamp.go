// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventQuizquizStampDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventQuizquizStampDao struct {
	internal.EventQuizquizStampDao
}

var (
	// EventQuizquizStamp is globally public accessible object for table event_quizquiz_stamp operations.
	EventQuizquizStamp = eventQuizquizStampDao{
		internal.EventQuizquizStamp,
	}
)

// Fill with you ideas below.

func (d *eventQuizquizStampDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
