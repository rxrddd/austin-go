package model

type SmsRecord struct {
	ID                int64  `gorm:"column:id" json:"id"`                                   //
	MessageTemplateID int64  `gorm:"column:message_template_id" json:"message_template_id"` // 消息模板ID
	Phone             string `gorm:"column:phone" json:"phone"`                             // 手机号
	SupplierID        int    `gorm:"column:supplier_id" json:"supplier_id"`                 // 发送短信渠道商的ID
	SupplierName      string `gorm:"column:supplier_name" json:"supplier_name"`             // 发送短信渠道商的名称
	MsgContent        string `gorm:"column:msg_content" json:"msg_content"`                 // 短信发送的内容
	SeriesID          string `gorm:"column:series_id" json:"series_id"`                     // 下发批次的ID
	ChargingNum       int    `gorm:"column:charging_num" json:"charging_num"`               // 计费条数
	ReportContent     string `gorm:"column:report_content" json:"report_content"`           // 回执内容
	Status            int    `gorm:"column:status" json:"status"`                           // 短信状态： 10.发送 20.成功 30.失败
	SendDate          int64  `gorm:"column:send_date" json:"send_date"`                     // 发送日期：20211112
	Created           int64  `gorm:"column:created" json:"created"`                         // 创建时间
	Updated           int64  `gorm:"column:updated" json:"updated"`                         // 更新时间
}

func (m SmsRecord) TableName() string {
	return "sms_record"
}
