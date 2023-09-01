package logic

import (
	"cloud_disk/core/define"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {
	uf := make([]*types.UserFile, 0)
	//分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size
	// Id                 int64  `json:"id"`
	// Identity           string `json:"identity"`
	// RepositoryIdentity string `json:"repositoryIdentity"`
	// Name               string `json:"name"`
	// Size               int64  `json:"size"`
	// Ext                string `json:"ext"`
	// Path               string `json:"Path"`
	result := l.svcCtx.DB.Debug().
		Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).Table("user_repository").
		Select("user_repository.Id,user_repository.identity,user_repository.repository_Identity,user_repository.name,repository_pool.size,user_repository.ext,repository_pool.path").
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Limit(size).Offset(offset).Scan(&uf)
	if result.Error != nil {
		return
	}
	//查询文件总数
	resp = &types.UserFileListReply{
		List:  uf,
		Count: result.RowsAffected,
	}
	return
}
