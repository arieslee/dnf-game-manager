// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// guildStatMonthDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type guildStatMonthDao struct {
	internal.GuildStatMonthDao
}

var (
	// GuildStatMonth is globally public accessible object for table guild_stat_month operations.
	GuildStatMonth = guildStatMonthDao{
		internal.GuildStatMonth,
	}
)

// Fill with you ideas below.

func (d *guildStatMonthDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
