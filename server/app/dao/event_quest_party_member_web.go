// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// eventQuestPartyMemberWebDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type eventQuestPartyMemberWebDao struct {
	internal.EventQuestPartyMemberWebDao
}

var (
	// EventQuestPartyMemberWeb is globally public accessible object for table event_quest_party_member_web operations.
	EventQuestPartyMemberWeb = eventQuestPartyMemberWebDao{
		internal.EventQuestPartyMemberWeb,
	}
)

// Fill with you ideas below.

func (d *eventQuestPartyMemberWebDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
