package logic

import (
	"context"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUpLoadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUpLoadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUpLoadLogic {
	return &FileUpLoadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUpLoadLogic) FileUpLoad(req *types.FileUpLoadRequest) (resp *types.FileUpLoadReply, err error) {
	rp := &models.RepositoryPool{
		Identity: helper.UUID(),
		Hash:     req.Hash,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     req.Path,
	}
	_, err = l.svcCtx.Engine.Insert(rp)
	if err != nil {
		return nil, err
	}

	resp = new(types.FileUpLoadReply)
	resp.Identity = rp.Identity
	return
}
