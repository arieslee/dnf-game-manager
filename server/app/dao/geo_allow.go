// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// geoAllowDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type geoAllowDao struct {
	internal.GeoAllowDao
}

var (
	// GeoAllow is globally public accessible object for table geo_allow operations.
	GeoAllow = geoAllowDao{
		internal.GeoAllow,
	}
)

// Fill with you ideas below.

func (d *geoAllowDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
