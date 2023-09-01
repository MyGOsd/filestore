package logic

import (
	"context"
	"errors"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequset) (resp *types.UserDetailReply, err error) {
	resp = &types.UserDetailReply{}
	ub := models.UserBasic{}
	result := l.svcCtx.DB.Where("identity=?", req.Identity).First(&ub)
	if result.RowsAffected == 0 {
		return nil, errors.New("该用户不存在")
	}
	resp.Name = ub.Name
	resp.Email = ub.Email
	return
}
