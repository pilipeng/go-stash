package dal

import (
	"context"
	"encoding/json"
	"github.com/kevwan/go-stash/stash/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

func MOrderFormat(ctx context.Context,conn sqlx.SqlConn,m model.PnMOrder)map[string]interface{}  {
	esMap:=make(map[string]interface{})
	//  'id', 'orderId', 'userId', 'filmId', 'filmsName', 'dimensional', 'cinemaId', 'cinemaName'
	//                ,'filmImg', 'amount', 'price', 'total', 'orderStatus', 'orderTime','foretellsData'
	//                , 'is_hide', 'payCardNum', 'cityId'
	esMap["id"]=m.Id
	esMap["order_sn"]=m.OrderId
	esMap["userId"]=m.UserId
	esMap["user_name"]=m.UserName
	esMap["orderStatus"]= OrderStatusAggMap("m_order", m.OrderStatus)
	esMap["pay_status"]=""
	esMap["post_fee"]=0
	esMap["total"]=m.Total
	esMap["orderTime"]=m.OrderTime*1000
	esMap["pay_time"]=0
	esMap["pay_limit_time"]=0
	esMap["goods_name"]=m.FilmsName
	esMap["cards"]=[]string{}
	if m.PayCardNum !=""{
		esMap["cards"]=strings.Split(m.PayCardNum,",")
	}

	goodsInfo:=make(map[string]interface{})
	goodsInfo["filmId"]=m.FilmId
	goodsInfo["filmsName"]=m.FilmsName
	goodsInfo["dimensional"]=m.Dimensional
	goodsInfo["cinemaId"]=m.CinemaId
	goodsInfo["cinemaName"]=m.CinemaName
	goodsInfo["filmImg"]=m.FilmImg
	goodsInfo["amount"]=m.Amount
	goodsInfo["cityId"]=m.CityId
	goodsInfo["price"]=m.Price
	goodsInfo["foretellsData"]=m.ForetellsData

	goodsInfoString,_:=json.Marshal(goodsInfo)
	esMap["goods_info"]=goodsInfoString
	esMap["order_data_type"]="m_order"
    if m.IsHide==1{
		esMap["common_is_show"]=0
	}else{
		esMap["common_is_show"]=1
	}
	return esMap
}
