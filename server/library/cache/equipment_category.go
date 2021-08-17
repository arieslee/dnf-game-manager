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

var syncLock *sync.Mutex
var EquipmentCategory = new(equipmentCategoryCache)

type equipmentCategoryCache struct{}

const (
	ActiveCacheKey     = "equipment:category:active"
	ActiveItemCacheKey = "eqi:cate:active:item:%d"
)

func (c *equipmentCategoryCache) SetList(listValue []*model.EquipmentCategory) {
	if listValue == nil || len(listValue) == 0 {
		return
	}
	data, _ := gjson.Encode(listValue)
	strData := gconv.String(data)
	_, _ = redis.GetInstance().Set(ActiveCacheKey, strData)
}

func (c *equipmentCategoryCache) GetList() []*model.EquipmentCategory {
	cacheData, _ := redis.GetInstance().GetString(ActiveCacheKey)
	if len(cacheData) == 0 {
		syncLock = new(sync.Mutex)
		syncLock.Lock()
		defer syncLock.Unlock()
		// 从数据库中取
		list, _ := dao.EquipmentCategory.Where(dao.EquipmentCategory.Columns.Status + "=1").Order(dao.EquipmentCategory.Columns.ListOrder + " ASC,id DESC").All()
		if list == nil {
			return nil
		}
		c.SetList(list)
		return list
	}
	var result []*model.EquipmentCategory
	_ = gjson.DecodeTo(cacheData, &result)
	return result
}

func (c *equipmentCategoryCache) SetItem(id uint64, value interface{}) {
	key := fmt.Sprintf(ActiveItemCacheKey, id)
	dataMap := gconv.MapStrStr(value)
	_ = redis.GetInstance().HSetMap(key, dataMap)
}

func (c *equipmentCategoryCache) GetItem(id uint64) *model.EquipmentCategory {
	key := fmt.Sprintf(ActiveItemCacheKey, id)
	cacheData, _ := redis.GetInstance().HGetMap(key)
	if cacheData == nil || len(cacheData) == 0 {
		row, _ := dao.EquipmentCategory.Fields(dao.EquipmentCategory.Columns).FieldsEx("created_at", "updated_at",
			"list_order", "status").Where("id=? AND status=1", id).One()
		if row == nil {
			return nil
		}
		c.SetItem(id, row)
		return row
	}
	item := &model.EquipmentCategory{}
	_ = gconv.Struct(cacheData, &item)
	return item
}
func (c *equipmentCategoryCache) DeleteById(cid uint64) {
	key := fmt.Sprintf(ActiveItemCacheKey, cid)
	_, _ = redis.GetInstance().Del(key)
}
func (c *equipmentCategoryCache) Delete() {
	list := c.GetList()
	if len(list) > 0 {
		for _, row := range list {
			c.DeleteById(row.Id)
		}
	}
	for _, row := range list {
		c.DeleteById(row.Id)
	}
	_, _ = redis.GetInstance().Del(ActiveCacheKey)
}

func (c *equipmentCategoryCache) QueryCacheDelete() {
	_, _ = g.DB().GetCache().UpdateExpire(common.BackendEquipmentCategoryAll, 0)
}
