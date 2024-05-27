package user

import (
	"context"
	"github.com/jinzhu/copier"
	"mygo/app/userCenter/user-rpc/userservice"

	"mygo/app/userCenter/user-api/internal/svc"
	"mygo/app/userCenter/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	registerResp, err := l.svcCtx.UserRpcClient.Register(l.ctx, &userservice.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		Code:     req.Code,
		Mobile:   req.Mobile,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, registerResp)
	return resp, nil
}
