// Package dao
package dao

import (
	"dnf-game-manager/app/dao/internal"
	"dnf-game-manager/library/helper"
	"github.com/gogf/gf/util/gconv"
)

// dnfStoryDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type dnfStoryDao struct {
	internal.DnfStoryDao
}

var (
	// DnfStory is globally public accessible object for table dnf_story operations.
	DnfStory = dnfStoryDao{
		internal.DnfStory,
	}
)

// Fill with you ideas below.

func (d *dnfStoryDao) GetColumns(prefix ...string) string {
	data := gconv.MapStrStr(d.Columns)
	return helper.FormatSelectColumn(data, prefix...)
}
