// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// dnfGameMessageDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type dnfGameMessageDao struct {
	internal.DnfGameMessageDao
}

var (
	// DnfGameMessage is globally public accessible object for table dnf_game_message operations.
	DnfGameMessage = dnfGameMessageDao{
		internal.DnfGameMessage,
	}
)

// Fill with you ideas below.

func (d *dnfGameMessageDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
