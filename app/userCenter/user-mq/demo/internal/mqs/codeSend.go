package mqs

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"mygo/app/userCenter/user-mq/demo/internal/svc"
)

type PaymentSuccess struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentSuccess {
	return &PaymentSuccess{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentSuccess) Consume(key, val string) error {
	logx.Infof("PaymentSuccess key :%s , val :%s", key, val)
	fmt.Printf("Send Code %s and key is %s", val, key)
	return nil
}
