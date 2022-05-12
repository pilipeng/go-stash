package dal

import (
	"context"
	"encoding/json"
	"github.com/kevwan/go-stash/stash/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

func DamaiOrderFormat(ctx context.Context, conn sqlx.SqlConn, order model.PnDamaiOrder) map[string]interface{} {
	//esMap:=make(map[string]interface{})
	esMap:=make(map[string]interface{})
	esMap["order_sn"]= order.OrderSn
	esMap["user_id"]=order.OrderId
	esMap["user_name"]=order.UserName
	esMap["order_status"]=order.OrderStatus
	esMap["pay_status"]=order.PayStatus
	esMap["post_fee"]=order.PostFee
	esMap["total_price"]=order.TotalPrice
	esMap["add_time"]=order.OrderTime*1000
	esMap["pay_time"]=order.PayTime
	esMap["pay_limit_time"]=0
	if order.CardNo==""{
		esMap["cards"]=[]string{}
	}else {
		esMap["cards"] = strings.Split(order.CardNo,",")
	}

	goodsInfo:=make(map[string]interface{})
	goodsInfo["goods_id"]=order.GoodsId
	goodsInfo["goods_name"]=order.GoodsName
	goodsInfo["sessions_id"]=order.SessionsId
	goodsInfo["sessions_name"]=order.SessionsName
	goodsInfo["showImg"]=order.ShowImg
	goodsInfo["start_time"]=order.StartTime
	goodsInfo["goods_num"]=order.Num
	goodsInfo["city_name"]=order.CityName
	goodsInfo["venue_name"]=order.VenueName
	goodsInfo["venue_address"]=order.VenueAddress

	goodsInfoString,_:=json.Marshal(goodsInfo)

	esMap["goods_info"]=goodsInfoString
	esMap["goods_name"]=order.GoodsName
	esMap["order_data_type"]="damai_order"
	if order.IsHide==1{
		esMap["common_is_show"]=0
	}else{
		esMap["common_is_show"]=1
	}

	return esMap
}
