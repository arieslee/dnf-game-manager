package request

type RoleRequest struct {
	CharacName string `json:"charac_name" v:"required#角色名称不能为空"`
	Job        int    `json:"job" v:"required|integer#职业不能为空|职业只能是整数"`
	Lev        int    `json:"lev"  v:"integer#等级不能为空|等级只能是整数"`
	Coin       int    `json:"coin" v:"integer#金币数量只能是整数"`
	Cera       int    `json:"cera" v:"integer#点卷数量只能是整数"`
	Mid        int    `json:"m_id"  v:"required|integer#账号UID不能为空|账号UID只能是整数"`
}
