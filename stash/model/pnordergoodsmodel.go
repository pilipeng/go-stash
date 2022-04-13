package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PnOrderGoodsModel = (*customPnOrderGoodsModel)(nil)

type (
	// PnOrderGoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnOrderGoodsModel.
	PnOrderGoodsModel interface {
		pnOrderGoodsModel
		FindAll(ctx context.Context,orderSn string) (*[]PnOrderGoods, error)
	}

	customPnOrderGoodsModel struct {
		*defaultPnOrderGoodsModel
	}
)

// NewPnOrderGoodsModel returns a model for the database table.
func NewPnOrderGoodsModel(conn sqlx.SqlConn) PnOrderGoodsModel {
	return &customPnOrderGoodsModel{
		defaultPnOrderGoodsModel: newPnOrderGoodsModel(conn),
	}
}

func (m *customPnOrderGoodsModel) FindAll(ctx context.Context,orderSn string) (*[]PnOrderGoods, error) {
	query := fmt.Sprintf("select %s from %s where `order_sn` = ? ", pnOrderGoodsRows, m.table)
	var resp []PnOrderGoods
	err := m.conn.QueryRowsCtx(ctx,&resp, query, orderSn)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}