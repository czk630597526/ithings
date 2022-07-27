// Code generated by goctl. DO NOT EDIT!
// Source: dm.proto

package dm

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AccessAuthReq            = dm.AccessAuthReq
	DataHubLogIndex          = dm.DataHubLogIndex
	DataHubLogIndexReq       = dm.DataHubLogIndexReq
	DataHubLogIndexResp      = dm.DataHubLogIndexResp
	DataSchemaIndex          = dm.DataSchemaIndex
	DataSchemaIndexResp      = dm.DataSchemaIndexResp
	DataSchemaLatestIndexReq = dm.DataSchemaLatestIndexReq
	DataSchemaLogIndexReq    = dm.DataSchemaLogIndexReq
	DataSdkLogIndex          = dm.DataSdkLogIndex
	DataSdkLogIndexReq       = dm.DataSdkLogIndexReq
	DataSdkLogIndexResp      = dm.DataSdkLogIndexResp
	DeviceInfo               = dm.DeviceInfo
	DeviceInfoDeleteReq      = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq       = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp      = dm.DeviceInfoIndexResp
	DeviceInfoReadReq        = dm.DeviceInfoReadReq
	FirmwareInfo             = dm.FirmwareInfo
	GetFirmwareInfoReq       = dm.GetFirmwareInfoReq
	GetFirmwareInfoResp      = dm.GetFirmwareInfoResp
	LoginAuthReq             = dm.LoginAuthReq
	ManageFirmwareReq        = dm.ManageFirmwareReq
	PageInfo                 = dm.PageInfo
	ProductInfo              = dm.ProductInfo
	ProductInfoDeleteReq     = dm.ProductInfoDeleteReq
	ProductInfoIndexReq      = dm.ProductInfoIndexReq
	ProductInfoIndexResp     = dm.ProductInfoIndexResp
	ProductInfoReadReq       = dm.ProductInfoReadReq
	ProductSchema            = dm.ProductSchema
	ProductSchemaReadReq     = dm.ProductSchemaReadReq
	ProductSchemaUpdateReq   = dm.ProductSchemaUpdateReq
	Response                 = dm.Response
	RootCheckReq             = dm.RootCheckReq
	SendActionReq            = dm.SendActionReq
	SendActionResp           = dm.SendActionResp
	SendPropertyReq          = dm.SendPropertyReq
	SendPropertyResp         = dm.SendPropertyResp

	Dm interface {
		// 新增设备
		DeviceInfoCreate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error)
		// 更新设备
		DeviceInfoUpdate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error)
		// 删除设备
		DeviceInfoDelete(ctx context.Context, in *DeviceInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
		// 获取设备信息列表
		DeviceInfoIndex(ctx context.Context, in *DeviceInfoIndexReq, opts ...grpc.CallOption) (*DeviceInfoIndexResp, error)
		// 获取设备信息详情
		DeviceInfoRead(ctx context.Context, in *DeviceInfoReadReq, opts ...grpc.CallOption) (*DeviceInfo, error)
		// 新增设备
		ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error)
		// 更新设备
		ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error)
		// 删除设备
		ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
		// 获取设备信息列表
		ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error)
		// 获取设备信息详情
		ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error)
		// 更新产品物模型
		ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error)
		// 获取产品物模型
		ProductSchemaRead(ctx context.Context, in *ProductSchemaReadReq, opts ...grpc.CallOption) (*ProductSchema, error)
		// 管理产品的固件
		ManageFirmware(ctx context.Context, in *ManageFirmwareReq, opts ...grpc.CallOption) (*Response, error)
		// 获取产品固件信息
		GetFirmwareInfo(ctx context.Context, in *GetFirmwareInfoReq, opts ...grpc.CallOption) (*GetFirmwareInfoResp, error)
		// 设备登录认证
		LoginAuth(ctx context.Context, in *LoginAuthReq, opts ...grpc.CallOption) (*Response, error)
		// 设备操作认证
		AccessAuth(ctx context.Context, in *AccessAuthReq, opts ...grpc.CallOption) (*Response, error)
		// 鉴定是否是root账号
		RootCheck(ctx context.Context, in *RootCheckReq, opts ...grpc.CallOption) (*Response, error)
		// 同步调用设备行为
		SendAction(ctx context.Context, in *SendActionReq, opts ...grpc.CallOption) (*SendActionResp, error)
		// 同步调用设备属性
		SendProperty(ctx context.Context, in *SendPropertyReq, opts ...grpc.CallOption) (*SendPropertyResp, error)
		// 获取设备sdk调试日志
		DataSdkLogIndex(ctx context.Context, in *DataSdkLogIndexReq, opts ...grpc.CallOption) (*DataSdkLogIndexResp, error)
		// 获取设备调试信息记录登入登出,操作
		DataHubLogIndex(ctx context.Context, in *DataHubLogIndexReq, opts ...grpc.CallOption) (*DataHubLogIndexResp, error)
		// 获取设备数据信息
		DataSchemaLatestIndex(ctx context.Context, in *DataSchemaLatestIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error)
		// 获取设备数据信息
		DataSchemaLogIndex(ctx context.Context, in *DataSchemaLogIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error)
	}

	defaultDm struct {
		cli zrpc.Client
	}

	directDm struct {
		svcCtx *svc.ServiceContext
		svr    dm.DmServer
	}
)

