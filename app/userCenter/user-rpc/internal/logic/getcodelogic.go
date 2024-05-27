package logic

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"mygo/app/userCenter/user-rpc/internal/svc"
	"mygo/app/userCenter/user-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodeLogic {
	return &GetCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCodeLogic) GetCode(in *pb.GetCodeReq) (*pb.GetCodeResp, error) {
	// todo: add your logic here and delete this line
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(9000) + 1000
	randomNumberStr := fmt.Sprintf("%04d", randomNumber)
	err := l.svcCtx.RDB.Setex(in.Mobile, randomNumberStr, 300)
	if err != nil {
		return &pb.GetCodeResp{Status: 400}, err
	}
	if err := l.svcCtx.KqPusherClient.Push(randomNumberStr); err != nil {
		logx.Errorf("KqPusherClient Push Error , err :%v", err)
		return &pb.GetCodeResp{Status: 400}, err
	} //这里是发送验证码到kafka
	return &pb.GetCodeResp{
		Status: 200,
	}, nil
}
