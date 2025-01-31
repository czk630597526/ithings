// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/logic/devicemanage"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
)

type DeviceManageServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedDeviceManageServer
}

func NewDeviceManageServer(svcCtx *svc.ServiceContext) *DeviceManageServer {
	return &DeviceManageServer{
		svcCtx: svcCtx,
	}
}

// 新增设备
func (s *DeviceManageServer) DeviceInfoCreate(ctx context.Context, in *dm.DeviceInfo) (*dm.Response, error) {
	l := devicemanagelogic.NewDeviceInfoCreateLogic(ctx, s.svcCtx)
	return l.DeviceInfoCreate(in)
}

// 更新设备
func (s *DeviceManageServer) DeviceInfoUpdate(ctx context.Context, in *dm.DeviceInfo) (*dm.Response, error) {
	l := devicemanagelogic.NewDeviceInfoUpdateLogic(ctx, s.svcCtx)
	return l.DeviceInfoUpdate(in)
}

// 删除设备
func (s *DeviceManageServer) DeviceInfoDelete(ctx context.Context, in *dm.DeviceInfoDeleteReq) (*dm.Response, error) {
	l := devicemanagelogic.NewDeviceInfoDeleteLogic(ctx, s.svcCtx)
	return l.DeviceInfoDelete(in)
}

// 获取设备信息列表
func (s *DeviceManageServer) DeviceInfoIndex(ctx context.Context, in *dm.DeviceInfoIndexReq) (*dm.DeviceInfoIndexResp, error) {
	l := devicemanagelogic.NewDeviceInfoIndexLogic(ctx, s.svcCtx)
	return l.DeviceInfoIndex(in)
}

// 获取设备信息详情
func (s *DeviceManageServer) DeviceInfoRead(ctx context.Context, in *dm.DeviceInfoReadReq) (*dm.DeviceInfo, error) {
	l := devicemanagelogic.NewDeviceInfoReadLogic(ctx, s.svcCtx)
	return l.DeviceInfoRead(in)
}

// 创建分组设备
func (s *DeviceManageServer) DeviceGatewayMultiCreate(ctx context.Context, in *dm.DeviceGatewayMultiCreateReq) (*dm.Response, error) {
	l := devicemanagelogic.NewDeviceGatewayMultiCreateLogic(ctx, s.svcCtx)
	return l.DeviceGatewayMultiCreate(in)
}

// 获取分组设备信息列表
func (s *DeviceManageServer) DeviceGatewayIndex(ctx context.Context, in *dm.DeviceGatewayIndexReq) (*dm.DeviceGatewayIndexResp, error) {
	l := devicemanagelogic.NewDeviceGatewayIndexLogic(ctx, s.svcCtx)
	return l.DeviceGatewayIndex(in)
}

// 删除分组设备
func (s *DeviceManageServer) DeviceGatewayMultiDelete(ctx context.Context, in *dm.DeviceGatewayMultiDeleteReq) (*dm.Response, error) {
	l := devicemanagelogic.NewDeviceGatewayMultiDeleteLogic(ctx, s.svcCtx)
	return l.DeviceGatewayMultiDelete(in)
}

// 设备计数
func (s *DeviceManageServer) DeviceInfoCount(ctx context.Context, in *dm.DeviceInfoCountReq) (*dm.DeviceInfoCountResp, error) {
	l := devicemanagelogic.NewDeviceInfoCountLogic(ctx, s.svcCtx)
	return l.DeviceInfoCount(in)
}

// 设备类型
func (s *DeviceManageServer) DeviceTypeCount(ctx context.Context, in *dm.DeviceTypeCountReq) (*dm.DeviceTypeCountResp, error) {
	l := devicemanagelogic.NewDeviceTypeCountLogic(ctx, s.svcCtx)
	return l.DeviceTypeCount(in)
}
