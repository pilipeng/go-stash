package dal

import (
	"context"
	"encoding/json"
	"github.com/kevwan/go-stash/stash/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

func ScenicOrderFormat(ctx context.Context,conn sqlx.SqlConn,order model.PnScenicOrder)map[string]interface{}  {
	esMap:=make(map[string]interface{})
	esMap["order_sn"]=order.OrderSn
	esMap["user_id"]=order.UserId
	esMap["user_name"]=order.UserName
	esMap["order_status"]=order.OrderStatus
	esMap["total_price"]=order.Amount
	esMap["add_time"]=order.OrderTime*1000
	esMap["pay_time"]=order.PayTime*1000
	esMap["post_fee"]=0
	esMap["pay_limit_time"]=0
	if order.PayCardNum==""{
		esMap["cards"]=[]string{}
	}else {
		esMap["cards"] = strings.Split(order.PayCardNum,",")
	}

	goodsInfo:=make(map[string]interface{})
	goodsInfo["goods_id"]=order.GoodsId
	goodsInfo["goods_name"]=order.GoodsName
	goodsInfo["goods_img"]=order.GoodsImg
	goodsInfo["resource_id"]=order.ResourceId
	goodsInfo["resource_name"]=order.ResourceName
	goodsInfo["count"]=order.Count
	goodsInfo["amount"]=order.Amount
	goodsInfo["price"]=order.Price
	goodsInfo["select_timer"]=order.SelectTimer
	goodsInfoString,_:=json.Marshal(goodsInfo)

	esMap["goods_info"]=goodsInfoString
	esMap["goods_name"]=order.GoodsName
	esMap["order_data_type"]="scenic_order"
	if order.IsHide==1{
		esMap["common_is_show"]=0
	}else {
		esMap["common_is_show"]=1
	}
	return esMap




}
