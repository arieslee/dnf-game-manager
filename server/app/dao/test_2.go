// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// test2Dao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type test2Dao struct {
	internal.Test2Dao
}

var (
	// Test2 is globally public accessible object for table test2 operations.
	Test2 = test2Dao{
		internal.Test2,
	}
)

// Fill with you ideas below.

func (d *test2Dao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
