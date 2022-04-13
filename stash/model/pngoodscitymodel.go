package model

import (
	"context"
	"fmt"
	"github.com/kevwan/go-stash/stash/lib"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PnGoodsCityModel = (*customPnGoodsCityModel)(nil)

type (
	// PnGoodsCityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnGoodsCityModel.
	PnGoodsCityModel interface {
		pnGoodsCityModel
		FindAllBySpuIds(ctx context.Context,spuIds []int64) (*[]PnGoodsCity, error)
	}

	customPnGoodsCityModel struct {
		*defaultPnGoodsCityModel
	}
)

// NewPnGoodsCityModel returns a model for the database table.
func NewPnGoodsCityModel(conn sqlx.SqlConn) PnGoodsCityModel {
	return &customPnGoodsCityModel{
		defaultPnGoodsCityModel: newPnGoodsCityModel(conn),
	}
}


func (m *customPnGoodsCityModel) FindAllBySpuIds(ctx context.Context,spuIds []int64) (*[]PnGoodsCity, error) {
	query := fmt.Sprintf("select %s from %s where `spu_id` in (?) ", pnGoodsCityRows, m.table)
	var resp []PnGoodsCity
	err := m.conn.QueryRows(&resp, query, lib.ArrayOrSlice2String(spuIds) )
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}