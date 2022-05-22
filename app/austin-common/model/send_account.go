package model

type SendAccount struct {
	ID          int64  `gorm:"column:id" json:"id"`                     //
	SendChannel int    `gorm:"column:send_channel" json:"send_channel"` // 发送渠道
	Config      string `gorm:"column:config" json:"config"`             // 账户配置
	Title       string `gorm:"column:title" json:"title"`               // 账号名称
}

func (m SendAccount) TableName() string {
	return "send_account"
}
