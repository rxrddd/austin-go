package handler

import (
	"austin-go/common/result"
	"net/http"

	"austin-go/app/austin-web/api/internal/logic"
	"austin-go/app/austin-web/api/internal/svc"
	"austin-go/app/austin-web/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSendLogic(r.Context(), svcCtx)
		resp, err := l.Send(req)
		result.HttpResult(r, w, resp, err)
	}
}
