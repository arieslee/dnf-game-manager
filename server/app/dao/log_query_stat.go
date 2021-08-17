// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// logQueryStatDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type logQueryStatDao struct {
	internal.LogQueryStatDao
}

var (
	// LogQueryStat is globally public accessible object for table log_query_stat operations.
	LogQueryStat = logQueryStatDao{
		internal.LogQueryStat,
	}
)

// Fill with you ideas below.

func (d *logQueryStatDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
