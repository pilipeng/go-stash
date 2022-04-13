package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PnKkOrderModel = (*customPnKkOrderModel)(nil)

type (
	// PnKkOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnKkOrderModel.
	PnKkOrderModel interface {
		pnKkOrderModel
	}

	customPnKkOrderModel struct {
		*defaultPnKkOrderModel
	}
)

// NewPnKkOrderModel returns a model for the database table.
func NewPnKkOrderModel(conn sqlx.SqlConn) PnKkOrderModel {
	return &customPnKkOrderModel{
		defaultPnKkOrderModel: newPnKkOrderModel(conn),
	}
}

