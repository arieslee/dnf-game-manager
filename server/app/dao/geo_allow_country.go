// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// geoAllowCountryDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type geoAllowCountryDao struct {
	internal.GeoAllowCountryDao
}

var (
	// GeoAllowCountry is globally public accessible object for table geo_allow_country operations.
	GeoAllowCountry = geoAllowCountryDao{
		internal.GeoAllowCountry,
	}
)

// Fill with you ideas below.

func (d *geoAllowCountryDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
