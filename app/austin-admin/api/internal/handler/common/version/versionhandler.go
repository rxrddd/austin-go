package version

import (
	"net/http"

	"austin-go/app/austin-admin/api/internal/logic/common/version"
	"austin-go/app/austin-admin/api/internal/svc"

	"austin-go/common/result"
)

func VersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := version.NewVersionLogic(r.Context(), svcCtx)
		resp, err := l.Version()
		result.HttpResult(r, w, resp, err)
	}
}
