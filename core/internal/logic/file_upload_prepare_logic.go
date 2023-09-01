package logic

import (
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrePareRequest) (resp *types.FileUploadPrePareReply, err error) {
	// todo: add your logic here and delete this line
	rp := new(models.RepositoryPool)
	result := l.svcCtx.DB.Where("hash = ?", req.Md5).First(rp)
	if result.Error != nil {
		return nil, result.Error
	}
	resp = &types.FileUploadPrePareReply{}
	if result.RowsAffected != 0 {
		//妙传成功
		resp.Identity = rp.Identity
	} else {
		//获取该文件的UploadID、key，用来进行文件的分片上传

	}
	return
}
