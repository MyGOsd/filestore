package handler

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud_disk/core/helper"
	"cloud_disk/core/internal/logic"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequset
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		file, fileHeard, err := r.FormFile("file")
		if err != nil {
			return
		}
		//判断文件是否已存在
		b := make([]byte, fileHeard.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := models.RepositoryPool{}
		result := svcCtx.DB.Where("hash = ?", hash).Find(&rp)
		if result.RowsAffected != 0 {
			httpx.OkJson(w, &types.FileUploadReply{
				Identity: rp.Identity,
				Ext:      rp.Ext,
				Name:     rp.Name,
			})
			return
		}
		//文件不存在，往oss中存储文件
		ossPath, err := helper.OssUpload(r)
		if err != nil {
			return
		}
		//往logic中传递request
		req.Name = fileHeard.Filename
		req.Ext = path.Ext(fileHeard.Filename)
		req.Hash = hash
		req.Size = fileHeard.Size
		req.Path = ossPath
		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
