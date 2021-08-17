package cache

import (
	"github.com/gogf/gf/text/gstr"
)

type managerCache struct {
}

var Manager = new(managerCache)

func (c *managerCache) DeleteAll(userMobile ...string) {
	EquipmentBrand.Delete()
	EquipmentCategory.Delete()
	Department.Delete()
	Config.DeleteAll(TypeEquipment)
	AdminRole.DeleteAll()
	var mobile string
	if len(userMobile) > 0 {
		mobile = userMobile[0]
	} else {
		mobile = ""
	}
	if len(mobile) > 0 {
		for _, mob := range userMobile {
			if gstr.Contains(mob, ",") {
				mobiles := gstr.Split(mob, ",")
				for _, m := range mobiles {
					User.DeleteByMobile(gstr.Trim(m))
				}
			} else {
				User.DeleteByMobile(gstr.Trim(mob))
			}
		}
	} else {
		User.DeleteAll()
	}
}
