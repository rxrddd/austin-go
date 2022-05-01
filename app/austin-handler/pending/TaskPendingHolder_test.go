package pending

import (
	"austin-go/app/austin-common/enums/channelType"
	"austin-go/app/austin-common/types"
	"testing"
	"time"
)

func TestNewTaskPendingHolder(t *testing.T) {
	go func() {
		for {
			time.Sleep(time.Second)
			Submit("official_accounts.notice", Task{TaskInfo: types.TaskInfo{
				SendChannel: channelType.OfficialAccounts,
			}})
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second / 2)
			Submit("sms.marketing", Task{TaskInfo: types.TaskInfo{
				SendChannel: channelType.Sms,
			}})
		}
	}()
	select {}
}
