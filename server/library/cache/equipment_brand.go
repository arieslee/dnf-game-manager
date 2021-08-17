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

var syncBrandLock *sync.Mutex
var EquipmentBrand = new(equipmentBrandCache)

type equipmentBrandCache struct{}

const (
	ActiveBrandCacheKey     = "equipment:brand:active"
	ActiveBrandItemCacheKey = "eqi:brand:active:item:%d"
)

func (c *equipmentBrandCache) SetList(listValue []*model.EquipmentBrand) {
	if listValue == nil || len(listValue) == 0 {
		return
	}
	data, _ := gjson.Encode(listValue)
	strData := gconv.String(data)
	_, _ = redis.GetInstance().Set(ActiveBrandCacheKey, strData)
}

func (c *equipmentBrandCache) GetList() []*model.EquipmentBrand {
	cacheData, _ := redis.GetInstance().GetString(ActiveBrandCacheKey)
	if len(cacheData) == 0 {
		syncBrandLock = new(sync.Mutex)
		syncBrandLock.Lock()
		defer syncBrandLock.Unlock()
		// 从数据库中取
		list, _ := dao.EquipmentBrand.Where(dao.EquipmentBrand.Columns.Status + "=1").Order(dao.EquipmentBrand.Columns.ListOrder + " ASC,id DESC").All()
		if list == nil {
			return nil
		}
		c.SetList(list)
		return list
	}
	var result []*model.EquipmentBrand
	_ = gjson.DecodeTo(cacheData, &result)
	return result
}

func (c *equipmentBrandCache) SetItem(id uint64, value interface{}) {
	key := fmt.Sprintf(ActiveBrandItemCacheKey, id)
	dataMap := gconv.MapStrStr(value)
	_ = redis.GetInstance().HSetMap(key, dataMap)
}

func (c *equipmentBrandCache) GetItem(id uint64) *model.EquipmentBrand {
	key := fmt.Sprintf(ActiveBrandItemCacheKey, id)
	cacheData, _ := redis.GetInstance().HGetMap(key)
	if cacheData == nil || len(cacheData) == 0 {
		row, _ := dao.EquipmentBrand.Fields(dao.EquipmentBrand.Columns).Fields("created_at", "updated_at", "list_order", "status").Where("id=? AND status=1", id).One()
		if row == nil {
			return nil
		}
		c.SetItem(id, row)
		return row
	}
	item := &model.EquipmentBrand{}
	_ = gconv.Struct(cacheData, &item)
	return item
}
func (c *equipmentBrandCache) DeleteById(id uint64) {
	key := fmt.Sprintf(ActiveBrandItemCacheKey, id)
	_, _ = redis.GetInstance().Del(key)
}
func (c *equipmentBrandCache) Delete() {
	list := c.GetList()
	if len(list) > 0 {
		for _, row := range list {
			c.DeleteById(row.Id)
		}
	}
	_, _ = redis.GetInstance().Del(ActiveBrandCacheKey)
}

func (c *equipmentBrandCache) QueryCacheDelete() {
	_, _ = g.DB().GetCache().UpdateExpire(common.BackendEquipmentBrandAll, 0)
}
