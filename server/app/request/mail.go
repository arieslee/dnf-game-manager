package request

type MailRequest struct {
	Mid    int `json:"mid" v:"required|integer#角色编号不能为空|角色编号只能是整数"`
	ItemId int `json:"item_id" v:"required|integer#物品编号不能为空|物品编号只能是整数"`
	Num    int `json:"num"  v:"required|integer#物品数量不能为空|物品数量只能是整数"`
	Coin   int `json:"coin" v:"integer#金币数量只能是整数"`
	Strong int `json:"strong" v:"integer#强化数量只能是整数"`
}
