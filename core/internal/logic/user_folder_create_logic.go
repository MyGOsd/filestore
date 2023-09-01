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

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateReply, err error) {
	// todo: add your logic here and delete this line
	//判断当前名称在该层级下是否存在
	data := models.UserRepository{}
	result := l.svcCtx.DB.Where("name = ? AND parent_id = ?", req.Name, req.ParentId).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected != 0 {
		return nil, errors.New("该文件名已存在")
	}
	//创建文件夹
	fold := &models.UserRepository{
		Identity:     helper.GetUUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
	}
	result = l.svcCtx.DB.Create(&fold)
	if result.Error != nil {
		return nil, result.Error
	}
	resp = &types.UserFolderCreateReply{
		Identity: fold.Identity,
	}
	return
}
