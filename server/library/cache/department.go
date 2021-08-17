package cache

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"hess/app/common"
	"hess/app/dao"
	"hess/app/model"
	"hess/library/redis"
	"sync"
)

var Department = new(departmentCache)
var syncDeptLock *sync.Mutex

type departmentCache struct{}

const (
	ActiveDeptCacheKey     = "department:active"
	ActiveDeptItemCacheKey = "department:active:item:%d"
)

func (c *departmentCache) SetList(listValue []*model.Department) {
	if listValue == nil || len(listValue) == 0 {
		return
	}
	data, _ := gjson.Encode(listValue)
	strData := gconv.String(data)
	_, _ = redis.GetInstance().Set(ActiveDeptCacheKey, strData)
}

func (c *departmentCache) GetList() []*model.Department {
	cacheData, _ := redis.GetInstance().GetString(ActiveDeptCacheKey)
	if len(cacheData) == 0 {
		syncDeptLock = new(sync.Mutex)
		syncDeptLock.Lock()
		defer syncDeptLock.Unlock()
		// 从数据库中取
		list, _ := dao.Department.Where(dao.Department.Columns.Status + "=1").Order(dao.Department.Columns.ListOrder + " ASC,id DESC").All()
		if list == nil {
			return nil
		}
		c.SetList(list)
		return list
	}
	var result []*model.Department
	_ = gjson.DecodeTo(cacheData, &result)
	return result
}

// SetItem id 如果从钉钉获取数据就对应钉钉中的dept_id,否则自己生成
func (c *departmentCache) SetItem(id uint64, value interface{}) {
	key := fmt.Sprintf(ActiveDeptItemCacheKey, id)
	dataMap := gconv.MapStrStr(value)
	_ = redis.GetInstance().HSetMap(key, dataMap)
}

// GetItem id 如果从钉钉获取数据就对应钉钉中的dept_id,否则自己生成
func (c *departmentCache) GetItem(id uint64) *model.Department {
	key := fmt.Sprintf(ActiveDeptItemCacheKey, id)
	cacheData, _ := redis.GetInstance().HGetMap(key)
	if cacheData == nil || len(cacheData) == 0 {
		row, _ := dao.Department.Fields(dao.Department.GetColumns("", "created_at", "updated_at", "list_order", "status")).Where("id=? AND status=1", id).One()
		if row == nil {
			return nil
		}
		c.SetItem(id, row)
		return row
	}
	item := &model.Department{}
	_ = gconv.Struct(cacheData, &item)
	return item
}
func (c *departmentCache) DeleteById(id uint64) {
	key := fmt.Sprintf(ActiveDeptItemCacheKey, id)
	_, _ = redis.GetInstance().Del(key)
}
func (c *departmentCache) Delete() {
	list := c.GetList()
	if len(list) > 0 {
		for _, row := range list {
			c.DeleteById(row.Id)
		}
	}
	_, _ = redis.GetInstance().Del(ActiveDeptCacheKey)
}

func (c *departmentCache) DeleteItem(id uint64) {
	key := fmt.Sprintf(ActiveDeptItemCacheKey, id)
	_, _ = redis.GetInstance().Del(key)
}
func (c *departmentCache) QueryCacheDelete() {
	_, _ = g.DB().GetCache().UpdateExpire(common.BackendDepartmentAll, 0)
}
