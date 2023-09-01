package logic

import (
	"context"
	"errors"
	"time"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileUpdateLogic {
	return &UserFileUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileUpdateLogic) UserFileUpdate(req *types.UserFileUpdateRequest, userIdentity string) (resp *types.UserFileUpdateReply, err error) {
	// todo: add your logic here and delete this line
	//判断当前名称在该层级下是否存在
	data := models.UserRepository{}
	result := l.svcCtx.DB.Where("name = ? AND parent_id = (SELECT parent_id FROM user_repository ur WHERE ur.identity = ?)", req.Name, req.Identity).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected != 0 {
		return nil, errors.New("该文件名已存在")
	}
	err = l.svcCtx.DB.Model(&models.UserRepository{}).Select("name", "Updated_at").Where("identity = ? AND user_Identity = ?", req.Identity, userIdentity).Updates(models.UserRepository{
		Name:       req.Name,
		Updated_at: time.Now(),
	}).Error
	if err != nil {
		return nil, err
	}
	return
}
