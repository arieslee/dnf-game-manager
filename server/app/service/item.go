package service

import (
	"dnf-game-manager/app/common"
	"dnf-game-manager/app/model"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	_ "github.com/mattn/go-sqlite3"
)

var Item = &itemService{
	Base: &common.ServiceBase{},
}

type itemService struct {
	Base *common.ServiceBase
}

func (s *itemService) List(r *ghttp.Request) (*model.ItemListResult, error) {
	kw := r.GetString("kw")
	where := "1=1"
	params := g.Slice{}
	if kw != "" {
		where += " AND item_name LIKE ?"
		params = append(params, "%"+kw+"%")
	}
	query := g.DB("item").Model("item_pool").Where(where, params)
	count, _ := query.Count()
	if count == 0 {
		return nil, errors.New("暂无内容")
	}
	result := new(model.ItemListResult)
	dataList := make([]*model.ItemDetail, 0)
	page, pageSize, _ := s.Base.PageComputedStructs(r, count, result)
	list, _ := query.Fields("id,item_name").Page(page, pageSize).All()
	_ = gconv.Structs(list, &dataList)
	result.DataList = dataList
	return result, nil
}
