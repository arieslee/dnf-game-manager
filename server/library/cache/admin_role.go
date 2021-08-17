// Package cache
package cache

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"hess/app/dao"
	"hess/library/redis"
	"sync"
)

var AdminRole = new(adminRoleCache)
var syncAdminRoleLock *sync.Mutex

type adminRoleCache struct{}

const (
	RoleCacheKey = "admin:role:%d" // %d 为role_id
)

func (c *adminRoleCache) Detail(roleId uint64) (g.Map, error) {
	row, _ := dao.AdminRole.Fields(dao.AdminRole.Columns).WherePri(roleId).One()
	if row == nil {
		return nil, errors.New("管理员角色不存在")
	}
	roleContent, _ := dao.AdminRoleContent.Fields("content").Where("role_id=?", roleId).One()
	if roleContent == nil {
		return nil, errors.New("管理员角色权限不存在")
	}

	mapData := gconv.Map(row)
	conStr := roleContent.Content
	if conStr == "all" {
		mapData["content"] = "all"
	} else {
		mapData["content"] = conStr
	}
	return mapData, nil
}

func (c *adminRoleCache) GetMapStrStr(roleId uint64) map[string]string {
	key := fmt.Sprintf(RoleCacheKey, roleId)
	data, _ := redis.GetInstance().HGetMap(key)
	if data == nil || len(data) <= 0 {
		if syncAdminRoleLock != nil {
			syncAdminRoleLock.Lock()
			defer syncAdminRoleLock.Unlock()
		}
		roleData, _ := c.Detail(roleId)
		if roleData == nil {
			return nil
		}
		cacheData := gconv.MapStrStr(roleData)
		_ = redis.GetInstance().HSetMap(key, cacheData)
		return cacheData
	}
	return data
}
func (c *adminRoleCache) Remove(roleId uint64) {
	key := fmt.Sprintf(RoleCacheKey, roleId)
	_, _ = redis.GetInstance().Del(key)
}
func (c *adminRoleCache) DeleteAll() {
	all, _ := dao.AdminRole.Fields(dao.AdminRole.Columns.Id).All()
	if len(all) > 0 {
		for _, v := range all {
			c.Remove(v.Id)
		}
	}
}
