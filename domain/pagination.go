package domain
import "math"

type Pagination interface {
	Valid() bool
	Page() int
	Pages() int
	Limit() int
	Total() int
	Offset() int
	SetTotal(total int)
	SetPages()
	ToResponse() *paginationResponse
}

type pagination struct {
	page  int
	pages int
	limit int
	total int
}

type paginationResponse struct {
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func NewPagination(page, limit int) Pagination {
	if page > 0 && limit == 0 {
		page = 1
	}
	pg := &pagination{
		page:  page,
		limit: limit,
	}
	return pg
}

func (pg pagination) Valid() bool {
	return pg.Page() > 0 && pg.Limit() >= 0
}

func (pg pagination) Page() int {
	return pg.page
}

func (pg pagination) Pages() int {
	return pg.pages
}

func (pg pagination) Limit() int {
	return pg.limit
}

func (pg pagination) Total() int {
	return pg.total
}

func (pg pagination) Offset() int {
	return (pg.Page() - 1) * pg.Limit()
}

func (pg *pagination) SetTotal(total int) {
	pg.total = total
	pg.SetPages()
}

func (pg *pagination) SetPages() {
	pg.pages = 1
	if pg.Limit() > 0 {
		pg.pages = int(math.Ceil(float64(pg.Total()) / float64(pg.Limit())))
	}
}

func (pg pagination) ToResponse() *paginationResponse {
	if pg.Valid() {
		return &paginationResponse{
			Page:  pg.Page(),
			Pages: pg.Pages(),
			Limit: pg.Limit(),
			Total: pg.Total(),
		}
	}
	return nil
}
