package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PnScenicOrderModel = (*customPnScenicOrderModel)(nil)

type (
	// PnScenicOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnScenicOrderModel.
	PnScenicOrderModel interface {
		pnScenicOrderModel
	}

	customPnScenicOrderModel struct {
		*defaultPnScenicOrderModel
	}
)

// NewPnScenicOrderModel returns a model for the database table.
func NewPnScenicOrderModel(conn sqlx.SqlConn) PnScenicOrderModel {
	return &customPnScenicOrderModel{
		defaultPnScenicOrderModel: newPnScenicOrderModel(conn),
	}
}
