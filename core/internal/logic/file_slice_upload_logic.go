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

type FileSliceUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileSliceUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileSliceUploadLogic {
	return &FileSliceUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileSliceUploadLogic) FileSliceUpload(req *types.FileSliceUploadRequest) (resp *types.FileSliceUploadReply, err error) {
	// todo: add your logic here and delete this line
	// todo: add your logic here and delete this line
	rp := &models.RepositoryPool{
		Identity:   helper.GetUUID(),
		Hash:       req.Hash,
		Name:       req.Name,
		Ext:        req.Ext,
		Size:       req.Size,
		Path:       req.Path,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	err = l.svcCtx.DB.Create(rp).Error
	if err != nil {
		return nil, err
	}
	resp = &types.FileSliceUploadReply{}
	resp.Identity = rp.Identity
	resp.Ext = rp.Ext
	resp.Name = rp.Name
	return
}
