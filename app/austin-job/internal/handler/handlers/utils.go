package handlers

import (
	"austin-go/app/austin-common/repo"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
)

func getContentModel(contentModel interface{}, v interface{}) {
	marshal, _ := jsonx.Marshal(contentModel)
	_ = jsonx.Unmarshal(marshal, &v)
}

func getAccount(ctx context.Context, sendAccount int, v interface{}) error {
	accountRepo := repo.NewSendAccountRepo()
	one, err := accountRepo.One(ctx, int64(sendAccount))
	if err != nil {
		return err
	}
	if one == nil {
		return fmt.Errorf("获取账号异常 sendAccount: %d", sendAccount)
	}

	err = jsonx.Unmarshal([]byte(one.Config), &v)
	if err != nil {
		return err
	}
	return nil
}
