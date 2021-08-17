// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// gmManifestNotuseDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type gmManifestNotuseDao struct {
	internal.GmManifestNotuseDao
}

var (
	// GmManifestNotuse is globally public accessible object for table gm_manifest_notuse operations.
	GmManifestNotuse = gmManifestNotuseDao{
		internal.GmManifestNotuse,
	}
)

// Fill with you ideas below.

func (d *gmManifestNotuseDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
