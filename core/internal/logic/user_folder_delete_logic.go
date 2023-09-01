package logic

import (
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderDeleteLogic {
	return &UserFolderDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderDeleteLogic) UserFolderDelete(req *types.UserFolderDeleteRequest, userIdentity string) (resp *types.UserFolderDeleteReply, err error) {
	// todo: add your logic here and delete this line
	result := l.svcCtx.DB.Debug().Where("user_Identity = ? AND identity = ?", userIdentity, req.Identity).Delete(&models.UserRepository{})
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
