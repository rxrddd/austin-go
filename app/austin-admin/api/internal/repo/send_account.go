package repo

import (
	"austin-go/app/austin-admin/api/internal/types"
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/model/cls"
	"austin-go/common/dbx"
	"austin-go/common/zutils/zsqlx"
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

func (a *SendAccountRepo) FindByUserName(ctx context.Context, username string) (item *model.SendAccount, err error) {
	err = a.getModel(ctx).Where(cls.ClsAccount.Username, username).Take(&item).Error
	return item, err
}

func (a *SendAccountRepo) One(ctx context.Context, id int64) (item *model.SendAccount, err error) {
	err = a.getModel(ctx).Where(cls.ClsAccount.ID, id).Take(&item).Error
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
func (a *SendAccountRepo) FindAll(ctx context.Context, req types.SendAccountListReq) (item []model.SendAccount, err error) {
	builder := zsqlx.NewBuilder()
	if len(req.Title) > 0 {
		builder.RLike(cls.ClsSendAccount.Title, req.Title)
	}
	if len(req.SendChannel) > 0 {
		builder.Eq(cls.ClsSendAccount.SendChannel, req.SendChannel)
	}
	cond, args := builder.End()
	err = a.getModel(ctx).Where(cond, args...).Find(&item).Error
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
	return a.getModel(ctx).Create(m).Error
}

func (a *SendAccountRepo) Edit(ctx context.Context, m *model.SendAccount) error {
	return a.getModel(ctx).Where(cls.ClsAccount.ID, m.ID).Updates(m).Error
}
