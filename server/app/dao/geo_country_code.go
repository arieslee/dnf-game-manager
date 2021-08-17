// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// geoCountryCodeDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type geoCountryCodeDao struct {
	internal.GeoCountryCodeDao
}

var (
	// GeoCountryCode is globally public accessible object for table geo_country_code operations.
	GeoCountryCode = geoCountryCodeDao{
		internal.GeoCountryCode,
	}
)

// Fill with you ideas below.

func (d *geoCountryCodeDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
