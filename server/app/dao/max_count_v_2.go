// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// maxCountV2Dao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type maxCountV2Dao struct {
	internal.MaxCountV2Dao
}

var (
	// MaxCountV2 is globally public accessible object for table max_count_v2 operations.
	MaxCountV2 = maxCountV2Dao{
		internal.MaxCountV2,
	}
)

// Fill with you ideas below.

func (d *maxCountV2Dao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
