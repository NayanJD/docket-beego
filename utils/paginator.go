package utils

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/utils/pagination"
	beego "github.com/beego/beego/v2/server/web"
)

var PAGINATION_DATA_KEY string = "pagination"

var defaultPageSize int = 5

var defaultSortColumn string = "created_at"

var defaultSortOrder string = "ASC"

var defaultPageNumber int = 1

type PaginationQuery struct {
	Page       int    `json:"page" form:"page"`
	PageSize   int    `json:"page_size" form:"page_size"`
	SortColumn string `json:"sort_attr" form:"sort_attr"`
	SortOrder  string `json:"order" form:"order"`
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

func GetPagninatedQs(ctrl beego.Controller, qs orm.QuerySeter) (orm.QuerySeter, error) {

	pgF := NewPaginationQuery()

	req := ctrl.Ctx.Request

	err := req.ParseForm()

	if err != nil {
		return nil, err
	}

	if err := beego.ParseForm(req.Form, &pgF); err != nil {
		logs.Error(err)
		return nil, err
	}

	count, _ := qs.Count()

	paginator := &Paginator{Paginator: pagination.NewPaginator(req, pgF.PageSize, count), page: pgF.Page}

	ctrl.Ctx.Input.SetData(PAGINATION_DATA_KEY, pgF)

	return qs.OrderBy(pgF.SortColumn).Limit(pgF.PageSize).Offset(paginator.Offset()), nil
}
