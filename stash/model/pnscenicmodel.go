package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PnScenicModel = (*customPnScenicModel)(nil)

type (
	// PnScenicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnScenicModel.
	PnScenicModel interface {
		pnScenicModel
	}

	customPnScenicModel struct {
		*defaultPnScenicModel
	}
)

// NewPnScenicModel returns a model for the database table.
func NewPnScenicModel(conn sqlx.SqlConn) PnScenicModel {
	return &customPnScenicModel{
		defaultPnScenicModel: newPnScenicModel(conn),
	}
}
