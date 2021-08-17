package cache

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"hess/app/dao"
	"hess/app/model"
	"hess/library/redis"
)

var User = new(userCache)

type userCache struct {
}

const (
	UserInfoCacheKey = "user:info:%d" // %d=userId
)

func (c *userCache) GetInfo(uid uint64) *model.UserDetail {
	if uid <= 0 {
		return nil
	}
	key := fmt.Sprintf(UserInfoCacheKey, uid)
	cacheData, _ := redis.GetInstance().HGetMap(key)
	//var userDetail *model.UserDetail
	userDetail := &model.UserDetail{}
	if cacheData == nil || len(cacheData) <= 0 {
		userInfo, _ := dao.User.Fields(dao.User.Columns).FieldsEx("updated_at,created_at,login_at,login_ip,login_times,user_agent").WherePri(uid).One()
		if userInfo == nil {
			return nil
		}
		if userInfo.Status != 1 {
			return nil
		}
		userInfoData, _ := dao.UserInfo.Fields(dao.UserInfo.Columns).Where("uid=?", uid).One()
		if userInfoData == nil {
			return nil
		}
		data := gconv.MapStrStr(userInfo)
		infoData := gconv.MapStrStr(userInfoData)
		userProfile := g.MapStrStr{}
		for k, v := range data {
			userProfile[k] = v
		}
		for k, v := range infoData {
			userProfile[k] = v
		}
		_ = redis.GetInstance().HSetMap(key, userProfile)
		err := gconv.Struct(userProfile, userDetail)
		if err != nil {
			return nil
		}
		return userDetail
	}
	err := gconv.Struct(cacheData, &userDetail)
	if err != nil {
		return nil
	}
	return userDetail
}

func (c *userCache) Delete(uid uint64) {
	key := fmt.Sprintf(UserInfoCacheKey, uid)
	_, _ = redis.GetInstance().Del(key)
}
func (c *userCache) DeleteByMobile(mobile string) {
	if mobile == "" {
		return
	}
	userData, _ := dao.User.Fields(dao.User.Columns.Id).Where(dao.User.Columns.Mobile+"=?", mobile).One()
	if userData == nil {
		return
	}
	c.Delete(userData.Id)
}
func (c *userCache) DeleteAll() {
	all, _ := dao.User.Fields(dao.User.Columns.Id).All()
	if len(all) > 0 {
		for _, v := range all {
			c.Delete(v.Id)
		}
	}
}
