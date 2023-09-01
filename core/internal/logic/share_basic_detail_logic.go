package logic

import (
	"context"
	"time"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailReply, err error) {
	//对分享记录的点击次数进行+1操作
	result := l.svcCtx.DB.Model(&models.ShareBasic{}).Where("identity = ?", req.Identity).Updates(map[string]interface{}{"click_num": gorm.Expr("click_num + ?", 1), "updated_at": time.Now()})
	if result.Error != nil {
		return nil, result.Error
	}
	//获取资源的详细信息
	resp = &types.ShareBasicDetailReply{}
	result = l.svcCtx.DB.Debug().Model(&models.ShareBasic{}).Select("repository_pool.identity,user_repository.name,repository_pool.ext,repository_pool.size,repository_pool.path").
		Joins("left join repository_pool on share_basic.repository_identity = repository_pool.identity").
		Joins("left join user_repository on user_repository.identity = share_basic.user_repository_identity").
		Where("share_basic.identity = ?", req.Identity).First(resp)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
