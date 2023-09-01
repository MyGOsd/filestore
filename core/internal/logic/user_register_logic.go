package logic

import (
	"context"
	"errors"
	"time"

	"cloud_disk/core/helper"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequset) (resp *types.UserRegisterReply, err error) {
	// todo: add your logic here and delete this line
	//判断Code是否一致
	code, err := l.svcCtx.RedisDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("未获取该邮箱的验证码")
	}
	if code != req.Code {
		err = errors.New("输入的验证码错误")
		return
	}
	//判断用户是否已存在
	cnt := l.svcCtx.DB.Where("name = ?", req.Name).Find(new(models.UserBasic)).RowsAffected
	if cnt > 0 {
		err = errors.New("用户名已存在")
		return
	}
	//数据入库
	user := models.UserBasic{
		Identity:   helper.GetUUID(),
		Name:       req.Name,
		Email:      req.Email,
		Password:   helper.Md5(req.Password),
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	result := l.svcCtx.DB.Create(&user)
	if result.RowsAffected == 0 {
		return nil, err
	}
	return
}
