package service

import (
	"dnf-game-manager/app/common"
	"dnf-game-manager/app/model"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Account = &accountService{
	Base: &common.ServiceBase{},
}

type accountService struct {
	Base *common.ServiceBase
}

func (s *accountService) List(r *ghttp.Request) (*model.AccountListResult, error) {
	where := "1=1"
	params := g.Slice{}
	kw := r.GetString("kw")
	if kw != "" {
		where += " AND INSTR(accountname,?)>0"
		params = append(params, kw)
	}
	query := g.DB().Schema("d_taiwan").
		Model("accounts").Where(where, params)
	cloneQuery := query.Clone()
	count, _ := cloneQuery.Count()
	if count == 0 {
		return nil, errors.New("暂无数据")
	}
	res := &model.AccountListResult{}
	dataList := make([]*model.Accounts, 0)
	page, pageSize, _ := s.Base.PageComputedStructs(r, count, res)
	list, err := query.
		Fields("UID,accountname,qq,dzuid,billing,VIP").
		Order("UID ASC").
		Page(page, pageSize).
		All()
	if err != nil {
		return nil, err
	}
	err = gconv.Structs(list, &dataList)
	if err != nil {
		g.Log().Line(true).Println(err.Error())
	}
	res.DataList = dataList
	return res, nil
}

func (s *accountService) DeleteAccount(r *ghttp.Request) error {
	id := r.GetInt("id")
	if id == 0 {
		return errors.New("要删除的id不能为空")
	}
	query := g.DB().Schema("d_taiwan").
		Model("accounts").Where("id=?", id)
	cloneQuery := query.Clone()
	row, _ := query.One()
	if row == nil {
		return errors.New("要删除的账号不存在")
	}
	_, err := cloneQuery.Delete()
	if err != nil {
		g.Log().Line(true).Debug("账号删除失败,err=" + err.Error())
		return errors.New("账号删除失败")
	}
	// 连同角色一起删除
	roleDelete := r.GetInt("role_delete", -1)

	if roleDelete == 1 {
		where := "1=1"
		params := g.Slice{}
		g.DB().Schema("taiwan_cain").
			Model("charac_info").Where(where, params).Delete()
	}
	return nil
}

func (s *accountService) RoleList(r *ghttp.Request) (*model.RoleListResult, error) {
	uid := r.GetUint64("uid")
	where := "delete_flag!=1"
	params := g.Slice{}
	if uid > 0 {
		where += " AND m_id=?"
		params = append(params, uid)
	}
	query := g.DB().Schema("taiwan_cain").
		Model("charac_info").Where(where, params)
	cloneQuery := query.Clone()
	count, _ := cloneQuery.Count()
	if count == 0 {
		return nil, errors.New("暂无数据")
	}
	res := &model.RoleListResult{}
	dataList := make([]*model.RoleListItem, 0)
	page, pageSize, _ := s.Base.PageComputedStructs(r, count, res)
	list, err := query.
		Fields("m_id,charac_no,charac_name,job,lev,grow_type").
		Order("charac_no ASC").
		Page(page, pageSize).
		All()
	if err != nil {
		return nil, err
	}
	result := g.Slice{}
	if len(list) > 0 {
		for _, row := range list {
			item := gconv.Map(row)
			cera, _ := g.DB().Schema("taiwan_billing").Model("cash_cera").Fields("cera").Where("account=?", row["m_id"]).One()
			if cera == nil {
				item["cera"] = 0
			} else {
				item["cera"] = cera["cera"]
			}
			result = append(result, item)
		}
		err := gconv.Structs(result, &dataList)
		if err != nil {
			g.Log().Line(true).Println(err.Error())
		}
	}
	res.DataList = dataList
	return res, nil
}
