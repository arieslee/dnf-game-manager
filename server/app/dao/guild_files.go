// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// guildFilesDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type guildFilesDao struct {
	internal.GuildFilesDao
}

var (
	// GuildFiles is globally public accessible object for table guild_files operations.
	GuildFiles = guildFilesDao{
		internal.GuildFiles,
	}
)

// Fill with you ideas below.

func (d *guildFilesDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
