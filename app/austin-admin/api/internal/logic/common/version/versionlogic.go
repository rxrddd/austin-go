package version

import (
	"context"

	"austin-go/app/austin-admin/api/internal/svc"
	"austin-go/app/austin-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) VersionLogic {
	return VersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VersionLogic) Version() (resp *types.VersionResp, err error) {
	return &types.VersionResp{Version: "1.5.0"}, nil
}
