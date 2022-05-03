package repo

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/model/cls"
	"austin-go/common/dbx"
	"austin-go/common/idgen"
	"context"
	"gorm.io/gorm"
)

type SendAccountRepo struct {
}

func NewSendAccountRepo() *SendAccountRepo {
	return &SendAccountRepo{}
}

func (a *SendAccountRepo) getModel(ctx context.Context) *gorm.DB {
	return dbx.GetDb(ctx).Model(&model.SendAccount{})
}

func (a *SendAccountRepo) One(ctx context.Context, id int64) (item *model.SendAccount, err error) {
	err = a.getModel(ctx).Where(cls.ClsSendAccount.ID, id).Limit(1).Find(&item).Error
	return item, err
}
func (a *SendAccountRepo) OneByField(ctx context.Context, field string, value interface{}) (item *model.SendAccount, err error) {
	err = a.getModel(ctx).Where(field, value).Take(&item).Error
	return item, err
}

func (a *SendAccountRepo) All(ctx context.Context) (item []model.SendAccount, err error) {
	err = a.getModel(ctx).Find(&item).Error
	return item, err
}

func (a *SendAccountRepo) ListByField(ctx context.Context, field string, value interface{}) (item []model.SendAccount, err error) {
	err = a.getModel(ctx).Where(field, value).Find(&item).Error
	return item, err
}

func (a *SendAccountRepo) ListByMap(ctx context.Context, m map[string]interface{}) (item []model.SendAccount, err error) {
	err = a.getModel(ctx).Where(m).Find(&item).Error
	return item, err
}

func (a *SendAccountRepo) Create(ctx context.Context, m *model.SendAccount) error {
	m.ID = idgen.NextID()
	return a.getModel(ctx).Create(m).Error
}

func (a *SendAccountRepo) Edit(ctx context.Context, m *model.SendAccount) error {
	return a.getModel(ctx).Where(cls.ClsSendAccount.ID, m.ID).Updates(m).Error
}

func (a *SendAccountRepo) DeleteByPrimaryKey(ctx context.Context, id int64) error {
	return a.getModel(ctx).Delete(cls.ClsSendAccount.ID, id).Error
}

func (a *SendAccountRepo) DeleteByField(ctx context.Context, field string, value interface{}) error {
	return a.getModel(ctx).Delete(field, value).Error
}

func (a *SendAccountRepo) DeleteByMap(ctx context.Context, m map[string]interface{}) error {
	return a.getModel(ctx).Delete(m).Error
}
