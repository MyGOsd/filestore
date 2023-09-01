package logic

import (
	"context"
	"time"

	"cloud_disk/core/helper"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	// todo: add your logic here and delete this line
	ur := new(models.UserRepository)
	result := l.svcCtx.DB.Where("identity = ?", req.UserRepositoryIdentity).First(ur)
	if result.Error != nil {
		return nil, result.Error
	}
	data := &models.ShareBasic{
		Identity:               helper.GetUUID(),
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity:     ur.RepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
		Created_at:             time.Now(),
		Updated_at:             time.Now(),
	}
	result = l.svcCtx.DB.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	resp = &types.ShareBasicCreateReply{
		Identity: data.Identity,
	}
	return
}
