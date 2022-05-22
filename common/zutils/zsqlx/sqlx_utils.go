package zsqlx

import (
	"austin-go/common/zutils/ormext"
	"gorm.io/gorm"
)

type GetPageParams struct {
	Current  int
	PageSize int
	Query    *gorm.DB
}

func (p GetPageParams) GetQuery() *gorm.DB {
	return p.Query
}
func GetPage(items interface{}, total *int64, params GetPageParams) error {
	err := params.GetQuery().
		Count(total).Error
	if err != nil {
		return err
	}
	if *total > 0 {
		err = params.GetQuery().
			Scopes(ormext.Paginate(params.Current, params.PageSize)).
			Find(&items).Error
	}
	return err
}
