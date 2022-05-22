package model

import (
	"austin-go/common/zutils/encrypt"
	"austin-go/common/zutils/randomx"
	"gorm.io/gorm"
)

type Account struct {
	ID        int64  `gorm:"column:id" json:"id"`                 // 主键
	CreatedBy string `gorm:"column:created_by" json:"created_by"` // 创建人
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedBy string `gorm:"column:updated_by" json:"updated_by"` // 更新人
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"` // 更新时间
	Username  string `gorm:"column:username" json:"username"`     // 账户
	Password  string `gorm:"column:password" json:"password"`     // 密码
	Salt      string `gorm:"column:salt" json:"salt"`             // 密码盐
	RoleID    string `gorm:"column:role_id" json:"role_id"`       // 角色ID,多角色都好隔开
	Nickname  string `gorm:"column:nickname" json:"nickname"`     // 昵称
	IsDelete  bool   `gorm:"column:is_delete" json:"is_delete"`   // 是否删除
}

func (m *Account) BeforeCreate(*gorm.DB) error {
	cur := randomx.RandStr(10)
	m.Salt = randomx.RandStr(10)
	m.Password = encrypt.MD5(encrypt.MD5(m.Password) + cur)
	return nil
}

func (m Account) CheckPassword(password string) bool {
	return m.Password == encrypt.MD5(encrypt.MD5(password)+m.Salt)
}

func (m Account) TableName() string {
	return "account"
}
