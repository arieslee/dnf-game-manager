package model

type BaseListResult struct {
	TotalCount int64 `json:"total_count"`
	TotalPage  int   `json:"total_page"`
	PageSize   int   `json:"page_size"`
	Page       int   `json:"current_page"`
}
type RoleListItem struct {
	Mid        int    `json:"m_id"`
	CharacNo   int    `json:"charac_no"`
	CharacName string `json:"charac_name"`
	Job        int    `json:"job"`
	Lev        int    `json:"lev"`
	GrowType   int    `json:"grow_type"`
	Cera       int    `json:"cera"`
}

// RoleListResult 列表结果
type RoleListResult struct {
	DataList []*RoleListItem `json:"data_list"` // 列表
	BaseListResult
}

type ItemDetail struct {
	Id       int    `json:"id"`
	ItemName string `json:"item_name"`
}

type ItemListResult struct {
	DataList []*ItemDetail `json:"data_list"` // 列表
	BaseListResult
}

// AccountListResult 列表结果
type AccountListResult struct {
	DataList []*Accounts `json:"data_list"` // 列表
	BaseListResult
}
