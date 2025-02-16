// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/kevwan/go-stash/stash/lib"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	pnDamaiOrderFieldNames          = builder.RawFieldNames(&PnDamaiOrder{})
	pnDamaiOrderRows                = strings.Join(pnDamaiOrderFieldNames, ",")
	pnDamaiOrderRowsExpectAutoSet   = strings.Join(stringx.Remove(pnDamaiOrderFieldNames, "`order_id`", "`create_time`", "`update_time`"), ",")
	pnDamaiOrderRowsWithPlaceHolder = strings.Join(stringx.Remove(pnDamaiOrderFieldNames, "`order_id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	pnDamaiOrderModel interface {
		Insert(ctx context.Context, data *PnDamaiOrder) (sql.Result, error)
		FindOne(ctx context.Context, orderId int64) (*PnDamaiOrder, error)
		Update(ctx context.Context, data *PnDamaiOrder) error
		Delete(ctx context.Context, orderId int64) error
	}

	defaultPnDamaiOrderModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PnDamaiOrder struct {
		OrderId          int64          `db:"order_id" json:"order_id,string"` // 主键，自增
		UserId           int64          `db:"user_id" json:"user_id,string"`
		UserName         string         `db:"user_name" json:"user_name"`          // 用户名（手机号）
		OrderSn          string         `db:"order_sn" json:"order_sn"`           // 订单号
		OrderStatus      int64          `db:"order_status" json:"order_status,string"`       //  0：待付款 1：已付款 2：已发货 3：已取消 4:未补差 5：已退款 6:已完成
		DeliverStatus    int64          `db:"deliver_status" json:"deliver_status,string"`     // 发货状态 0：未发货 1：已发货 2：已退货
		PostFee          float64        `db:"post_fee" json:"post_fee,string"`           // 邮费
		TotalPrice       float64        `db:"total_price" json:"total_price,string"`        // 商品总价
		TotalBasePrice   float64        `db:"total_base_price" json:"total_base_price,string"`   // 订单商品成本总价
		PayPrice         float64        `db:"pay_price" json:"pay_price,string"`          // 应付金额(商品总价+邮费)
		CardNo           string         `db:"card_no" json:"card_no"`            // 支付卡号
		CardType         int64          `db:"card_type" json:"card_type,string"`          // 卡类型  1：储值卡 2：次卡
		CardTypeName     string         `db:"card_type_name" json:"card_type_name"`     // 卡类型名称
		CardPayPrice     string         `db:"card_pay_price" json:"card_pay_price"`     // 卡券支付金额
		CardPayOneCard   float64        `db:"card_pay_one_card" json:"card_pay_one_card,string"`  // 次卡支付金额
		CardResponse     lib.JsonNullString `db:"card_response" json:"card_response"`      // 卡消费返回信息
		OrderTime        int64          `db:"order_time" json:"order_time,string"`         // 订单时间
		PayTime          int64          `db:"pay_time" json:"pay_time,string"`           // 支付时间
		Receiver         string         `db:"receiver" json:"receiver"`           // 收货人
		ReceiverPhone    string         `db:"receiver_phone" json:"receiver_phone"`     // 收货人手机号
		GoodsName        string         `db:"goods_name" json:"goods_name"`         // 商品名称
		GoodsId          int64          `db:"goods_id" json:"goods_id,string"`           // 演出商品id
		SessionsName     string         `db:"sessions_name" json:"sessions_name"`      // 场次名称
		SessionsId       int64          `db:"sessions_id" json:"sessions_id,string"`        // 场次id
		PriceId          int64          `db:"price_id" json:"price_id,string"`           // 价格id
		SeatData         lib.JsonNullString `db:"seat_data" json:"seat_data	"`          // 座位信息
		IdSeat           int64          `db:"id_seat" json:"id_seat,string"`            // 0:未选座 1：选座
		DamaiPayResult   lib.JsonNullString `db:"damai_pay_result" json:"damai_pay_result"`   // 大麦支付返回结果
		DamaiOrderSn     string         `db:"damai_order_sn" json:"damai_order_sn"`     // 大麦订单号
		DamaiOrderStatus int64          `db:"damai_order_status" json:"damai_order_status,string"` // 大麦订单状态
		UserInfo         string         `db:"user_info" json:"user_info"`          // 观影人信息
		StartTime        string         `db:"start_time" json:"start_time"`         // 演出时间
		DeliveryTypes    int64          `db:"delivery_types" json:"delivery_types,string"`     // 1和3为电子票，2和4为纸质票。
		PayLimitTime     int64          `db:"pay_limit_time" json:"pay_limit_time,string"`     // 过期时间
		Num              int64          `db:"num" json:"num,string"`                // 数量
		PayStatus        int64          `db:"pay_status" json:"pay_status,string"`         // 付款状态 0：待付款 1：已付款 2：未付款 3:支付失败 4:未补差价5：已完成
		IsHide           int64          `db:"is_hide" json:"is_hide,string"`            // 0:正常1：关闭
		WxResponse       lib.JsonNullString `db:"wx_response" json:"wx_response"`        // 微信回调结果
		ShowImg          lib.JsonNullString `db:"showImg" json:"showImg"`            // 商品图片
		Price            float64        `db:"price" json:"price,string"`              // 单价
		VenueName        lib.JsonNullString `db:"venue_name" json:"venue_name"`         // 场馆名称
		VenueAddress     lib.JsonNullString `db:"venue_address" json:"venue_address"`      // 场馆地址
		CityName         lib.JsonNullString `db:"city_name" json:"city_name"`          // 城市名称
		WxPrice          float64        `db:"wx_price" json:"wx_price,string"`           // 微信补差金额
		IsSeat           int64          `db:"is_seat" json:"is_seat,string"`            // 1:有座位信息 0：没有座位信息
		OrderMaizuoSn    string         `db:"order_maizuo_sn" json:"order_maizuo_sn"`    // 卖座订单号
		Platform         int64          `db:"platform" json:"platform,string"`           // 1：pc 2:手机
	}
)

