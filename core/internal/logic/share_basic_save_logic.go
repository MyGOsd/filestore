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

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveReply, err error) {
	// todo: add your logic here and delete this line
	//获取资源详情
	rp := new(models.RepositoryPool)
	result := l.svcCtx.DB.Where("identity = ?", req.RepositoryIdentity).First(&rp)
	if result.RowsAffected == 0 {
		return nil, errors.New("该资源不存在")
	}
	//user_repository 资源保存
	ur := &models.UserRepository{
		Identity:           helper.GetUUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                rp.Ext,
		Name:               rp.Name,
		Created_at:         time.Now(),
		Updated_at:         time.Now(),
	}
	result = l.svcCtx.DB.Create(&ur)
	if result.Error != nil {
		return nil, result.Error
	}
	resp = &types.ShareBasicSaveReply{
		Identity: ur.Identity,
	}
	return
}
