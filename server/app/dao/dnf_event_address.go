// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// dnfEventAddressDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type dnfEventAddressDao struct {
	internal.DnfEventAddressDao
}

var (
	// DnfEventAddress is globally public accessible object for table dnf_event_address operations.
	DnfEventAddress = dnfEventAddressDao{
		internal.DnfEventAddress,
	}
)

// Fill with you ideas below.

func (d *dnfEventAddressDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
