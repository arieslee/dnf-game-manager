// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// guildBbsDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type guildBbsDao struct {
	internal.GuildBbsDao
}

var (
	// GuildBbs is globally public accessible object for table guild_bbs operations.
	GuildBbs = guildBbsDao{
		internal.GuildBbs,
	}
)

// Fill with you ideas below.

func (d *guildBbsDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
