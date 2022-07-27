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
	productSchemaFieldNames          = builder.RawFieldNames(&ProductSchema{})
	productSchemaRows                = strings.Join(productSchemaFieldNames, ",")
	productSchemaRowsExpectAutoSet   = strings.Join(stringx.Remove(productSchemaFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	productSchemaRowsWithPlaceHolder = strings.Join(stringx.Remove(productSchemaFieldNames, "`productID`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	productSchemaModel interface {
		Insert(ctx context.Context, data *ProductSchema) (sql.Result, error)
		FindOne(ctx context.Context, productID string) (*ProductSchema, error)
		Update(ctx context.Context, newData *ProductSchema) error
		Delete(ctx context.Context, productID string) error
	}

	defaultProductSchemaModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ProductSchema struct {
		ProductID   string       `db:"productID"` // 产品id
		Schema      string       `db:"schema"`    // 物模型模板
		CreatedTime time.Time    `db:"createdTime"`
		UpdatedTime sql.NullTime `db:"updatedTime"`
		DeletedTime sql.NullTime `db:"deletedTime"`
	}
)

func newProductSchemaModel(conn sqlx.SqlConn) *defaultProductSchemaModel {
	return &defaultProductSchemaModel{
		conn:  conn,
		table: "`product_schema`",
	}
}

func (m *defaultProductSchemaModel) Delete(ctx context.Context, productID string) error {
	query := fmt.Sprintf("delete from %s where `productID` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, productID)
	return err
}

func (m *defaultProductSchemaModel) FindOne(ctx context.Context, productID string) (*ProductSchema, error) {
	query := fmt.Sprintf("select %s from %s where `productID` = ? limit 1", productSchemaRows, m.table)
	var resp ProductSchema
	err := m.conn.QueryRowCtx(ctx, &resp, query, productID)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductSchemaModel) Insert(ctx context.Context, data *ProductSchema) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, productSchemaRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProductID, data.Schema, data.CreatedTime, data.UpdatedTime, data.DeletedTime)
	return ret, err
}

func (m *defaultProductSchemaModel) Update(ctx context.Context, data *ProductSchema) error {
	query := fmt.Sprintf("update %s set %s where `productID` = ?", m.table, productSchemaRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Schema, data.CreatedTime, data.UpdatedTime, data.DeletedTime, data.ProductID)
	return err
}

func (m *defaultProductSchemaModel) tableName() string {
	return m.table
}
