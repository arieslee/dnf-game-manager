// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// bakDnfItemInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type bakDnfItemInfoDao struct {
	internal.BakDnfItemInfoDao
}

var (
	// BakDnfItemInfo is globally public accessible object for table bak_dnf_item_info operations.
	BakDnfItemInfo = bakDnfItemInfoDao{
		internal.BakDnfItemInfo,
	}
)

// Fill with you ideas below.

func (d *bakDnfItemInfoDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