func NewDm(cli zrpc.Client) Dm {
	return &defaultDm{
		cli: cli,
	}
}

func NewDirectDm(svcCtx *svc.ServiceContext, svr dm.DmServer) Dm {
	return &directDm{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 新增设备
func (m *defaultDm) DeviceInfoCreate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.DeviceInfoCreate(ctx, in, opts...)
}

// 新增设备
func (d *directDm) DeviceInfoCreate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.DeviceInfoCreate(ctx, in)
}

// 更新设备
func (m *defaultDm) DeviceInfoUpdate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.DeviceInfoUpdate(ctx, in, opts...)
}

// 更新设备
func (d *directDm) DeviceInfoUpdate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.DeviceInfoUpdate(ctx, in)
}

// 删除设备
func (m *defaultDm) DeviceInfoDelete(ctx context.Context, in *DeviceInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.DeviceInfoDelete(ctx, in, opts...)
}

// 删除设备
func (d *directDm) DeviceInfoDelete(ctx context.Context, in *DeviceInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.DeviceInfoDelete(ctx, in)
}

// 获取设备信息列表
func (m *defaultDm) DeviceInfoIndex(ctx context.Context, in *DeviceInfoIndexReq, opts ...grpc.CallOption) (*DeviceInfoIndexResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.DeviceInfoIndex(ctx, in, opts...)
}

// 获取设备信息列表
func (d *directDm) DeviceInfoIndex(ctx context.Context, in *DeviceInfoIndexReq, opts ...grpc.CallOption) (*DeviceInfoIndexResp, error) {
	return d.svr.DeviceInfoIndex(ctx, in)
}

// 获取设备信息详情
func (m *defaultDm) DeviceInfoRead(ctx context.Context, in *DeviceInfoReadReq, opts ...grpc.CallOption) (*DeviceInfo, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.DeviceInfoRead(ctx, in, opts...)
}

// 获取设备信息详情
func (d *directDm) DeviceInfoRead(ctx context.Context, in *DeviceInfoReadReq, opts ...grpc.CallOption) (*DeviceInfo, error) {
	return d.svr.DeviceInfoRead(ctx, in)
}

// 新增设备
func (m *defaultDm) ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.ProductInfoCreate(ctx, in, opts...)
}

// 新增设备
func (d *directDm) ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ProductInfoCreate(ctx, in)
}

// 更新设备
func (m *defaultDm) ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.ProductInfoUpdate(ctx, in, opts...)
}

// 更新设备
func (d *directDm) ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ProductInfoUpdate(ctx, in)
}

// 删除设备
func (m *defaultDm) ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.ProductInfoDelete(ctx, in, opts...)
}

// 删除设备
func (d *directDm) ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ProductInfoDelete(ctx, in)
}

// 获取设备信息列表
func (m *defaultDm) ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.ProductInfoIndex(ctx, in, opts...)
}

// 获取设备信息列表
func (d *directDm) ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error) {
	return d.svr.ProductInfoIndex(ctx, in)
}

// 获取设备信息详情
func (m *defaultDm) ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.ProductInfoRead(ctx, in, opts...)
}

// 获取设备信息详情
func (d *directDm) ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	return d.svr.ProductInfoRead(ctx, in)
}

// 更新产品物模型
func (m *defaultDm) ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.ProductSchemaUpdate(ctx, in, opts...)
}

// 更新产品物模型
func (d *directDm) ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ProductSchemaUpdate(ctx, in)
}

// 获取产品物模型
func (m *defaultDm) ProductSchemaRead(ctx context.Context, in *ProductSchemaReadReq, opts ...grpc.CallOption) (*ProductSchema, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.ProductSchemaRead(ctx, in, opts...)
}

// 获取产品物模型
func (d *directDm) ProductSchemaRead(ctx context.Context, in *ProductSchemaReadReq, opts ...grpc.CallOption) (*ProductSchema, error) {
	return d.svr.ProductSchemaRead(ctx, in)
}

