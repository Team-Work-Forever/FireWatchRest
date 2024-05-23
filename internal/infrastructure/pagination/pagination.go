package pagination

import "math"

type Pagination struct {
	PageSize   uint64 `json:"page_size,omitempty"`
	Page       uint64 `json:"page,omitempty"`
	TotalPages uint64 `json:"total_pages,omitempty"`
}

func (p *Pagination) GetOffset() int {
	return (int(p.GetPage()) - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	return int(p.PageSize)
}

func (p *Pagination) GetPage() uint64 {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) SetTotalPages(total int) {
	p.TotalPages = uint64(math.Ceil(float64(total) / float64(p.GetLimit())))
}
