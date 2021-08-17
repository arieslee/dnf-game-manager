// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// accountCerashopRestrictDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type accountCerashopRestrictDao struct {
	internal.AccountCerashopRestrictDao
}

var (
	// AccountCerashopRestrict is globally public accessible object for table account_cerashop_restrict operations.
	AccountCerashopRestrict = accountCerashopRestrictDao{
		internal.AccountCerashopRestrict,
	}
)

// Fill with you ideas below.

func (d *accountCerashopRestrictDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
