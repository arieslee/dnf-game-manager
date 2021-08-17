// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// guildBodyDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type guildBodyDao struct {
	internal.GuildBodyDao
}

var (
	// GuildBody is globally public accessible object for table guild_body operations.
	GuildBody = guildBodyDao{
		internal.GuildBody,
	}
)

// Fill with you ideas below.

func (d *guildBodyDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
