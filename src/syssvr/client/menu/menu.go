// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package client

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckTokenReq     = sys.CheckTokenReq
	CheckTokenResp    = sys.CheckTokenResp
	ConfigResp        = sys.ConfigResp
	DateRange         = sys.DateRange
	JwtToken          = sys.JwtToken
	LoginLogCreateReq = sys.LoginLogCreateReq
	LoginLogIndexData = sys.LoginLogIndexData
	LoginLogIndexReq  = sys.LoginLogIndexReq
	LoginLogIndexResp = sys.LoginLogIndexResp
	LoginReq          = sys.LoginReq
	LoginResp         = sys.LoginResp
	Map               = sys.Map
	MenuCreateReq     = sys.MenuCreateReq
	MenuData          = sys.MenuData
	MenuDeleteReq     = sys.MenuDeleteReq
	MenuIndexReq      = sys.MenuIndexReq
	MenuIndexResp     = sys.MenuIndexResp
	MenuUpdateReq     = sys.MenuUpdateReq
	OperLogCreateReq  = sys.OperLogCreateReq
	OperLogIndexData  = sys.OperLogIndexData
	OperLogIndexReq   = sys.OperLogIndexReq
	OperLogIndexResp  = sys.OperLogIndexResp
	PageInfo          = sys.PageInfo
	Response          = sys.Response
	RoleCreateReq     = sys.RoleCreateReq
	RoleDeleteReq     = sys.RoleDeleteReq
	RoleIndexData     = sys.RoleIndexData
	RoleIndexReq      = sys.RoleIndexReq
	RoleIndexResp     = sys.RoleIndexResp
	RoleMenuUpdateReq = sys.RoleMenuUpdateReq
	RoleUpdateReq     = sys.RoleUpdateReq
	UserCreateReq     = sys.UserCreateReq
	UserCreateResp    = sys.UserCreateResp
	UserDeleteReq     = sys.UserDeleteReq
	UserIndexReq      = sys.UserIndexReq
	UserIndexResp     = sys.UserIndexResp
	UserInfo          = sys.UserInfo
	UserReadReq       = sys.UserReadReq
	UserReadResp      = sys.UserReadResp
	UserUpdateReq     = sys.UserUpdateReq

	Menu interface {
		MenuCreate(ctx context.Context, in *MenuCreateReq, opts ...grpc.CallOption) (*Response, error)
		MenuIndex(ctx context.Context, in *MenuIndexReq, opts ...grpc.CallOption) (*MenuIndexResp, error)
		MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*Response, error)
		MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*Response, error)
	}

	defaultMenu struct {
		cli zrpc.Client
	}

	directMenu struct {
		svcCtx *svc.ServiceContext
		svr    sys.MenuServer
	}
)

func NewMenu(cli zrpc.Client) Menu {
	return &defaultMenu{
		cli: cli,
	}
}

func NewDirectMenu(svcCtx *svc.ServiceContext, svr sys.MenuServer) Menu {
	return &directMenu{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultMenu) MenuCreate(ctx context.Context, in *MenuCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewMenuClient(m.cli.Conn())
	return client.MenuCreate(ctx, in, opts...)
}

func (d *directMenu) MenuCreate(ctx context.Context, in *MenuCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.MenuCreate(ctx, in)
}

func (m *defaultMenu) MenuIndex(ctx context.Context, in *MenuIndexReq, opts ...grpc.CallOption) (*MenuIndexResp, error) {
	client := sys.NewMenuClient(m.cli.Conn())
	return client.MenuIndex(ctx, in, opts...)
}

func (d *directMenu) MenuIndex(ctx context.Context, in *MenuIndexReq, opts ...grpc.CallOption) (*MenuIndexResp, error) {
	return d.svr.MenuIndex(ctx, in)
}

func (m *defaultMenu) MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewMenuClient(m.cli.Conn())
	return client.MenuUpdate(ctx, in, opts...)
}

func (d *directMenu) MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.MenuUpdate(ctx, in)
}

func (m *defaultMenu) MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewMenuClient(m.cli.Conn())
	return client.MenuDelete(ctx, in, opts...)
}

func (d *directMenu) MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.MenuDelete(ctx, in)
}
