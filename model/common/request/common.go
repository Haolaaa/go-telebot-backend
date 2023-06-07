package request

type PageInfo struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Keyword  string `json:"keyword" form:"keyword"`
}

type GetByID struct {
	ID int `json:"id" form:"id"`
}

func (r *GetByID) Uint() uint {
	return uint(r.ID)
}

type IDsReq struct {
	IDs []int `json:"ids" form:"ids"`
}

type GetAuthorityID struct {
	AuthorityID uint `json:"authorityID" form:"authorityID"`
}

type Empty struct{}
