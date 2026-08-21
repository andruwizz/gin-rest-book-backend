package model

type Pagination struct {
	Records      int64 `json:"records"`
	TotalRecords int64 `json:"total_records"`
	Limit        int   `json:"limit"`
	Page         int   `json:"page"`
	TotalPage    int   `json:"total_page"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPageNumber() - 1) * p.GetPageLimit()
}

func (p *Pagination) GetPageLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPageNumber() int {
	if p.Page <= 1 {
		p.Page = 1
	}
	return p.Page
}
