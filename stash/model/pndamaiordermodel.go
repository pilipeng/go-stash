package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PnDamaiOrderModel = (*customPnDamaiOrderModel)(nil)

type (
	// PnDamaiOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnDamaiOrderModel.
	PnDamaiOrderModel interface {
		pnDamaiOrderModel
	}

	customPnDamaiOrderModel struct {
		*defaultPnDamaiOrderModel
	}
)

// NewPnDamaiOrderModel returns a model for the database table.
func NewPnDamaiOrderModel(conn sqlx.SqlConn) PnDamaiOrderModel {
	return &customPnDamaiOrderModel{
		defaultPnDamaiOrderModel: newPnDamaiOrderModel(conn),
	}
}
