package repo

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/model/cls"
	"austin-go/common/dbx"
	"austin-go/common/idgen"
	"context"
	"gorm.io/gorm"
)

type AccountRepo struct {
}

func NewAccountRepo() *AccountRepo {
	return &AccountRepo{}
}

func (a *AccountRepo) getModel(ctx context.Context) *gorm.DB {
	return dbx.GetDb(ctx).Model(&model.Account{})
}

func (a *AccountRepo) FindByUserName(ctx context.Context, username string) (item *model.Account, err error) {
	err = a.getModel(ctx).Where(cls.ClsAccount.Username, username).Take(&item).Error
	return item, err
}

func (a *AccountRepo) One(ctx context.Context, id int64) (item *model.Account, err error) {
	err = a.getModel(ctx).Where(cls.ClsAccount.ID, id).Take(&item).Error
	return item, err
}
func (a *AccountRepo) OneByField(ctx context.Context, field string, value interface{}) (item *model.Account, err error) {
	err = a.getModel(ctx).Where(field, value).Take(&item).Error
	return item, err
}

func (a *AccountRepo) All(ctx context.Context) (item []model.Account, err error) {
	err = a.getModel(ctx).Find(&item).Error
	return item, err
}

func (a *AccountRepo) ListByField(ctx context.Context, field string, value interface{}) (item []model.Account, err error) {
	err = a.getModel(ctx).Where(field, value).Find(&item).Error
	return item, err
}

func (a *AccountRepo) ListByMap(ctx context.Context, m map[string]interface{}) (item []model.Account, err error) {
	err = a.getModel(ctx).Where(m).Find(&item).Error
	return item, err
}

func (a *AccountRepo) Create(ctx context.Context, m *model.Account) error {
	m.ID = idgen.NextID()
	return a.getModel(ctx).Create(m).Error
}

func (a *AccountRepo) Edit(ctx context.Context, m *model.Account) error {
	return a.getModel(ctx).Where(cls.ClsAccount.ID, m.ID).Updates(m).Error
}
