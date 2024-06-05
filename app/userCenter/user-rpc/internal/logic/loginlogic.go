package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"mygo/app/userCenter/user-rpc/internal/svc"
	"mygo/app/userCenter/user-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.DB.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	// 检查密码是否正确
	if user.Password != in.Password { // 注意：在实际应用中，你应该对密码进行加密处理
		return nil, errors.New("invalid password")
	}

	token, err := generateToken(in.Username)
	if err != nil {
		return nil, err
	}
	fmt.Println(token, "我是token2")
	return &pb.LoginResp{
		Status:       200,
		Token:        token,
		ExpireTime:   time.Now().Add(time.Hour * 24).UnixMilli(), // 你需要根据你的需求来设置这个值
		RefreshAfter: time.Now().Add(time.Hour * 12).UnixMilli(), // 你需要根据你需求来设置这个值
	}, nil
}

func generateToken(username string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "htoday",
		"sub": username,
		"exp": time.Now().Add(time.Hour * 24).UnixMilli(),
	})
	secretKey := []byte("666666")
	//fmt.Println(token, "我是token1")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(tokenString)
		fmt.Println(err)
	}
	return tokenString, err
}
