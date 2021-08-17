// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// limitCreateCharacterIpDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type limitCreateCharacterIpDao struct {
	internal.LimitCreateCharacterIpDao
}

var (
	// LimitCreateCharacterIp is globally public accessible object for table limit_create_character_ip operations.
	LimitCreateCharacterIp = limitCreateCharacterIpDao{
		internal.LimitCreateCharacterIp,
	}
)

// Fill with you ideas below.

func (d *limitCreateCharacterIpDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
