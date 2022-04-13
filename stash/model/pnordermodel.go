package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ PnOrderModel = (*customPnOrderModel)(nil)

type (
	// PnOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnOrderModel.
	PnOrderModel interface {
		pnOrderModel
		FindAllByOrderSn(ctx context.Context, orderSns []string) (*[]PnOrder, error)
	}

	customPnOrderModel struct {
		*defaultPnOrderModel
	}
)

// NewPnOrderModel returns a model for the database table.
func NewPnOrderModel(conn sqlx.SqlConn) PnOrderModel {
	return &customPnOrderModel{
		defaultPnOrderModel: newPnOrderModel(conn),
	}
}
func (m *customPnOrderModel) FindAllByOrderSn(ctx context.Context, orderSns []string) (*[]PnOrder, error) {
	query := fmt.Sprintf("select %s from %s where `order_sn` in (?)  ", pnOrderRows, m.table)
	var resp []PnOrder
	err := m.conn.QueryRowsCtx(ctx, &resp, query, strings.Join(orderSns, ","))
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}