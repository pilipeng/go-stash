package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PnMOrderModel = (*customPnMOrderModel)(nil)

type (
	// PnMOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPnMOrderModel.
	PnMOrderModel interface {
		pnMOrderModel
	}

	customPnMOrderModel struct {
		*defaultPnMOrderModel
	}

	OrderSnapshot []struct {
		SupplierID        int64   `json:"-"`
		Supplier          string  `json:"-"`
		SuppliersNameReal string  `json:"-"`
		GoodsID           int64   `json:"-"`
		SkuID             int64   `json:"-"`
		ZySn              string  `json:"-"`
		SkuSn             string  `json:"-"`
		FinanceSn         string  `json:"-"`
		BrandID           int64   `json:"-"`
		CateID            int64   `json:"-"`
		FirstCateID       int64   `json:"-"`
		SecondCateID      int64   `json:"-"`
		ThirdCateID       int64   `json:"-"`
		GoodsName         string  `json:"goods_name,omitempty"`
		Count             int64   `json:"-"`
		ZtPrice           float64 `json:"-"`
		BasePrice         string  `json:"-"`
		Price             float64 `json:"-"`
		GoodsImg          string  `json:"-"`
		Skuattr           string  `json:"-"`
		Subgoods          string  `json:"-"`
		CakeTime          string  `json:"-"`
		IsVirtual         int8    `json:"-"`
		Gifts             string  `json:"-"`
	}
)

// NewPnMOrderModel returns a model for the database table.
func NewPnMOrderModel(conn sqlx.SqlConn) PnMOrderModel {
	return &customPnMOrderModel{
		defaultPnMOrderModel: newPnMOrderModel(conn),
	}
}
