package paging

type Paging struct {
	Offset int   `json:"offset" form:"page"`
	Limit  int   `json:"limit" form:"limit"`
	Total  int64 `json:"total" form:"-"`
}

func (p *Paging) Fulfill() {
	if p.Offset <= 0 {
		p.Offset = 1
	}

	if p.Limit <= 0 {
		p.Limit = 50
	}
}