// 管理产品的固件
func (m *defaultDm) ManageFirmware(ctx context.Context, in *ManageFirmwareReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.ManageFirmware(ctx, in, opts...)
}

// 管理产品的固件
func (d *directDm) ManageFirmware(ctx context.Context, in *ManageFirmwareReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ManageFirmware(ctx, in)
}

// 获取产品固件信息
func (m *defaultDm) GetFirmwareInfo(ctx context.Context, in *GetFirmwareInfoReq, opts ...grpc.CallOption) (*GetFirmwareInfoResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.GetFirmwareInfo(ctx, in, opts...)
}

// 获取产品固件信息
func (d *directDm) GetFirmwareInfo(ctx context.Context, in *GetFirmwareInfoReq, opts ...grpc.CallOption) (*GetFirmwareInfoResp, error) {
	return d.svr.GetFirmwareInfo(ctx, in)
}

// 设备登录认证
func (m *defaultDm) LoginAuth(ctx context.Context, in *LoginAuthReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.LoginAuth(ctx, in, opts...)
}

// 设备登录认证
func (d *directDm) LoginAuth(ctx context.Context, in *LoginAuthReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.LoginAuth(ctx, in)
}

// 设备操作认证
func (m *defaultDm) AccessAuth(ctx context.Context, in *AccessAuthReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.AccessAuth(ctx, in, opts...)
}

// 设备操作认证
func (d *directDm) AccessAuth(ctx context.Context, in *AccessAuthReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.AccessAuth(ctx, in)
}

// 鉴定是否是root账号
func (m *defaultDm) RootCheck(ctx context.Context, in *RootCheckReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.RootCheck(ctx, in, opts...)
}

// 鉴定是否是root账号
func (d *directDm) RootCheck(ctx context.Context, in *RootCheckReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RootCheck(ctx, in)
}

// 同步调用设备行为
func (m *defaultDm) SendAction(ctx context.Context, in *SendActionReq, opts ...grpc.CallOption) (*SendActionResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.SendAction(ctx, in, opts...)
}

// 同步调用设备行为
func (d *directDm) SendAction(ctx context.Context, in *SendActionReq, opts ...grpc.CallOption) (*SendActionResp, error) {
	return d.svr.SendAction(ctx, in)
}

// 同步调用设备属性
func (m *defaultDm) SendProperty(ctx context.Context, in *SendPropertyReq, opts ...grpc.CallOption) (*SendPropertyResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.SendProperty(ctx, in, opts...)
}

// 同步调用设备属性
func (d *directDm) SendProperty(ctx context.Context, in *SendPropertyReq, opts ...grpc.CallOption) (*SendPropertyResp, error) {
	return d.svr.SendProperty(ctx, in)
}

// 获取设备sdk调试日志
func (m *defaultDm) DataSdkLogIndex(ctx context.Context, in *DataSdkLogIndexReq, opts ...grpc.CallOption) (*DataSdkLogIndexResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.DataSdkLogIndex(ctx, in, opts...)
}

// 获取设备sdk调试日志
func (d *directDm) DataSdkLogIndex(ctx context.Context, in *DataSdkLogIndexReq, opts ...grpc.CallOption) (*DataSdkLogIndexResp, error) {
	return d.svr.DataSdkLogIndex(ctx, in)
}

// 获取设备调试信息记录登入登出,操作
func (m *defaultDm) DataHubLogIndex(ctx context.Context, in *DataHubLogIndexReq, opts ...grpc.CallOption) (*DataHubLogIndexResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.DataHubLogIndex(ctx, in, opts...)
}

// 获取设备调试信息记录登入登出,操作
func (d *directDm) DataHubLogIndex(ctx context.Context, in *DataHubLogIndexReq, opts ...grpc.CallOption) (*DataHubLogIndexResp, error) {
	return d.svr.DataHubLogIndex(ctx, in)
}

// 获取设备数据信息
func (m *defaultDm) DataSchemaLatestIndex(ctx context.Context, in *DataSchemaLatestIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.DataSchemaLatestIndex(ctx, in, opts...)
}

// 获取设备数据信息
func (d *directDm) DataSchemaLatestIndex(ctx context.Context, in *DataSchemaLatestIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error) {
	return d.svr.DataSchemaLatestIndex(ctx, in)
}

// 获取设备数据信息
func (m *defaultDm) DataSchemaLogIndex(ctx context.Context, in *DataSchemaLogIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.DataSchemaLogIndex(ctx, in, opts...)
}

// 获取设备数据信息
func (d *directDm) DataSchemaLogIndex(ctx context.Context, in *DataSchemaLogIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error) {
	return d.svr.DataSchemaLogIndex(ctx, in)
}
