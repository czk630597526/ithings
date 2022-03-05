package dm

import (
	"context"
	"github.com/go-things/things/shared/errors"
	"github.com/go-things/things/shared/utils"
	"github.com/go-things/things/src/dmsvr/dm"

	"github.com/go-things/things/src/webapi/internal/svc"
	"github.com/go-things/things/src/webapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeviceDescribeLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeviceDescribeLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDeviceDescribeLogLogic {
	return GetDeviceDescribeLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeviceDescribeLogLogic) GetDeviceDescribeLog(req types.GetDeviceDescribeLogReq) (*types.GetDeviceDescribeLogResp, error) {
	l.Infof("GetDeviceDescribeLog|req=%+v", req)
	resp, err := l.svcCtx.DmRpc.GetDeviceDescribeLog(l.ctx, &dm.GetDeviceDescribeLogReq{
		DeviceName: req.DeviceName,
		ProductID:  req.ProductID,
		TimeStart:  req.TimeStart,
		TimeEnd:    req.TimeEnd,
		Limit:      req.Limit,
	})
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s|rpc.GetDeviceDescribeLog|req=%v|err=%+v", utils.FuncName(), req, er)
		return nil, er
	}
	info := make([]*types.DeviceDescribeLog, 0, len(resp.Data))
	for _, v := range resp.Data {
		info = append(info, &types.DeviceDescribeLog{
			Timestamp:  v.Timestamp,
			Action:     v.Action,
			RequestID:  v.RequestID,
			TranceID:   v.TranceID,
			Topic:      v.Topic,
			Content:    v.Content,
			ResultType: v.ResultType,
		})
	}
	return &types.GetDeviceDescribeLogResp{Data: info}, err
}