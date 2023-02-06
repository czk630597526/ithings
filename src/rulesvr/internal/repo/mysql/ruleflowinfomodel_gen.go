// Code generated by goctl. DO NOT EDIT!

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	ruleFlowInfoFieldNames          = builder.RawFieldNames(&RuleFlowInfo{})
	ruleFlowInfoRows                = strings.Join(ruleFlowInfoFieldNames, ",")
	ruleFlowInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(ruleFlowInfoFieldNames, "`id`", "`updatedTime`", "`deletedTime`", "`createdTime`"), ",")
	ruleFlowInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(ruleFlowInfoFieldNames, "`id`", "`updatedTime`", "`deletedTime`", "`createdTime`"), "=?,") + "=?"
)

type (
	ruleFlowInfoModel interface {
		Insert(ctx context.Context, data *RuleFlowInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RuleFlowInfo, error)
		FindOneByName(ctx context.Context, name string) (*RuleFlowInfo, error)
		Update(ctx context.Context, data *RuleFlowInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRuleFlowInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RuleFlowInfo struct {
		Id          int64        `db:"id"`          // id
		Name        string       `db:"name"`        // 流的名称
		Password    string       `db:"password"`    // 登录密码
		Desc        string       `db:"desc"`        // 描述
		IsDisabled  int64        `db:"isDisabled"`  // 是否禁用 1:是 2:否
		CreatedTime time.Time    `db:"createdTime"` // 创建时间
		UpdatedTime time.Time    `db:"updatedTime"` // 更新时间
		DeletedTime sql.NullTime `db:"deletedTime"` // 删除时间，默认为空，表示未删除，非空表示已删除
	}
)

func newRuleFlowInfoModel(conn sqlx.SqlConn) *defaultRuleFlowInfoModel {
	return &defaultRuleFlowInfoModel{
		conn:  conn,
		table: "`rule_flow_info`",
	}
}

func (m *defaultRuleFlowInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRuleFlowInfoModel) FindOne(ctx context.Context, id int64) (*RuleFlowInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ruleFlowInfoRows, m.table)
	var resp RuleFlowInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRuleFlowInfoModel) FindOneByName(ctx context.Context, name string) (*RuleFlowInfo, error) {
	var resp RuleFlowInfo
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", ruleFlowInfoRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRuleFlowInfoModel) Insert(ctx context.Context, data *RuleFlowInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, ruleFlowInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Password, data.Desc, data.IsDisabled)
	return ret, err
}

func (m *defaultRuleFlowInfoModel) Update(ctx context.Context, newData *RuleFlowInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ruleFlowInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Name, newData.Password, newData.Desc, newData.IsDisabled, newData.Id)
	return err
}

func (m *defaultRuleFlowInfoModel) tableName() string {
	return m.table
}