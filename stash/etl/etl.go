package etl

import (
	"context"
	"encoding/json"
	"github.com/kevwan/go-stash/stash/config"
	"github.com/kevwan/go-stash/stash/dal"
	"github.com/kevwan/go-stash/stash/lib"
	"github.com/kevwan/go-stash/stash/model"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
)

var (
	err  error
)
type MapList struct {
	EsIndexName string
	Item        interface{}
}


func Process(ctx context.Context,processor *config.Cluster,conn sqlx.SqlConn ,key string, msg []byte) map[string]MapList {
	var object json.RawMessage
	binlogObj := config.BinLog{
		Data: &object,
	}
	if err := json.Unmarshal(msg, &binlogObj); err != nil {
		panic(err)
	}


	logx.Infof("binlog-data-object:%s", object)
	var spuCityCodes map[int64][]int64
	var goodsSpuIds []int64

	list := make(map[string]MapList)
	tableName := binlogObj.Table
	switch binlogObj.Table {
	case "pn_goods":
		var l []model.PnGoods
		if err := json.Unmarshal(object, &l); err != nil {
			panic(err)
		}

		lib.StructColumn(&goodsSpuIds,l,"SpuId","")
		spuCityCodes = dal.GoodsCityCodes(ctx,conn, goodsSpuIds)
		for _, item := range l {
			list[strconv.FormatInt(item.SkuId, 10)] = MapList{
				EsIndexName: dal.Table2index(tableName, 0),
				Item:        dal.GoodsFormat(ctx,conn, item, spuCityCodes),
			}
		}

	case "pn_goods_city":
		var m []model.PnGoodsCity
		if err = json.Unmarshal(object, &m); err != nil {
			panic(err)
		}
		var spuIdS []int64
		lib.StructColumn(&spuIdS, m, "SpuId", "")
		spuIdS = lib.RemoveRepIntByMap(spuIdS)
		//goods:is_min_price判断搜索展示哪个sku
		for _, spuId := range spuIdS {
			mL, err := model.NewPnGoodsModel(conn).FindAllBySpuId(ctx,spuId)
			spuCityCodes = dal.GoodsCityCodes(ctx,conn, spuIdS)
			if err != nil {
				panic(err)
			}
			for _, i2 := range *mL {
				list[strconv.FormatInt(i2.SkuId, 10)] = MapList{
					EsIndexName: dal.Table2index(tableName, 0),
					Item:        dal.GoodsFormat(ctx,conn, i2, spuCityCodes),
				}
			}
		}
	case "pn_hot_word":
		var m []model.PnHotWord
		if err = json.Unmarshal(object, &m); err != nil {
			panic(err)
		}
		for _, i3 := range m {
			list[strconv.FormatInt(i3.Id, 10)] = MapList{
				EsIndexName: dal.Table2index(tableName, 0),
				Item:        i3,
			}
		}
	case "pn_order":
		var mqOrderList []model.PnOrder
		if err = json.Unmarshal(object, &mqOrderList); err != nil {
			panic(err)
		}
		var orderSnList []string
		lib.StructColumn(&orderSnList, mqOrderList, "OrderSn", "")
		orderList, err := model.NewPnOrderModel(conn).FindAllByOrderSn(ctx,orderSnList)
		if err != nil {
			panic(err)
		}
		for _, orderInfo := range *orderList {
			list[orderInfo.OrderSn] = MapList{
				EsIndexName: dal.Table2index(tableName, orderInfo.OrderTime),
				Item:        dal.OrderFormat(ctx,conn, orderInfo),
			}
		}
	case "pn_order_goods":
		var mqOrderGoodsList []model.PnOrderGoods
		if err = json.Unmarshal(object,&mqOrderGoodsList); err != nil{
			panic(err)
		}
		var orderSnList []string
		lib.StructColumn(&orderSnList,mqOrderGoodsList,"OrderSn","")
		orderList, err := model.NewPnOrderModel(conn).FindAllByOrderSn(ctx,orderSnList)
		if err != nil {
			panic(err)
		}
		for _, orderInfo := range *orderList {
			list[orderInfo.OrderSn] = MapList{
				EsIndexName: dal.Table2index(tableName, orderInfo.OrderTime),
				Item:        dal.OrderFormat(ctx,conn, orderInfo),
			}
		}
	case "pn_scenic":
		var scenicList []model.PnScenic
		if err = json.Unmarshal(object,&scenicList);err!=nil{
			logx.Errorf("%s",err)
			panic(err)
		}
		for _, scenic := range scenicList {
			list[strconv.FormatInt(scenic.Id,10)]=MapList{
				EsIndexName: dal.Table2index(tableName, 0),
				Item:        dal.SceincFormat(ctx,conn,scenic),
			}
		}
	case "pn_m_cinema":
		var cinemaList []model.PnMCinema
		if err=json.Unmarshal(object,&cinemaList);err!=nil{
			logx.Errorf("%s",err)
			panic(err)
		}
		for _, cinema := range cinemaList {
			list[strconv.FormatInt(cinema.Id,10)]=MapList{
				EsIndexName: dal.Table2index(tableName, 0),
				Item:        dal.McinemaFormat(ctx,conn,cinema),
			}
		}
	case "pn_m_films":
		var mFilms []model.PnMFilms
		if err=json.Unmarshal(object,&mFilms);err!=nil{
			logx.Errorf("%s",err)
			panic(err)
		}
		for _, film := range mFilms {
			list[strconv.FormatInt(film.Id,10)]=MapList{
				EsIndexName: dal.Table2index(tableName,0),
				Item:        dal.MFilmsFormat(ctx,conn,film),
			}
		}
	case "pn_damai_show":
		var damaiShowList []model.PnDamaiShow
		if err=json.Unmarshal(object,&damaiShowList);err!=nil{
			logx.Errorf("%s",err)
			panic(err)
		}
		for _, damaiShow := range damaiShowList {
			list[strconv.FormatInt(damaiShow.Id,10)]=MapList{
				EsIndexName: dal.Table2index(tableName,0),
				Item:        dal.DamaiShowFormat(ctx,conn,damaiShow),
			}
		}
	case "pn_m_order":
		var mOrderList []model.PnMOrder
		if err=json.Unmarshal(object,&mOrderList);err!=nil{
			logx.Errorf("%s",err)
			panic(err)
		}
		for _, mOrder := range mOrderList {
			list[strconv.FormatInt(mOrder.Id,10)]=MapList{
				EsIndexName: dal.Table2index(tableName,0),
				Item:  dal.MOrderFormat(ctx,conn,mOrder)     ,
			}
		}
	case "pn_damai_order":
		var damaiOrderList []model.PnDamaiOrder
		if err=json.Unmarshal(object,&damaiOrderList);err!=nil{
			logx.Errorf("%s",err)
			panic(err)
		}
		for _, damaiOrder := range damaiOrderList {
			list[strconv.FormatInt(damaiOrder.OrderId,10)]=MapList{
				EsIndexName: dal.Table2index(tableName,0),
				Item:  dal.DamaiOrderFormat(ctx,conn,damaiOrder)      ,
			}
		}
	case "pn_scenic_order":
		var scenicOrderList []model.PnScenicOrder
		if err=json.Unmarshal(object,&scenicOrderList);err!=nil{
			logx.Errorf("%s",err)
			panic(err)
		}
		for _, scenicOrder := range scenicOrderList {
			list[strconv.FormatInt(scenicOrder.Id,10)]=MapList{
				EsIndexName: dal.Table2index(tableName,0),
				Item:  dal.ScenicOrderFormat(ctx,conn,scenicOrder)      ,
			}
		}
	default:
		logx.Errorf("[warning] mq table no handle:%s",msg)
	}

	return list
}
