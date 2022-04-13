package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PnCakeFlavorModel = (*customPnCakeFlavorModel)(nil)

type (
	// PnCakeFlavorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnCakeFlavorModel.
	PnCakeFlavorModel interface {
		pnCakeFlavorModel
	}

	customPnCakeFlavorModel struct {
		*defaultPnCakeFlavorModel
	}
)

// NewPnCakeFlavorModel returns a model for the database table.
func NewPnCakeFlavorModel(conn sqlx.SqlConn) PnCakeFlavorModel {
	return &customPnCakeFlavorModel{
		defaultPnCakeFlavorModel: newPnCakeFlavorModel(conn),
	}
}
