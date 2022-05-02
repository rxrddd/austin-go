package repo

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/model/cls"
	"austin-go/common/dbx"
	"austin-go/common/gormc"
	"austin-go/common/idgen"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

type MessageTemplateRepo struct {
	cache gormc.CachedConn
}

func NewMessageTemplateRepo(c cache.CacheConf, opts ...cache.Option) *MessageTemplateRepo {
	return &MessageTemplateRepo{
		cache: gormc.NewConn(c, opts...),
	}
}

func (a *MessageTemplateRepo) getModel(ctx context.Context) *gorm.DB {
	return dbx.GetDb(ctx).Model(&model.MessageTemplate{})
}

func (a *MessageTemplateRepo) One(ctx context.Context, id int64) (item model.MessageTemplate, err error) {
	key := fmt.Sprintf("messagetemplate_%d", id)
	err = a.cache.QueryRowCtx(ctx, &item, key, func(ctx context.Context, v interface{}) error {
		return a.getModel(ctx).Where(cls.ClsMessageTemplate.ID, id).Limit(1).Find(&item).Error
	})
	return
}
func (a *MessageTemplateRepo) OneByField(ctx context.Context, field string, value interface{}) (item *model.MessageTemplate, err error) {
	err = a.getModel(ctx).Where(field, value).Take(&item).Error
	return item, err
}

func (a *MessageTemplateRepo) All(ctx context.Context) (item []model.MessageTemplate, err error) {
	err = a.getModel(ctx).Find(&item).Error
	return item, err
}

func (a *MessageTemplateRepo) ListByField(ctx context.Context, field string, value interface{}) (item []model.MessageTemplate, err error) {
	err = a.getModel(ctx).Where(field, value).Find(&item).Error
	return item, err
}

func (a *MessageTemplateRepo) ListByMap(ctx context.Context, m map[string]interface{}) (item []model.MessageTemplate, err error) {
	err = a.getModel(ctx).Where(m).Find(&item).Error
	return item, err
}

func (a *MessageTemplateRepo) Create(ctx context.Context, m *model.MessageTemplate) error {
	m.ID = idgen.NextID()
	return a.getModel(ctx).Create(m).Error
}

func (a *MessageTemplateRepo) Edit(ctx context.Context, m *model.MessageTemplate) error {
	return a.getModel(ctx).Where(cls.ClsMessageTemplate.ID, m.ID).Updates(m).Error
}

func (a *MessageTemplateRepo) DeleteByPrimaryKey(ctx context.Context, id int64) error {
	return a.getModel(ctx).Delete(cls.ClsMessageTemplate.ID, id).Error
}

func (a *MessageTemplateRepo) DeleteByField(ctx context.Context, field string, value interface{}) error {
	return a.getModel(ctx).Delete(field, value).Error
}

func (a *MessageTemplateRepo) DeleteByMap(ctx context.Context, m map[string]interface{}) error {
	return a.getModel(ctx).Delete(m).Error
}
