package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PnMCinemaModel = (*customPnMCinemaModel)(nil)

type (
	// PnMCinemaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnMCinemaModel.
	PnMCinemaModel interface {
		pnMCinemaModel
	}

	customPnMCinemaModel struct {
		*defaultPnMCinemaModel
	}
)

// NewPnMCinemaModel returns a model for the database table.
func NewPnMCinemaModel(conn sqlx.SqlConn) PnMCinemaModel {
	return &customPnMCinemaModel{
		defaultPnMCinemaModel: newPnMCinemaModel(conn),
	}
}
