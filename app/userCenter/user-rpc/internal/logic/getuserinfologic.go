package logic

import (
	"context"
	"mygo/app/userCenter/user-rpc/internal/svc"
	"mygo/app/userCenter/user-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.DB.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &pb.GetUserInfoResp{
		UserId:   user.UserId,
		Mobile:   user.Mobile,
		Money:    user.Money,
		Status:   200,
		Username: user.Username,
	}, nil
}
