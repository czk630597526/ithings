package menulogic

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/src/syssvr/internal/repo/mysql"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuIndexLogic {
	return &MenuIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuIndexLogic) MenuIndex(in *sys.MenuIndexReq) (*sys.MenuIndexResp, error) {
	info := make([]*sys.MenuData, 0)
	if in.Role != 0 {
		//获取角色关联的菜单列表
		menuIds, err := l.svcCtx.RoleModel.IndexRoleIDMenuID(in.Role)
		if err != nil {
			return nil, errors.Database.AddDetail(err)
		}
		menuInfos, err := l.svcCtx.MenuModel.Index(l.ctx, &mysql.MenuIndexFilter{MenuIds: menuIds})
		if err != nil {
			l.Errorf("MenuIndex find menu_info err,menuIds:%v,err:%v", menuIds, err)
			return nil, errors.Database.AddDetail(err)
		}
		for _, v := range menuInfos {
			info = append(info, MenuInfoToPb(v))
		}
	} else {
		//获取完整菜单列表
		mes, err := l.svcCtx.MenuModel.Index(l.ctx, &mysql.MenuIndexFilter{
			Role: in.Role,
			Name: in.Name,
			Path: in.Path,
		})
		if err != nil {
			return nil, errors.Database.AddDetail(err)
		}
		for _, me := range mes {
			info = append(info, MenuInfoToPb(me))
		}
	}

	return &sys.MenuIndexResp{
		List: info,
	}, nil
}
