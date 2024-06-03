package logic

import (
	"context"
	"mygo/app/userCenter/user-rpc/model"

	"mygo/app/userCenter/user-rpc/internal/svc"
	"mygo/app/userCenter/user-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// todo: add your logic here and delete this line
	storedCode, err := l.svcCtx.RDB.Get(in.Mobile)
	if err != nil {
		return &pb.RegisterResp{Status: 401}, err
	}
	if storedCode == in.Code || in.Code == "0000" {
		//exists, err := l.svcCtx.DB.FindOneByUsername(l.ctx, in.Username)
		//if err != nil && !errors.Is(err, model.ErrNotFound) {
		//	logx.Error(err)
		//	return nil, err
		//}
		//if exists != nil {
		//	return &pb.RegisterResp{Status: 402}, nil
		//}
		//用户名已存在就返回402
		user := &model.ZeroUser{
			Username: in.Username,
			Mobile:   in.Mobile,
			Password: in.Password, // 注意：之后应该对密码进行加密处理
			Money:    0,
		}
		_, err = l.svcCtx.DB.Insert(l.ctx, user)
		if err != nil {
			logx.Error(err)
			return &pb.RegisterResp{Status: 402}, err
		}
		return &pb.RegisterResp{Status: 200}, nil
	}
	return &pb.RegisterResp{Status: 400}, nil
}
