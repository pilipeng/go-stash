// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/kevwan/go-stash/stash/lib"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
)

var (
	pnOrderFieldNames          = builder.RawFieldNames(&PnOrder{})
	pnOrderRows                = strings.Join(pnOrderFieldNames, ",")
	pnOrderRowsExpectAutoSet   = strings.Join(stringx.Remove(pnOrderFieldNames, "`order_id`", "`create_time`", "`update_time`"), ",")
	pnOrderRowsWithPlaceHolder = strings.Join(stringx.Remove(pnOrderFieldNames, "`order_id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	pnOrderModel interface {
		Insert(ctx context.Context, data *PnOrder) (sql.Result, error)
		FindOne(ctx context.Context, orderId int64) (*PnOrder, error)
		Update(ctx context.Context, data *PnOrder) error
		Delete(ctx context.Context, orderId int64) error
	}

	defaultPnOrderModel struct {
		conn  sqlx.SqlConn
		table string
	}
	PnOrder struct {
		OrderId             int64          `db:"order_id" json:"order_id,string"` // 主键，自增
		UserId              int64          `db:"user_id" json:"user_id,string"`   // 用户id
		UserName            string         `db:"user_name" json:"user_name"`
		OrderSn             string         `db:"order_sn" json:"order_sn"`                          // 订单号
		ParentOrderSn       string         `db:"parent_order_sn" json:"parent_order_sn"`            // 父订单号，订单为子订单时该字段有值
		OrderType           int64          `db:"order_type" json:"order_type,string"`               // 0：父订单 1：子订单
		ZtOrderId           string         `db:"zt_order_id" json:"zt_order_id"`                    // 中台订单号
		OrderStatus         int64          `db:"order_status" json:"order_status,string"`           //  0：待付款 1：已付款 2：已发货 3：已取消 4:退款中 5：已退款 6:已完成 7:支付失败
		PayStatus           int64          `db:"pay_status" json:"pay_status,string"`               // 付款状态 0：待付款 1：已付款 2：未付款 3:支付失败 4:未补差价
		DeliverStatus       int64          `db:"deliver_status" json:"deliver_status,string"`       // 发货状态 0：未发货 1：已发货 2：已退货
		PayType             int64          `db:"pay_type" json:"pay_type,string"`                   // 支付方式  0 微信 1支付宝
		PostFee             float64        `db:"post_fee" json:"post_fee,string"`                   // 邮费
		TotalPrice          float64        `db:"total_price" json:"total_price,string"`             // 商品总价
		TotalBasePrice      float64        `db:"total_base_price" json:"total_base_price,string"`   // 总成本
		RefundBasePrice     float64        `db:"refund_base_price" json:"refund_base_price,string"` // 已退成本
		PayPrice            float64        `db:"pay_price" json:"pay_price,string"`                 // 应付金额(商品总价+邮费)
		CardNo              string         `db:"card_no" json:"card_no"`                            // 支付卡号
		CardType            int64          `db:"card_type" json:"card_type,string"`                 // 卡类型  1：  红卡 2：蓝卡 3：绿卡
		CardTypeName        string         `db:"card_type_name" json:"card_type_name"`              // 卡类型名称
		CardPayPrice        float64        `db:"card_pay_price" json:"card_pay_price,string"`       // 卡券支付金额
		CardPayOneCard      string         `db:"card_pay_one_card" json:"card_pay_one_card"`        // 次卡支付金额
		CardResponse        lib.JsonNullString `db:"card_response" json:"card_response,omitempty"`      // 卡消费返回信息
		Remark              string         `db:"remark" json:"remark,omitempty"`                    // 附言
		OrderTime           int64          `db:"order_time" json:"order_time,string"`               // 订单时间
		PayTime             int64          `db:"pay_time" json:"pay_time,string"`                   // 支付时间
		Receiver            string         `db:"receiver" json:"receiver"`                          // 收货人
		ReceiverPhone       string         `db:"receiver_phone" json:"receiver_phone"`              // 收货人手机号
		OrderSnapshot       string    `db:"order_snapshot" json:"order_snapshot"`              // 商品快照
		OrderExtendSnapshot string         `db:"order_extend_snapshot" json:"-"`                    // 订单其他信息快照（商品分组信息、商品价、邮费、总计等）
		ReceiverSnapshot    string         `db:"receiver_snapshot" json:"receiver_snapshot"`        // 收货信息快照
		ZtResponse          string         `db:"zt_response" json:"zt_response"`                    // 中台接口响应信息
		PayLimitTime        int64          `db:"pay_limit_time" json:"pay_limit_time,string"`       // 未支付过期时间
		WxResponse          string         `db:"wx_response" json:"wx_response"`                    // 微信回调结果
		Express             string         `db:"express"  json:"express"`                           // 快递名称
		ExpressNo           string         `db:"express_no" json:"express_no"`                      // 快递单号
		IsShow              int64          `db:"is_show" json:"is_show,string"`                     // 订单是否显示 (拆单时，父订单此字段为0) 0：否 1：是
		SupplierId          int64          `db:"supplier_id" json:"supplier_id,string"`             // 供应商id
		Supplier            string         `db:"supplier"  json:"supplier"`                         // 供应商
		IsGift              int64          `db:"is_gift" json:"is_gift,string"`                     // 是否礼包订单 0：否 1：是
		CompanyUserId       int64          `db:"company_user_id" json:"company_user_id,string"`     // 企业专区用户id 0:  商城 1：企业专区 2：次卡专区
		CompanydeGoodsName  string         `db:"companyde_goods_name" json:"companyde_goods_name"`  // 次卡专区礼包名
		SkusName            string         `db:"skus_name" json:"skus_name"`                        // 订单商品名
		Platform            int64          `db:"platform" json:"platform,string"`                   // 订单平台  1：PC   2：手机
		CompanyGoodsId      int64          `db:"company_goods_id" json:"company_goods_id,string"`   // 企业专区套餐id
		UpdateTime          lib.Datetime    `db:"update_time" json:"update_time,string"`                     // 修改时间
		//UpdateTime          time.Time `db:"update_time" json:"update_time"`                            // 修改时间
		IsHide              int64   `db:"is_hide" json:"is_hide,string"`                             // 是否被删除
		PriceDifference     float64 `db:"price_difference" json:"price_difference,string"`           // 微信补差金额（废弃）
		PriceDifferenceBack float64 `db:"price_difference_back" json:"price_difference_back,string"` // 微信补差金额备份（废弃）
		OrderMaizuoSn       string  `db:"orderMaizuoSn" json:"orderMaizuoSn"`                      // 卖座订单号
		SurplusMoney        float64 `db:"surplusMoney" json:"surplusMoney,string"`                  // 微信补差金额
		//PaySerialNo         string  `db:"pay_serial_no" json:"pay_serial_no"`                        // 第三方支付流水号
	}

	/*PnOrder struct {
		OrderId             int64          `db:"order_id"` // 主键，自增
		UserId              int64          `db:"user_id"`  // 用户id
		UserName            string         `db:"user_name"`
		OrderSn             string         `db:"order_sn"`              // 订单号
		ParentOrderSn       string         `db:"parent_order_sn"`       // 父订单号，订单为子订单时该字段有值
		OrderType           int64          `db:"order_type"`            // 0：父订单 1：子订单
		ZtOrderId           string         `db:"zt_order_id"`           // 中台订单号
		OrderStatus         int64          `db:"order_status"`          //  0：待付款 1：已付款 2：已发货 3：已取消 4:退款中 5：已退款 6:已完成 7:支付失败
		PayStatus           int64          `db:"pay_status"`            // 付款状态 0：待付款 1：已付款 2：未付款 3:支付失败 4:未补差价
		DeliverStatus       int64          `db:"deliver_status"`        // 发货状态 0：未发货 1：已发货 2：已退货
		PayType             int64          `db:"pay_type"`              // 支付方式  0 微信 1支付宝
		PostFee             float64        `db:"post_fee"`              // 邮费
		TotalPrice          float64        `db:"total_price"`           // 商品总价
		TotalBasePrice      float64        `db:"total_base_price"`      // 总成本
		RefundBasePrice     float64        `db:"refund_base_price"`     // 已退成本
		PayPrice            float64        `db:"pay_price"`             // 应付金额(商品总价+邮费)
		CardNo              string         `db:"card_no"`               // 支付卡号
		CardType            int64          `db:"card_type"`             // 卡类型  1：  红卡 2：蓝卡 3：绿卡
		CardTypeName        string         `db:"card_type_name"`        // 卡类型名称
		CardPayPrice        float64        `db:"card_pay_price"`        // 卡券支付金额
		CardPayOneCard      string         `db:"card_pay_one_card"`     // 次卡支付金额
		CardResponse        lib.JsonNullString `db:"card_response"`         // 卡消费返回信息
		Remark              string         `db:"remark"`                // 附言
		OrderTime           int64          `db:"order_time"`            // 订单时间
		PayTime             int64          `db:"pay_time"`              // 支付时间
		Receiver            string         `db:"receiver"`              // 收货人
		ReceiverPhone       string         `db:"receiver_phone"`        // 收货人手机号
		OrderSnapshot       string         `db:"order_snapshot"`        // 商品快照
		OrderExtendSnapshot string         `db:"order_extend_snapshot"` // 订单其他信息快照（商品分组信息、商品价、邮费、总计等）
		ReceiverSnapshot    string         `db:"receiver_snapshot"`     // 收货信息快照
		ZtResponse          string         `db:"zt_response"`           // 中台接口响应信息
		PayLimitTime        int64          `db:"pay_limit_time"`        // 未支付过期时间
		WxResponse          string         `db:"wx_response"`           // 微信回调结果
		Express             string         `db:"express"`               // 快递名称
		ExpressNo           string         `db:"express_no"`            // 快递单号
		IsShow              int64          `db:"is_show"`               // 订单是否显示 (拆单时，父订单此字段为0) 0：否 1：是
		SupplierId          int64          `db:"supplier_id"`           // 供应商id
		Supplier            string         `db:"supplier"`              // 供应商
		IsGift              int64          `db:"is_gift"`               // 是否礼包订单 0：否 1：是
		CompanyUserId       int64          `db:"company_user_id"`       // 企业专区用户id 0:  商城 1：企业专区 2：次卡专区
		CompanydeGoodsName  string         `db:"companyde_goods_name"`  // 次卡专区礼包名
		SkusName            string         `db:"skus_name"`             // 订单商品名
		Platform            int64          `db:"platform" json:"platform,string"`              // 订单平台  1：PC   2：手机
		CompanyGoodsId      int64          `db:"company_goods_id"`      // 企业专区套餐id
		UpdateTime          lib.Datetime      `db:"update_time"`           // 修改时间
		IsHide              int64          `db:"is_hide"`               // 是否被删除
		PriceDifference     float64        `db:"price_difference"`      // 微信补差金额（废弃）
		PriceDifferenceBack float64        `db:"price_difference_back"` // 微信补差金额备份（废弃）
		OrderMaizuoSn       string         `db:"orderMaizuoSn"`         // 卖座订单号
		SurplusMoney        float64        `db:"surplusMoney"`          // 微信补差金额
		PaySerialNo         string         `db:"pay_serial_no"`         // 第三方支付流水号
	}*/
)

func newPnOrderModel(conn sqlx.SqlConn) *defaultPnOrderModel {
	return &defaultPnOrderModel{
		conn:  conn,
		table: "`pn_order`",
	}
}

func (m *defaultPnOrderModel) Insert(ctx context.Context, data *PnOrder) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pnOrderRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.UserName, data.OrderSn, data.ParentOrderSn, data.OrderType, data.ZtOrderId, data.OrderStatus, data.PayStatus, data.DeliverStatus, data.PayType, data.PostFee, data.TotalPrice, data.TotalBasePrice, data.RefundBasePrice, data.PayPrice, data.CardNo, data.CardType, data.CardTypeName, data.CardPayPrice, data.CardPayOneCard, data.CardResponse, data.Remark, data.OrderTime, data.PayTime, data.Receiver, data.ReceiverPhone, data.OrderSnapshot, data.OrderExtendSnapshot, data.ReceiverSnapshot, data.ZtResponse, data.PayLimitTime, data.WxResponse, data.Express, data.ExpressNo, data.IsShow, data.SupplierId, data.Supplier, data.IsGift, data.CompanyUserId, data.CompanydeGoodsName, data.SkusName, data.Platform, data.CompanyGoodsId, data.IsHide, data.PriceDifference, data.PriceDifferenceBack, data.OrderMaizuoSn, data.SurplusMoney)
	return ret, err
}

