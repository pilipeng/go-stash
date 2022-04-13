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
	pnMOrderFieldNames          = builder.RawFieldNames(&PnMOrder{})
	pnMOrderRows                = strings.Join(pnMOrderFieldNames, ",")
	pnMOrderRowsExpectAutoSet   = strings.Join(stringx.Remove(pnMOrderFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	pnMOrderRowsWithPlaceHolder = strings.Join(stringx.Remove(pnMOrderFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	pnMOrderModel interface {
		Insert(ctx context.Context, data *PnMOrder) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PnMOrder, error)
		FindOneByOrderId(ctx context.Context, orderId lib.JsonNullString) (*PnMOrder, error)
		Update(ctx context.Context, data *PnMOrder) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPnMOrderModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PnMOrder struct {
		Id             int64          `db:"id"`
		OrderId        lib.JsonNullString `db:"orderId"`
		UserId         int64          `db:"userId"` // 用户id
		UserName       string         `db:"userName"`
		SendMobile     string         `db:"sendMobile"`
		FilmId         int64          `db:"filmId"`      // 电影id
		FilmsName      string         `db:"filmsName"`   // 电影名
		Dimensional    string         `db:"dimensional"` // 电影类型
		CinemaId       int64          `db:"cinemaId"`    // 影院id
		CinemaName     string         `db:"cinemaName"`  // 影院名
		CityId         lib.JsonNullInt64  `db:"cityId"`
		FilmImg        string         `db:"filmImg"`     // 电影图片
		SeatName       string         `db:"seatName"`    // 座位名
		SeatIds        string         `db:"seatIds"`     // 座位
		Amount         int64          `db:"amount"`      // 票张数
		Price          int64          `db:"price"`       // 售价
		Orprice        int64          `db:"orprice"`     // 原价
		Total          int64          `db:"total"`       // 总价
		OrderStatus    int64          `db:"orderStatus"` // 0：未付款  1：已付款  2：已扣点未补差价  3：已退款（退卡退差价） 4：已退款（退卡未退差价）5:已取消   6:已确认(第三方收到订单)   7: 已发码  8：已完成(虚拟，不存在) 9.退款失败10.出票失败
		OrderTime      int64          `db:"orderTime"`   // 订单提交时间
		SendMsg        lib.JsonNullString `db:"sendMsg"`     // 短信发送信息
		PayCardNum     string         `db:"payCardNum"`
		PayCardSport   string         `db:"payCardSport"` // 支付点数/次数
		PayCardType    int64          `db:"payCardType"`  // 卡类型
		PayCardMoney   float64        `db:"payCardMoney"`
		PayCardName    string         `db:"payCardName"`    // 卡名称
		ConfirmId      string         `db:"confirmId"`      // 确认id
		OfferOrderId   string         `db:"offerOrderId"`   // 序列号
		ThirdConfirmId string         `db:"thirdConfirmId"` // 验证码
		Atm            string         `db:"atm"`            // 取票方式
		SendStatus     int64          `db:"sendStatus"`     // 信息发送状态 0：未发送  1：短信发送   2：微信发送 3：失败信息发送
		Platform       int64          `db:"Platform"`       // 0：PC  1：手机
		IsSpecial      int64          `db:"isSpecial"`      // 是否是特价票1:是,0:不是
		WxRefund       string         `db:"wxRefund"`       // 微信退款
		WxResponse     string         `db:"wx_response"`    // 微信回调信息
		OpenId         string         `db:"openId"`         // 微信opendid
		ForetellsData  lib.JsonNullString `db:"foretellsData"`  // 排期信息
		SurplusMoney   float64        `db:"surplusMoney"`
		RealInfo       string         `db:"realInfo"`       // 观影人信息
		UpdateTime     lib.Datetime      `db:"update_time"`    // 更新时间
		IsHide         int64          `db:"is_hide"`        // 是否隐藏  1：隐藏
		OrderMaizuoSn  string         `db:"orderMaizuoSn"`  // 卖座订单号
		TimeOrderId    string         `db:"timeOrderId"`    // 时光网主id
		TimeSubOrderId string         `db:"timeSubOrderId"` // 时光网子id
		ErHyi          int64          `db:"erHyi"`          // 是否二换一
		ApiOrderSn     string         `db:"apiOrderSn"`     // 接口订单号
		ZyOrderSn      string         `db:"zyOrderSn"`      // 三方资源订单号
		OrderType      int64          `db:"orderType"`      // 订单类型 0:卖座 1:时光  2:品诺卖座3：影托邦，4：淘票票，5：时光，6：鼎新
	}
)

func newPnMOrderModel(conn sqlx.SqlConn) *defaultPnMOrderModel {
	return &defaultPnMOrderModel{
		conn:  conn,
		table: "`pn_m_order`",
	}
}

func (m *defaultPnMOrderModel) Insert(ctx context.Context, data *PnMOrder) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pnMOrderRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.OrderId, data.UserId, data.UserName, data.SendMobile, data.FilmId, data.FilmsName, data.Dimensional, data.CinemaId, data.CinemaName, data.CityId, data.FilmImg, data.SeatName, data.SeatIds, data.Amount, data.Price, data.Orprice, data.Total, data.OrderStatus, data.OrderTime, data.SendMsg, data.PayCardNum, data.PayCardSport, data.PayCardType, data.PayCardMoney, data.PayCardName, data.ConfirmId, data.OfferOrderId, data.ThirdConfirmId, data.Atm, data.SendStatus, data.Platform, data.IsSpecial, data.WxRefund, data.WxResponse, data.OpenId, data.ForetellsData, data.SurplusMoney, data.RealInfo, data.IsHide, data.OrderMaizuoSn, data.TimeOrderId, data.TimeSubOrderId, data.ErHyi, data.ApiOrderSn, data.ZyOrderSn, data.OrderType)
	return ret, err
}

func (m *defaultPnMOrderModel) FindOne(ctx context.Context, id int64) (*PnMOrder, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pnMOrderRows, m.table)
	var resp PnMOrder
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPnMOrderModel) FindOneByOrderId(ctx context.Context, orderId lib.JsonNullString) (*PnMOrder, error) {
	var resp PnMOrder
	query := fmt.Sprintf("select %s from %s where `orderId` = ? limit 1", pnMOrderRows, m.table)
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

func (m *defaultPnMOrderModel) Update(ctx context.Context, data *PnMOrder) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pnMOrderRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.OrderId, data.UserId, data.UserName, data.SendMobile, data.FilmId, data.FilmsName, data.Dimensional, data.CinemaId, data.CinemaName, data.CityId, data.FilmImg, data.SeatName, data.SeatIds, data.Amount, data.Price, data.Orprice, data.Total, data.OrderStatus, data.OrderTime, data.SendMsg, data.PayCardNum, data.PayCardSport, data.PayCardType, data.PayCardMoney, data.PayCardName, data.ConfirmId, data.OfferOrderId, data.ThirdConfirmId, data.Atm, data.SendStatus, data.Platform, data.IsSpecial, data.WxRefund, data.WxResponse, data.OpenId, data.ForetellsData, data.SurplusMoney, data.RealInfo, data.IsHide, data.OrderMaizuoSn, data.TimeOrderId, data.TimeSubOrderId, data.ErHyi, data.ApiOrderSn, data.ZyOrderSn, data.OrderType, data.Id)
	return err
}

func (m *defaultPnMOrderModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPnMOrderModel) tableName() string {
	return m.table
}
