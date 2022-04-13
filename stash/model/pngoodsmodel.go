package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PnGoodsModel = (*customPnGoodsModel)(nil)

type (
	// PnGoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnGoodsModel.
	PnGoodsModel interface {
		pnGoodsModel
		FindAllBySpuId(ctx context.Context, spuId int64) (*[]PnGoods, error)
	}

	customPnGoodsModel struct {
		*defaultPnGoodsModel
	}
)

// NewPnGoodsModel returns a model for the database table.
func NewPnGoodsModel(conn sqlx.SqlConn) PnGoodsModel {
	return &customPnGoodsModel{
		defaultPnGoodsModel: newPnGoodsModel(conn),
	}
}

func (m *customPnGoodsModel) FindAllBySpuId(ctx context.Context,spuId int64) (*[]PnGoods, error) {
	query := fmt.Sprintf("select %s from %s where `spu_id` = ? limit 1", pnGoodsRows, m.table)
	var resp []PnGoods
	err := m.conn.QueryRowsCtx(ctx,&resp, query, spuId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}