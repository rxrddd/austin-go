package channelType

const (
	Im                 int = 10
	Push               int = 20
	Sms                int = 30
	Email              int = 40
	OfficialAccounts   int = 50
	MiniProgram        int = 60
	EnterpriseWeChat   int = 70
	DingDing           int = 80
	DingDingWorkNotice int = 90
)

var (
	TypeText = map[int]string{
		Im:                 "IM(站内信)",
		Push:               "push(通知栏)",
		Sms:                "sms(短信)",
		Email:              "email(邮件)",
		OfficialAccounts:   "OfficialAccounts(服务号)",
		MiniProgram:        "miniProgram(小程序)",
		EnterpriseWeChat:   "EnterpriseWeChat(企业微信)",
		DingDing:           "dingDingRobot(钉钉机器人)",
		DingDingWorkNotice: "dingDingWorkNotice(钉钉工作通知)",
	}
	TypeCodeEn = map[int]string{
		Im:                 "im",
		Push:               "push",
		Sms:                "sms",
		Email:              "email",
		OfficialAccounts:   "official_accounts",
		MiniProgram:        "mini_program",
		EnterpriseWeChat:   "enterprise_we_chat",
		DingDing:           "ding_ding_robot",
		DingDingWorkNotice: "ding_ding_work_notice",
	}
)
