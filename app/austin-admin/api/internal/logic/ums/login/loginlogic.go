package login

import (
	"austin-go/app/austin-admin/api/internal/svc"
	"austin-go/app/austin-admin/api/internal/types"
	"austin-go/common/xerr"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginResp, err error) {

	info, err := l.svcCtx.AccountRepo.FindByUserName(l.ctx, req.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.NewErrMsg("用户或密码错误")
	}

	if !info.CheckPassword(req.Password) {
		return nil, xerr.NewErrMsg("用户或密码错误")
	}

	resp = new(types.LoginResp)
	now := time.Now().Unix()
	token, _ := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, info.ID)
	if err != nil {
		return nil, xerr.NewErrMsg("登录token生成异常")
	}
	resp.Token = token
	resp.UserInfo = types.UserInfo{
		UserID:    info.ID,
		UserName:  info.Nickname,
		Dashboard: "0",
		Role: []string{
			"SA",
			"admin",
			"Auditor",
		},
	}
	return
}
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
