package user

import (
	"context"
	"mygo/app/userCenter/user-rpc/userservice"

	"mygo/app/userCenter/user-api/internal/svc"
	"mygo/app/userCenter/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageLogic) SendMessage(req *types.SendMessageReq) (resp *types.SendMessageResp, err error) {
	// todo: add your logic here and delete this line
	sendCodeResp, err := l.svcCtx.UserRpcClient.GetCode(l.ctx, &userservice.GetCodeReq{
		Mobile: req.Mobile,
	})
	if err != nil {
		return nil, err
	}

	return &types.SendMessageResp{
		Status: sendCodeResp.Status,
	}, nil
}
