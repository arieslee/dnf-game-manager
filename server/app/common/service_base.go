// Package common
package common

import (
	"database/sql"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"math"
	"strings"
)

type ServiceBase struct{}

func (s *ServiceBase) PageComputed(r *ghttp.Request, totalCount int) (int, int, int) {
	page := r.GetInt("page")
	if page == 0 {
		page = 1
	}
	pageSize := r.GetInt("pageSize")
	if pageSize == 0 {
		pageSize = g.Cfg().GetInt("app.pageSize")
	}
	//go里面如果两个int相除，math.Ceil不会向上取整的,所以，必须转成float现做除法
	divPage := gconv.Float64(totalCount) / gconv.Float64(pageSize)
	totalPage := math.Ceil(divPage)
	return page, pageSize, int(totalPage)
}

// PageComputedStructs 计算分页并绑定到result
func (s *ServiceBase) PageComputedStructs(r *ghttp.Request, totalCount int, bindPointer interface{}) (int, int, int) {
	page, pageSize, totalPage := s.PageComputed(r, totalCount)
	binder := g.Map{
		"page_size":    pageSize,
		"total_page":   totalPage,
		"total_count":  totalCount,
		"current_page": page,
	}
	err := gconv.Struct(binder, bindPointer)
	if err != nil {
		g.Log().Line(true).Println(err.Error())
		return 0, 0, 0
	}
	return page, pageSize, totalPage
}

func (s *ServiceBase) GetStartEndAt(r *ghttp.Request) (uint, uint, error) {
	request := r.GetMap()
	startAt := gconv.String(request["start_at"])
	endAt := gconv.String(request["end_at"])
	var (
		startTime uint
		endTime   uint
	)
	if len(startAt) == 0 {
		startTime = 0
	} else {
		startAtTime, _ := gtime.StrToTime(gconv.String(request["start_at"]))
		startTime = uint(startAtTime.Unix())
	}
	if len(endAt) == 0 {
		endTime = 0
	} else {
		endAtTime, _ := gtime.StrToTime(gconv.String(request["end_at"]))
		endTime = uint(endAtTime.Unix())
	}
	if startTime > 0 && endTime > 0 && endTime <= startTime {
		return startTime, endTime, errors.New("有效期的结束时间必须大于开始时间")
	}
	return startTime, endTime, nil
}
func (s *ServiceBase) GetBatchIds(r *ghttp.Request) (string, error) {
	ids := r.Get("ids")
	idsArray := gconv.SliceStr(ids)
	if len(idsArray) < 1 {
		return "", errors.New("请选择要操作的内容")
	}
	in := ""
	for _, v := range idsArray {
		in += v + ","
	}
	in = strings.TrimSuffix(in, ",")
	return in, nil
}
func (s *ServiceBase) UpdateSuccess(res sql.Result) error {
	affNum, _ := res.RowsAffected()
	if affNum <= 0 {
		return errors.New("修改失败")
	}
	return nil
}
func (s *ServiceBase) DeleteSuccess(res sql.Result) error {
	affNum, _ := res.RowsAffected()
	if affNum <= 0 {
		return errors.New("删除失败")
	}
	return nil
}
