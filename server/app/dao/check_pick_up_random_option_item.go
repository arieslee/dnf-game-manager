// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// checkPickUpRandomOptionItemDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type checkPickUpRandomOptionItemDao struct {
	internal.CheckPickUpRandomOptionItemDao
}

var (
	// CheckPickUpRandomOptionItem is globally public accessible object for table check_pick_up_random_option_item operations.
	CheckPickUpRandomOptionItem = checkPickUpRandomOptionItemDao{
		internal.CheckPickUpRandomOptionItem,
	}
)

// Fill with you ideas below.

func (d *checkPickUpRandomOptionItemDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
