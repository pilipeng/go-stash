package dal

import (
	"context"
	"github.com/kevwan/go-stash/stash/model"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

func GoodsFormat(ctx context.Context,conn sqlx.SqlConn, tableData model.PnGoods,spuCityCodes map[int64][]int64) map[string]interface{} {
	esMap := make(map[string]interface{})

	esMap["index_name"] = "fx_goods"
	esMap["add_time"] =time.Now().UnixNano() / 1e6
	esMap["brand_id"] = tableData.BrandId
	esMap["cat_id"] = tableData.CatId
	esMap["flavor_id"] = 0
	esMap["flavor_name"] = ""
	if tableData.CatId == 26 {
		cfm, err := model.NewPnCakeFlavorModel(conn).FindOne(ctx,tableData.SpuId)
		if err != nil {
			logx.Errorf("no find spu_id:%s in fx_cakeFlavor ", tableData.SpuId)
		}
		if cfm != nil {
			esMap["flavor_id"] = cfm.FlavorId
			esMap["flavor_name"] = cfm.FlavorName
		}else {
			esMap["flavor_id"] = 0
			esMap["flavor_name"] = ""
		}
	}
	//esMap["city_code"] = GoodsCityCodes(conn, tableData)
	if cityCodes,ok:=spuCityCodes[tableData.SpuId];ok {
		esMap["city_code"] = cityCodes
	}else{
		esMap["city_code"] = []int{}
	}

	esMap["curr_cat_id"] = tableData.CurrCatId
	esMap["goods_id"] = tableData.GoodsId
	esMap["goods_img"] = tableData.GoodsImg
	esMap["goods_name"] = tableData.GoodsName
	esMap["goods_price"] = tableData.GoodsPrice
	esMap["goods_price_sell"] = tableData.GoodsPriceSell
	esMap["is_show"] = tableData.IsShow
	esMap["is_show_sellout"] = tableData.IsShowSellout //IsShowSellout:     tableData.IsShowSellout,
	esMap["is_recommond"] = tableData.IsRecommond
	esMap["is_virtual"] = tableData.IsVirtual
	esMap["on_sale"] = tableData.OnSale
	esMap["sale_count"] = tableData.SaleCount
	esMap["second_cat_id"] = tableData.SecondCatId
	esMap["sku_id"] = tableData.SkuId
	esMap["sku_sn"] = tableData.SkuSn
	esMap["sort_tag"] = tableData.SortTag
	esMap["spu_id"] = tableData.SpuId
	esMap["finance_sn"] = tableData.FinanceSn
	esMap["supplier"] = tableData.Supplier
	esMap["supplier_id"] = tableData.SupplierId
	esMap["suppliers_name_real"] = tableData.SuppliersNameReal
	esMap["third_cat_id"] = tableData.ThirdCatId
	esMap["update_time"] = tableData.UpdateTime * 1e3
	esMap["common_is_show"] = goodsCommonIsShow(tableData.IsShow)
	esMap["is_hot"] = tableData.IsHot
	esMap["is_new"] = tableData.IsNew

	return esMap
}


func GoodsCityCodes(ctx context.Context,conn sqlx.SqlConn, goodsSpuIds []int64) map[int64][]int64 {
	list, err := model.NewPnGoodsCityModel(conn).FindAllBySpuIds(ctx,goodsSpuIds)
	if err != nil {
		//panic(err)
		logx.Must(err)
	}

	spuCityCodes:=make(map[int64][]int64, len(goodsSpuIds))

	for _, cityInfo := range *list {
		spuCityCodes[cityInfo.SpuId]= append(spuCityCodes[cityInfo.SpuId],cityInfo.CityCode)
	}
	//lib.StructColumn(&cityIdList, *list, "CityCode", "")
	return spuCityCodes
}

func goodsCommonIsShow(isShow int64) int64 {
	if isShow == 1 {
		return 1
	} else {
		return 0
	}
}
