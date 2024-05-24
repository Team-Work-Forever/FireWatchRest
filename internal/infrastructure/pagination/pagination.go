package pagination

import (
	"math"
	"strconv"
)

type Pagination struct {
	PageSize   uint64 `json:"page_size,omitempty"`
	Page       uint64 `json:"page,omitempty"`
	TotalPages uint64 `json:"total_pages,omitempty"`
}

func New(pageString, pageSizeString string) (*Pagination, error) {
	page, err := strconv.ParseUint(pageString, 10, 64)

	if err != nil {
		return nil, err
	}

	pageSize, err := strconv.ParseUint(pageSizeString, 10, 64)

	if err != nil {
		return nil, err
	}

	return &Pagination{
		Page:     page,
		PageSize: pageSize,
	}, nil
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
