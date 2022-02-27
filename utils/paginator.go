package utils

import (
	"net/http"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/utils/pagination"
)

var defaultPageSize int = 5

var defaultSortColumn string = "created_at"

var defaultSortOrder string = "ASC"

var defaultPageNumber int = 1

type PaginationQuery struct {
	Page       int    `form:"page"`
	PageSize   int    `form:"page_size"`
	SortColumn string `form:"sort_attr"`
	SortOrder  string `form:"order"`
}

type Paginator struct {
	*pagination.Paginator
	page int
}

func NewPaginationQuery() PaginationQuery {
	return PaginationQuery{
		Page:       defaultPageNumber,
		PageSize:   defaultPageSize,
		SortColumn: defaultSortColumn,
		SortOrder:  defaultSortOrder,
	}
}

func (p *Paginator) Page() int {
	if p.page > p.PageNums() {
		p.page = p.PageNums()
	}
	if p.page <= 0 {
		p.page = 1
	}
	return p.page
}

func GetPagninatedQs(pq PaginationQuery, req *http.Request, qs orm.QuerySeter) (orm.QuerySeter, *Paginator, error) {
	count, err := qs.Count()

	if err != nil {
		return nil, nil, err
	}

	paginator := &Paginator{Paginator: pagination.NewPaginator(req, pq.PageSize, count), page: pq.Page}

	return qs.OrderBy(pq.SortColumn).Limit(pq.PageSize).Offset(paginator.Offset()), paginator, nil
}
