package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PnMFilmsModel = (*customPnMFilmsModel)(nil)

type (
	// PnMFilmsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnMFilmsModel.
	PnMFilmsModel interface {
		pnMFilmsModel
	}

	customPnMFilmsModel struct {
		*defaultPnMFilmsModel
	}
)

// NewPnMFilmsModel returns a model for the database table.
func NewPnMFilmsModel(conn sqlx.SqlConn) PnMFilmsModel {
	return &customPnMFilmsModel{
		defaultPnMFilmsModel: newPnMFilmsModel(conn),
	}
}
