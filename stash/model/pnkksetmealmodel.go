package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PnKkSetmealModel = (*customPnKkSetmealModel)(nil)

type (
	// PnKkSetmealModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnKkSetmealModel.
	PnKkSetmealModel interface {
		pnKkSetmealModel
	}

	customPnKkSetmealModel struct {
		*defaultPnKkSetmealModel
	}
)

// NewPnKkSetmealModel returns a model for the database table.
func NewPnKkSetmealModel(conn sqlx.SqlConn) PnKkSetmealModel {
	return &customPnKkSetmealModel{
		defaultPnKkSetmealModel: newPnKkSetmealModel(conn),
	}
}
