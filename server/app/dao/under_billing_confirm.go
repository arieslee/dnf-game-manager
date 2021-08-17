// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// underBillingConfirmDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type underBillingConfirmDao struct {
	internal.UnderBillingConfirmDao
}

var (
	// UnderBillingConfirm is globally public accessible object for table under_billing_confirm operations.
	UnderBillingConfirm = underBillingConfirmDao{
		internal.UnderBillingConfirm,
	}
)

// Fill with you ideas below.

func (d *underBillingConfirmDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
