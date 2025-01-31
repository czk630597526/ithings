package interact

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/disvr/pb/di"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PropertyReadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPropertyReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PropertyReadLogic {
	return &PropertyReadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PropertyReadLogic) PropertyRead(req *types.DeviceInteractRespReadReq) (resp *types.DeviceInteractSendPropertyResp, err error) {
	dmReq := &di.RespReadReq{
		ProductID:   req.ProductID,
		DeviceName:  req.DeviceName,
		ClientToken: req.ClientToken,
	}
	dmResp, err := l.svcCtx.DeviceInteract.PropertyRead(l.ctx, dmReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.PropertyRead req=%v err=%+v", utils.FuncName(), req, er)
		return nil, er
	}
	return &types.DeviceInteractSendPropertyResp{
		Code:        dmResp.Code,
		Status:      dmResp.Status,
		ClientToken: dmResp.ClientToken,
		Data:        dmResp.Data,
	}, nil
}
