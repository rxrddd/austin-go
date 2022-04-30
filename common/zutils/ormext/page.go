package ormext

import (
	"gorm.io/gorm"
)

func Paginate(current, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := current
		if current < 1 {
			page = 1
		}
		if pageSize < 1 {
			pageSize = 20
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

type GetPageParams struct {
	Query      *gorm.DB
	CountQuery *gorm.DB
	OrderBY    string
	GroupBY    string
	Current    int
	PageSize   int
}

func (p GetPageParams) GetQuery() *gorm.DB {
	return p.Query
}
func GetPage(items interface{}, total *int64, params GetPageParams) error {
	countQuery := params.GetQuery()
	if params.CountQuery != nil {
		countQuery = params.CountQuery
	}
	err := countQuery.
		Count(total).Error
	if err != nil {
		return err
	}
	if *total > 0 {
		q := params.GetQuery().
			Scopes(Paginate(params.Current, params.PageSize))
		if len(params.OrderBY) > 0 {
			q = q.Order(params.OrderBY)
		}
		err = q.Find(items).Error
	}
	return err
}
