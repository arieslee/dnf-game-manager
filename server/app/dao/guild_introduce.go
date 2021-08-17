// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// guildIntroduceDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type guildIntroduceDao struct {
	internal.GuildIntroduceDao
}

var (
	// GuildIntroduce is globally public accessible object for table guild_introduce operations.
	GuildIntroduce = guildIntroduceDao{
		internal.GuildIntroduce,
	}
)

// Fill with you ideas below.

func (d *guildIntroduceDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
