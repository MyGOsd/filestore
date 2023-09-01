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

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	// todo: add your logic here and delete this line
	//parentID
	parentDate := new(models.UserRepository)
	result := l.svcCtx.DB.Where("identity = ? AND user_Identity = ?", req.ParentIdentity, userIdentity).Find(&parentDate)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("该文件夹不存在")
	}
	//更新记录的parentId
	result = l.svcCtx.DB.Where("identity = ?", req.Identity).Model(&models.UserRepository{}).
		Updates(&models.UserRepository{
			ParentId:   int64(parentDate.Id),
			Updated_at: time.Now(),
		})
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
