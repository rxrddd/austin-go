package ormext

import "gorm.io/gorm"

func NormalSearch(db *gorm.DB) *gorm.DB {
	return db.Where("is_delete", 0)
}
