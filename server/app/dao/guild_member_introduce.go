// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// guildMemberIntroduceDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type guildMemberIntroduceDao struct {
	internal.GuildMemberIntroduceDao
}

var (
	// GuildMemberIntroduce is globally public accessible object for table guild_member_introduce operations.
	GuildMemberIntroduce = guildMemberIntroduceDao{
		internal.GuildMemberIntroduce,
	}
)

// Fill with you ideas below.

func (d *guildMemberIntroduceDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
