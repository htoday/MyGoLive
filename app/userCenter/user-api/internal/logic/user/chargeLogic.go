package user

import (
	"context"

	"mygo/app/userCenter/user-api/internal/svc"
	"mygo/app/userCenter/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChargeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChargeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChargeLogic {
	return &ChargeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChargeLogic) Charge(req *types.ChargeReq) (resp *types.ChargeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
