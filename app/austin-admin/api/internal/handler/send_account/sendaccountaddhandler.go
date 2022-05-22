package send_account

import (
	"net/http"

	"austin-go/app/austin-admin/api/internal/logic/send_account"
	"austin-go/app/austin-admin/api/internal/svc"
	"austin-go/app/austin-admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"austin-go/common/result"
)

func SendAccountAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendAccountAddReq
		if err := httpx.Parse(r, &req); err != nil {
			result.HttpResult(r, w, nil, err)
			return
		}

		l := send_account.NewSendAccountAddLogic(r.Context(), svcCtx)
		err := l.SendAccountAdd(req)
		result.HttpResult(r, w, nil, err)
	}
}
