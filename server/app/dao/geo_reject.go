// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// geoRejectDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type geoRejectDao struct {
	internal.GeoRejectDao
}

var (
	// GeoReject is globally public accessible object for table geo_reject operations.
	GeoReject = geoRejectDao{
		internal.GeoReject,
	}
)

// Fill with you ideas below.

func (d *geoRejectDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
