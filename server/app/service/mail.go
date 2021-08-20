package service

import (
	"dnf-game-manager/app/common"
	"dnf-game-manager/app/request"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

var Mail = &mailService{
	Base: &common.ServiceBase{},
}

type mailService struct {
	Base *common.ServiceBase
}

func (s *mailService) Send(r *ghttp.Request) error {
	req := &request.MailRequest{}
	if err := r.Parse(req); err != nil {
		return err
	}
	schema := "taiwan_cain_2nd"
	table := "letter"
	dateTime := gtime.Now().Format("Y-m-d H:i:s")
	data := g.Map{
		"charac_no":        req.CharacNo,
		"send_charac_name": "Game Master",
		"letter_text":      req.LetterText,
		"reg_date":         dateTime,
	}
	db := g.DB().Schema(schema)
	res, err := db.Model(table).Data(data).Insert()
	if err != nil {
		return err
	}
	lastId, _ := res.LastInsertId()
	postalData := g.Map{
		"occ_time":          dateTime,
		"send_charac_name":  data["send_charac_name"],
		"receive_charac_no": req.CharacNo,
		"item_id":           req.ItemId,
		"add_info":          req.Num,
		"upgrade":           req.Strong,
		"amplify_option":    0,
		"amplify_value":     0,
		"gold":              req.Coin,
		"seal_flag":         0,
		"letter_id":         lastId,
		"seperate_upgrade":  0,
	}
	postalTable := "postal"
	_, err = db.Model(postalTable).Data(postalData).Insert()
	if err != nil {
		return err
	}
	return nil
}
