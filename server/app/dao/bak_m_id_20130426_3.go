// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// bakMId201304263Dao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type bakMId201304263Dao struct {
	internal.BakMId201304263Dao
}

var (
	// BakMId201304263 is globally public accessible object for table bak_m_id_20130426_3 operations.
	BakMId201304263 = bakMId201304263Dao{
		internal.BakMId201304263,
	}
)

// Fill with you ideas below.

func (d *bakMId201304263Dao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
