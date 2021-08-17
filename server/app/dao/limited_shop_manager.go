// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// limitedShopManagerDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type limitedShopManagerDao struct {
	internal.LimitedShopManagerDao
}

var (
	// LimitedShopManager is globally public accessible object for table limited_shop_manager operations.
	LimitedShopManager = limitedShopManagerDao{
		internal.LimitedShopManager,
	}
)

// Fill with you ideas below.

func (d *limitedShopManagerDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
