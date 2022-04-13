package dal

import (
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func Table2index(tableName string, createTime int64) string {
	var idx string
	switch tableName {
	case "pn_goods":
		idx = "by_goods"
	case "pn_goods_city":
		idx = "by_goods"
	case "pn_hot_word":
		idx = "by_hot_word"
	case "pn_order":
		idx = "by_order"
		idx = GetOrderIndexName(idx, createTime)
	case "pn_order_goods":
		idx = "by_order"
		idx = GetOrderIndexName(idx, createTime)
	case "pn_scenic":
		idx = "by_scenic"
	case "pn_m_cinema":
		idx="by_m_cinema"
	case "pn_m_films":
		idx="by_m_cinema"
	case "pn_damai_show":
		idx="by_m_cinema"
	case "pn_m_order":
		idx="by_m_cinema"
	case "pn_scenic_order":
		idx="by_m_cinema"
	default:
		idx = ""
		logx.Errorf("Table2index no find map:%s", tableName)
	}
	return idx
}

func GetOrderIndexName(indexPre string, createTime int64) string {
	timeObj := time.Unix(createTime, 0)
	return indexPre + "_" + timeObj.Format("200601")
}

//聚合订单状态
const ORDER_AGGS_STATUS_PAY_WAIT int32 = 100      //待付款
const ORDER_AGGS_STATUS_PAY_ALREADY int32 = 101   //已付款
const ORDER_AGGS_STATUS_SEND_ALREADY int32 = 102  //已发货
const ORDER_AGGS_STATUS_CANCEL int32 = 103        //已取消
const ORDER_AGGS_STATUS_REFUND int32 = 104        //退款中
const ORDER_AGGS_STATUS_REFUNDED int32 = 105      //已退款
const ORDER_AGGS_STATUS_COMPLETED int32 = 106     //已完成
const ORDER_AGGS_STATUS_PAY_FALED int32 = 107     //支付失败
const ORDER_AGGS_STATUS_NO_PAY_CHAJIA int32 = 108 //未补差价
const ORDER_AGGS_STATUS_REFUND_FAILED int32 = 109 //退款失败
const ORDER_AGGS_STATUS_UNDEFINED int32 = 999     //未匹配

func OrderStatusAggMap(tableName string, orderStatus int64) int32 {
	esAggsOrderStatus := ORDER_AGGS_STATUS_UNDEFINED
	switch tableName {
	case "goods_order":
		switch orderStatus {
		case 0:
			esAggsOrderStatus= ORDER_AGGS_STATUS_PAY_WAIT
		case 1:
			esAggsOrderStatus= ORDER_AGGS_STATUS_PAY_ALREADY
		case 2:
			esAggsOrderStatus= ORDER_AGGS_STATUS_SEND_ALREADY
		case 3:
			esAggsOrderStatus= ORDER_AGGS_STATUS_CANCEL
		case 4:
			esAggsOrderStatus= ORDER_AGGS_STATUS_REFUND
		case 5:
			esAggsOrderStatus= ORDER_AGGS_STATUS_REFUNDED
		case 6:
			esAggsOrderStatus= ORDER_AGGS_STATUS_COMPLETED
		case 7:
			esAggsOrderStatus= ORDER_AGGS_STATUS_PAY_FALED
		}
	case "m_order":
		switch orderStatus {
		case 0:
			esAggsOrderStatus=ORDER_AGGS_STATUS_PAY_ALREADY
		case 1:
			esAggsOrderStatus=ORDER_AGGS_STATUS_PAY_ALREADY
		case 2:
			esAggsOrderStatus=ORDER_AGGS_STATUS_NO_PAY_CHAJIA
		case 3:
			esAggsOrderStatus=ORDER_AGGS_STATUS_REFUNDED
		case 4:
			esAggsOrderStatus=ORDER_AGGS_STATUS_REFUNDED
		case 5:
			esAggsOrderStatus=ORDER_AGGS_STATUS_CANCEL
		case 6:
			esAggsOrderStatus=ORDER_AGGS_STATUS_PAY_ALREADY
		case 7:
			esAggsOrderStatus=ORDER_AGGS_STATUS_PAY_ALREADY
		case 8:
			esAggsOrderStatus=ORDER_AGGS_STATUS_COMPLETED
		case 9:
			esAggsOrderStatus=ORDER_AGGS_STATUS_REFUND_FAILED
		}
	case "damai_order":
		switch orderStatus {
		case 0:
			esAggsOrderStatus = ORDER_AGGS_STATUS_PAY_WAIT
		case 1:
			esAggsOrderStatus = ORDER_AGGS_STATUS_PAY_ALREADY
		case 2:
			esAggsOrderStatus = ORDER_AGGS_STATUS_SEND_ALREADY
		case 3:
			esAggsOrderStatus = ORDER_AGGS_STATUS_CANCEL
		case 4:
			esAggsOrderStatus = ORDER_AGGS_STATUS_NO_PAY_CHAJIA
		case 5:
			esAggsOrderStatus = ORDER_AGGS_STATUS_REFUNDED
		case 6:
			esAggsOrderStatus = ORDER_AGGS_STATUS_COMPLETED
		}
	case "scenic_order":
		switch orderStatus {
		case 0:
			esAggsOrderStatus=ORDER_AGGS_STATUS_PAY_WAIT
		case 1:
			esAggsOrderStatus=ORDER_AGGS_STATUS_PAY_ALREADY
		case 2:
			esAggsOrderStatus=ORDER_AGGS_STATUS_NO_PAY_CHAJIA
		case 3:
			esAggsOrderStatus=ORDER_AGGS_STATUS_REFUNDED
		case 4:
			esAggsOrderStatus=ORDER_AGGS_STATUS_REFUNDED
		case 5:
			esAggsOrderStatus=ORDER_AGGS_STATUS_CANCEL
		case 6:
			esAggsOrderStatus=ORDER_AGGS_STATUS_PAY_ALREADY
		case 7:
			esAggsOrderStatus=ORDER_AGGS_STATUS_PAY_ALREADY
		case 8:
			esAggsOrderStatus=ORDER_AGGS_STATUS_COMPLETED
		}

	}
	return esAggsOrderStatus
}
