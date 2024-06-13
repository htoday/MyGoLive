package logic

import (
	"context"
	"fmt"

	"mygo/app/userCenter/user-rpc/internal/svc"
	"mygo/app/userCenter/user-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SpendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSpendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SpendLogic {
	return &SpendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SpendLogic) Spend(in *pb.SpendReq) (*pb.SpendResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.DB.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		logx.Error(err)
		return &pb.SpendResp{
			Status: 401,
		}, err
	}
	if user.Money < in.Num {
		return &pb.SpendResp{
			Status: 402,
		}, nil
	}
	user.Money -= in.Num
	err = l.svcCtx.DB.Update(l.ctx, user)
	fmt.Println(err)
	if err != nil {
		logx.Error(err)
		return &pb.SpendResp{
			Status: 401,
		}, err
	}
	return &pb.SpendResp{
		Status: 200,
	}, nil
}
