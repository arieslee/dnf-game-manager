// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// guildEventDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type guildEventDao struct {
	internal.GuildEventDao
}

var (
	// GuildEvent is globally public accessible object for table guild_event operations.
	GuildEvent = guildEventDao{
		internal.GuildEvent,
	}
)

// Fill with you ideas below.

func (d *guildEventDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
