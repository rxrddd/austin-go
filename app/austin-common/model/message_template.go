package model

type MessageTemplate struct {
	ID                  int64  `gorm:"column:id" json:"id"`                                     //
	Name                string `gorm:"column:name" json:"name"`                                 // 标题
	AuditStatus         int    `gorm:"column:audit_status" json:"audit_status"`                 // 当前消息审核状态： 10.待审核 20.审核成功 30.被拒绝
	FlowID              string `gorm:"column:flow_id" json:"flow_id"`                           // 工单ID
	MsgStatus           int8   `gorm:"column:msg_status" json:"msg_status"`                     // 当前消息状态：10.新建 20.停用 30.启用 40.等待发送 50.发送中 60.发送成功 70.发送失败
	CronTaskID          int64  `gorm:"column:cron_task_id" json:"cron_task_id"`                 // 定时任务Id (xxl-job-admin返回)
	CronCrowdPath       string `gorm:"column:cron_crowd_path" json:"cron_crowd_path"`           // 定时发送人群的文件路径
	ExpectPushTime      string `gorm:"column:expect_push_time" json:"expect_push_time"`         // 期望发送时间：0:立即发送 定时任务以及周期任务:cron表达式
	IDType              int    `gorm:"column:id_type" json:"id_type"`                           // 消息的发送ID类型：10. userId 20.did 30.手机号 40.openId 50.email 60.企业微信userId
	SendChannel         int    `gorm:"column:send_channel" json:"send_channel"`                 // 消息发送渠道：10.IM 20.Push 30.短信 40.Email 50.公众号 60.小程序 70.企业微信
	TemplateType        int    `gorm:"column:template_type" json:"template_type"`               // 10.运营类 20.技术类接口调用
	TemplateSn          string `gorm:"column:template_sn" json:"template_sn"`                   // 发送消息的模版ID
	MsgType             int    `gorm:"column:msg_type" json:"msg_type"`                         // 10.通知类消息 20.营销类消息 30.验证码类消息
	ShieldType          int    `gorm:"column:shield_type" json:"shield_type"`                   // 10.夜间不屏蔽 20.夜间屏蔽 30.夜间屏蔽(次日早上9点发送)
	MsgContent          string `gorm:"column:msg_content" json:"msg_content"`                   // 消息内容 占位符用{$var}表示
	SendAccount         int    `gorm:"column:send_account" json:"send_account"`                 // 发送账号 一个渠道下可存在多个账号
	Creator             string `gorm:"column:creator" json:"creator"`                           // 创建者
	Updator             string `gorm:"column:updator" json:"updator"`                           // 更新者
	Auditor             string `gorm:"column:auditor" json:"auditor"`                           // 审核人
	Team                string `gorm:"column:team" json:"team"`                                 // 业务方团队
	Proposer            string `gorm:"column:proposer" json:"proposer"`                         // 业务方
	IsDeleted           int    `gorm:"column:is_deleted" json:"is_deleted"`                     // 是否删除：0.不删除 1.删除
	Created             int32  `gorm:"column:created" json:"created"`                           // 创建时间
	Updated             int32  `gorm:"column:updated" json:"updated"`                           // 更新时间
	DeduplicationConfig string `gorm:"column:deduplication_config" json:"deduplication_config"` // 限流配置
}

func (m MessageTemplate) TableName() string {
	return "message_template"
}