func newPnDamaiOrderModel(conn sqlx.SqlConn) *defaultPnDamaiOrderModel {
	return &defaultPnDamaiOrderModel{
		conn:  conn,
		table: "`pn_damai_order`",
	}
}

func (m *defaultPnDamaiOrderModel) Insert(ctx context.Context, data *PnDamaiOrder) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pnDamaiOrderRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.UserName, data.OrderSn, data.OrderStatus, data.DeliverStatus, data.PostFee, data.TotalPrice, data.TotalBasePrice, data.PayPrice, data.CardNo, data.CardType, data.CardTypeName, data.CardPayPrice, data.CardPayOneCard, data.CardResponse, data.OrderTime, data.PayTime, data.Receiver, data.ReceiverPhone, data.GoodsName, data.GoodsId, data.SessionsName, data.SessionsId, data.PriceId, data.SeatData, data.IdSeat, data.DamaiPayResult, data.DamaiOrderSn, data.DamaiOrderStatus, data.UserInfo, data.StartTime, data.DeliveryTypes, data.PayLimitTime, data.Num, data.PayStatus, data.IsHide, data.WxResponse, data.ShowImg, data.Price, data.VenueName, data.VenueAddress, data.CityName, data.WxPrice, data.IsSeat, data.OrderMaizuoSn, data.Platform)
	return ret, err
}

func (m *defaultPnDamaiOrderModel) FindOne(ctx context.Context, orderId int64) (*PnDamaiOrder, error) {
	query := fmt.Sprintf("select %s from %s where `order_id` = ? limit 1", pnDamaiOrderRows, m.table)
	var resp PnDamaiOrder
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

func (m *defaultPnDamaiOrderModel) Update(ctx context.Context, data *PnDamaiOrder) error {
	query := fmt.Sprintf("update %s set %s where `order_id` = ?", m.table, pnDamaiOrderRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.UserName, data.OrderSn, data.OrderStatus, data.DeliverStatus, data.PostFee, data.TotalPrice, data.TotalBasePrice, data.PayPrice, data.CardNo, data.CardType, data.CardTypeName, data.CardPayPrice, data.CardPayOneCard, data.CardResponse, data.OrderTime, data.PayTime, data.Receiver, data.ReceiverPhone, data.GoodsName, data.GoodsId, data.SessionsName, data.SessionsId, data.PriceId, data.SeatData, data.IdSeat, data.DamaiPayResult, data.DamaiOrderSn, data.DamaiOrderStatus, data.UserInfo, data.StartTime, data.DeliveryTypes, data.PayLimitTime, data.Num, data.PayStatus, data.IsHide, data.WxResponse, data.ShowImg, data.Price, data.VenueName, data.VenueAddress, data.CityName, data.WxPrice, data.IsSeat, data.OrderMaizuoSn, data.Platform, data.OrderId)
	return err
}

func (m *defaultPnDamaiOrderModel) Delete(ctx context.Context, orderId int64) error {
	query := fmt.Sprintf("delete from %s where `order_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, orderId)
	return err
}

func (m *defaultPnDamaiOrderModel) tableName() string {
	return m.table
}
