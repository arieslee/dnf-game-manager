package cache

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"hess/app/dao"
	"hess/app/model"
	"hess/library/redis"
)

var Config *configCache

type configCache struct {
}

const (
	ConfigCacheKey = "config:%s:"
	TypeEquipment  = "equipment"
)

func (c *configCache) Set(cfgType string, listValue []*model.Config) {
	if len(listValue) <= 0 {
		return
	}
	data := g.MapStrStr{}
	for _, val := range listValue {
		data[val.CfgName] = val.CfgValue
	}
	if len(data) <= 0 {
		return
	}
	jsonData, _ := gjson.Encode(data)
	key := fmt.Sprintf(ConfigCacheKey, cfgType)
	redis.GetInstance().Set(key, gconv.String(jsonData))
}
func (c *configCache) Delete(cfgType string) {
	key := fmt.Sprintf(ConfigCacheKey, cfgType)
	_, _ = redis.GetInstance().Del(key)
}
func (c *configCache) DeleteAll(cfgType ...string) {
	if len(cfgType) == 0 {
		return
	}
	for _, ct := range cfgType {
		c.Delete(ct)
	}
}

func (c *configCache) Get(cfgType string) g.MapStrStr {
	cacheKey := fmt.Sprintf(ConfigCacheKey, cfgType)
	cacheData, err := redis.GetInstance().GetString(cacheKey)
	if err != nil {
		return nil
	}
	if len(cacheData) == 0 {
		list, err := dao.Config.Fields("cfg_name,cfg_value,cfg_type").Where("cfg_type=?", cfgType).All()
		if err != nil {
			return nil
		}
		if list == nil {
			return nil
		}
		c.Set(TypeEquipment, list)
		newData := g.MapStrStr{}
		for _, v := range list {
			newData[v.CfgName] = v.CfgValue
		}
		return newData
	} else {
		jsonStr, _ := gjson.Decode(cacheData)
		return gconv.MapStrStr(jsonStr)
	}
}
func (c *configCache) GetEquipmentValue(key string) string {
	return c.GetValue(TypeEquipment, key)
}
func (c *configCache) GetValue(cfgType, key string) string {
	data := c.Get(cfgType)
	if data == nil {
		return ""
	}
	mapData := gconv.MapStrStr(data)
	if value, ok := mapData[key]; ok {
		return value
	}
	return ""
}