func (m *defaultPnOrderModel) FindOne(ctx context.Context, orderId int64) (*PnOrder, error) {
	query := fmt.Sprintf("select %s from %s where `order_id` = ? limit 1", pnOrderRows, m.table)
	var resp PnOrder
	err := m.conn.QueryRowCtx(ctx, &resp, query, orderId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPnOrderModel) Update(ctx context.Context, data *PnOrder) error {
	query := fmt.Sprintf("update %s set %s where `order_id` = ?", m.table, pnOrderRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.UserName, data.OrderSn, data.ParentOrderSn, data.OrderType, data.ZtOrderId, data.OrderStatus, data.PayStatus, data.DeliverStatus, data.PayType, data.PostFee, data.TotalPrice, data.TotalBasePrice, data.RefundBasePrice, data.PayPrice, data.CardNo, data.CardType, data.CardTypeName, data.CardPayPrice, data.CardPayOneCard, data.CardResponse, data.Remark, data.OrderTime, data.PayTime, data.Receiver, data.ReceiverPhone, data.OrderSnapshot, data.OrderExtendSnapshot, data.ReceiverSnapshot, data.ZtResponse, data.PayLimitTime, data.WxResponse, data.Express, data.ExpressNo, data.IsShow, data.SupplierId, data.Supplier, data.IsGift, data.CompanyUserId, data.CompanydeGoodsName, data.SkusName, data.Platform, data.CompanyGoodsId, data.IsHide, data.PriceDifference, data.PriceDifferenceBack, data.OrderMaizuoSn, data.SurplusMoney, data.OrderId)
	return err
}

func (m *defaultPnOrderModel) Delete(ctx context.Context, orderId int64) error {
	query := fmt.Sprintf("delete from %s where `order_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, orderId)
	return err
}

func (m *defaultPnOrderModel) tableName() string {
	return m.table
}
