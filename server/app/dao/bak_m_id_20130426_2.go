// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// bakMId201304262Dao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type bakMId201304262Dao struct {
	internal.BakMId201304262Dao
}

var (
	// BakMId201304262 is globally public accessible object for table bak_m_id_20130426_2 operations.
	BakMId201304262 = bakMId201304262Dao{
		internal.BakMId201304262,
	}
)

// Fill with you ideas below.

func (d *bakMId201304262Dao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
