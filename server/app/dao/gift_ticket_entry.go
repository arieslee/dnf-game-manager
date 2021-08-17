// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// giftTicketEntryDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type giftTicketEntryDao struct {
	internal.GiftTicketEntryDao
}

var (
	// GiftTicketEntry is globally public accessible object for table gift_ticket_entry operations.
	GiftTicketEntry = giftTicketEntryDao{
		internal.GiftTicketEntry,
	}
)

// Fill with you ideas below.

func (d *giftTicketEntryDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
