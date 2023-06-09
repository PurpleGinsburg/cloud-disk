package handler

import (
	"crypto/md5"
	"net/http"
	"path"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUpLoadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUpLoadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		//判断文件是否已存在
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		//md5.Sum(b)返回一个16进制的数组,转换
		md5 := md5.Sum(b)
		m := make([]byte, len(md5))
		copy(m, md5[:])
		hash := string(m) //hex.DecodeString(md5.Sum(b)) //fmt.Sprint("%v", md5.Sum(b))

		rp := new(models.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			return
		}
		if has {
			httpx.OkJson(w, &types.FileUpLoadReply{Identity: rp.Identity})
			return
		}

		//文件不存在，往COS中存储文件
		cosPath, err := helper.CosUpload(r)
		if err != nil {
			return
		}

		//往logic中传递req
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = cosPath

		l := logic.NewFileUpLoadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpLoad(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
