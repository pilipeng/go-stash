package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PnHotWordModel = (*customPnHotWordModel)(nil)

type (
	// PnHotWordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnHotWordModel.
	PnHotWordModel interface {
		pnHotWordModel
	}

	customPnHotWordModel struct {
		*defaultPnHotWordModel
	}
)

// NewPnHotWordModel returns a model for the database table.
func NewPnHotWordModel(conn sqlx.SqlConn) PnHotWordModel {
	return &customPnHotWordModel{
		defaultPnHotWordModel: newPnHotWordModel(conn),
	}
}
