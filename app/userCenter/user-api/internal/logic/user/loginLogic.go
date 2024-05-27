package user

import (
	"context"
	"mygo/app/userCenter/user-rpc/userservice"

	"mygo/app/userCenter/user-api/internal/svc"
	"mygo/app/userCenter/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	loginResp, err := l.svcCtx.UserRpcClient.Login(l.ctx, &userservice.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	//_ = copier.Copy(resp, loginResp)
	return &types.LoginResp{
		Status: int(loginResp.Status),
	}, nil
}
