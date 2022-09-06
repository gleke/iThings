package productmanagelogic

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/events"
	"github.com/i-Things/things/shared/utils"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductInfoDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductInfoDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductInfoDeleteLogic {
	return &ProductInfoDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除设备
func (l *ProductInfoDeleteLogic) ProductInfoDelete(in *dm.ProductInfoDeleteReq) (*dm.Response, error) {
	err := l.DropProduct(in)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.DmDB.Delete(l.ctx, in.ProductID)
	if err != nil {
		l.Errorf("%s.Delete err=%+v", utils.FuncName(), err)
		return nil, errors.Database.AddDetail(err)
	}
	err = l.svcCtx.DataUpdate.ProductSchemaUpdate(l.ctx, &events.DataUpdateInfo{ProductID: in.ProductID})
	if err != nil {
		return nil, err
	}

	return &dm.Response{}, nil
}
func (l *ProductInfoDeleteLogic) DropProduct(in *dm.ProductInfoDeleteReq) error {
	pt, err := l.svcCtx.SchemaRepo.GetSchemaModel(l.ctx, in.ProductID)
	if err != nil {
		return errors.System.AddDetail(err)
	}
	err = l.svcCtx.HubLogRepo.DropProduct(l.ctx, in.ProductID)
	if err != nil {
		l.Errorf("%s.HubLogRepo.DropProduct err=%v", utils.FuncName(), err)
		return errors.Database.AddDetail(err)
	}
	err = l.svcCtx.SDKLogRepo.DropProduct(l.ctx, in.ProductID)
	if err != nil {
		l.Errorf("%s.SDKLogRepo.DropProduct err=%v", utils.FuncName(), err)
		return errors.Database.AddDetail(err)
	}
	err = l.svcCtx.SchemaManaRepo.DropProduct(l.ctx, pt, in.ProductID)
	if err != nil {
		l.Errorf("%s.SchemaManaRepo.DropProduct err=%+v", utils.FuncName(), err)
		return errors.Database.AddDetail(err)
	}
	err = l.svcCtx.SchemaRepo.ClearCache(l.ctx, in.ProductID)
	if err != nil {
		l.Errorf("%s.SchemaRepo.ClearCache err=%+v", utils.FuncName(), err)
		return errors.Database.AddDetail(err)
	}
	return nil
}
