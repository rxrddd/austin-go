package repo

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/model/cls"
	"austin-go/common/dbx"
	"austin-go/common/idgen"
	"context"
	"gorm.io/gorm"
)

type SmsRecordRepo struct {
}

func NewSmsRecordRepo() *SmsRecordRepo {
	return &SmsRecordRepo{}
}

func (a *SmsRecordRepo) getModel(ctx context.Context) *gorm.DB {
	return dbx.GetDb(ctx).Model(&model.SmsRecord{})
}

//func (a *SmsRecordRepo) Page(ctx context.Context, req types.SmsRecordListReq) (items []model.SmsRecord, total int64, err error) {
//	builder := zsqlx.NewBuilder()
//	if len(req.Name) > 0 {
//		builder.Like(cls.ClsSmsRecord.Name, req.Name)
//	}
//	cond, args := builder.End()
//	err = paginate.GetPage(&items, &total, paginate.GetPageParams{
//		Current:  req.Current,
//		PageSize: req.PageSize,
//		Query:    a.getModel(ctx).Where(cond, args...),
//	})
//	return items, total, err
//}

func (a *SmsRecordRepo) One(ctx context.Context, id int64) (item *model.SmsRecord, err error) {
	err = a.getModel(ctx).Where(cls.ClsSmsRecord.ID, id).Take(&item).Error
	return item, err
}
func (a *SmsRecordRepo) OneByField(ctx context.Context, field string, value interface{}) (item *model.SmsRecord, err error) {
	err = a.getModel(ctx).Where(field, value).Take(&item).Error
	return item, err
}

func (a *SmsRecordRepo) All(ctx context.Context) (item []model.SmsRecord, err error) {
	err = a.getModel(ctx).Find(&item).Error
	return item, err
}

func (a *SmsRecordRepo) ListByField(ctx context.Context, field string, value interface{}) (item []model.SmsRecord, err error) {
	err = a.getModel(ctx).Where(field, value).Find(&item).Error
	return item, err
}

func (a *SmsRecordRepo) ListByMap(ctx context.Context, m map[string]interface{}) (item []model.SmsRecord, err error) {
	err = a.getModel(ctx).Where(m).Find(&item).Error
	return item, err
}

func (a *SmsRecordRepo) Create(ctx context.Context, m *model.SmsRecord) error {
	m.ID = idgen.NextID()
	return a.getModel(ctx).Create(m).Error
}
func (a *SmsRecordRepo) BatchCreate(ctx context.Context, m []model.SmsRecord) error {
	for i, record := range m {
		if record.ID == 0 {
			m[i].ID = idgen.NextID()
		}
	}
	return a.getModel(ctx).Create(m).Error
}

func (a *SmsRecordRepo) Edit(ctx context.Context, m *model.SmsRecord) error {
	return a.getModel(ctx).Where(cls.ClsSmsRecord.ID, m.ID).Updates(m).Error
}

func (a *SmsRecordRepo) DeleteByPrimaryKey(ctx context.Context, id int64) error {
	return a.getModel(ctx).Delete(cls.ClsSmsRecord.ID, id).Error
}

func (a *SmsRecordRepo) DeleteByField(ctx context.Context, field string, value interface{}) error {
	return a.getModel(ctx).Delete(field, value).Error
}

func (a *SmsRecordRepo) DeleteByMap(ctx context.Context, m map[string]interface{}) error {
	return a.getModel(ctx).Delete(m).Error
}
