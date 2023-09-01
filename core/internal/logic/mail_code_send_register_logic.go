package logic

import (
	"context"
	"errors"
	"time"

	"cloud_disk/core/define"
	"cloud_disk/core/helper"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequset) (resp *types.MailCodeSendReply, err error) {
	// todo: add your logic here and delete this line
	//该邮箱在数据库不存在

	cnt := l.svcCtx.DB.Where("email = ?", req.Email).Find(new(models.UserBasic)).RowsAffected
	if cnt > 0 {
		err = errors.New("该邮箱已被注册")
		return
	}
	//生成随机验证码
	code := helper.RandCode()
	//存储验证码
	l.svcCtx.RedisDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))
	err = helper.MailSendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return
}
