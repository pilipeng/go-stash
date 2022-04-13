package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PnDamaiShowModel = (*customPnDamaiShowModel)(nil)

type (
	// PnDamaiShowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnDamaiShowModel.
	PnDamaiShowModel interface {
		pnDamaiShowModel
	}

	customPnDamaiShowModel struct {
		*defaultPnDamaiShowModel
	}
)

// NewPnDamaiShowModel returns a model for the database table.
func NewPnDamaiShowModel(conn sqlx.SqlConn) PnDamaiShowModel {
	return &customPnDamaiShowModel{
		defaultPnDamaiShowModel: newPnDamaiShowModel(conn),
	}
}
