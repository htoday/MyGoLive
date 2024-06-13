package logic

import (
	"context"

	"mygo/app/userCenter/user-rpc/internal/svc"
	"mygo/app/userCenter/user-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChargeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChargeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChargeLogic {
	return &ChargeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChargeLogic) Charge(in *pb.ChargeReq) (*pb.ChargeResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.DB.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		logx.Error(err)
		return &pb.ChargeResp{
			Status: 401,
		}, err
	}
	user.Money += in.Num
	err = l.svcCtx.DB.Update(l.ctx, user)
	if err != nil {
		logx.Error(err)
		return &pb.ChargeResp{
			Status: 401,
		}, err
	}
	return &pb.ChargeResp{
		Status: 200,
	}, nil
}
