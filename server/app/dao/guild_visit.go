// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// guildVisitDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type guildVisitDao struct {
	internal.GuildVisitDao
}

var (
	// GuildVisit is globally public accessible object for table guild_visit operations.
	GuildVisit = guildVisitDao{
		internal.GuildVisit,
	}
)

// Fill with you ideas below.

func (d *guildVisitDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
