// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// limitCreateCharacterDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type limitCreateCharacterDao struct {
	internal.LimitCreateCharacterDao
}

var (
	// LimitCreateCharacter is globally public accessible object for table limit_create_character operations.
	LimitCreateCharacter = limitCreateCharacterDao{
		internal.LimitCreateCharacter,
	}
)

// Fill with you ideas below.

func (d *limitCreateCharacterDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
