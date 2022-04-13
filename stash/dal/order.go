package dal

import (
	"context"
	"encoding/json"
	"github.com/kevwan/go-stash/stash/lib"
	"github.com/kevwan/go-stash/stash/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"regexp"
	"strings"
)

func OrderFormat(ctx context.Context,conn sqlx.SqlConn, orderInfo model.PnOrder) map[string]interface{} {
	ESMap := make(map[string]interface{})

	ESMap["order_sn"] = orderInfo.OrderSn
	ESMap["user_id"] = orderInfo.UserId
	ESMap["order_status"] = OrderStatusAggMap("goods_order", orderInfo.OrderStatus)
	ESMap["pay_status"] = orderInfo.PayStatus
	ESMap["post_fee"] = orderInfo.PostFee
	ESMap["total_price"] = orderInfo.TotalPrice
	ESMap["pay_price"] = orderInfo.PayPrice
	ESMap["add_time"] = orderInfo.OrderTime * 1000
	ESMap["pay_time"] = orderInfo.PayTime * 1000
	ESMap["pay_limit_time"] = orderInfo.PayLimitTime

	glj, goodsNameList := orderGoodsInfo(ctx,conn, orderInfo)
	ESMap["goods_info"] = glj
	ESMap["goods_name"] = goodsNameList
	ESMap["order_data_type"] = "goods_order"
	ESMap["common_is_show"] = commonIsShow(orderInfo)
	ESMap["cards"] = cards(orderInfo)

	return ESMap
}

//获取商品快照和商品名称切片
func orderGoodsInfo(ctx context.Context,conn sqlx.SqlConn, orderInfo model.PnOrder) (string, []string) {
	var orderSnap []byte
	var goodsNameList []string
	//orderSnapshotBytes := orderInfo.OrderSnapshot.([]byte)
	//orderSnapshotBytes := orderInfo.OrderSnapshot
	orderSnapshotStr := orderInfo.OrderSnapshot
	orderSnapshotBytes := []byte(orderSnapshotStr)
	if orderInfo.IsShow == 1 {
		if len(orderSnapshotStr) > 0 {
			//转换博影旧数据特殊字符
			pattern := "/\t+/i"
			re, _ := regexp.Compile(pattern)
			//orderSnapshotBytes :=[]byte(orderSnapshotStr)
			orderSnap = re.ReplaceAll([]byte(orderSnapshotBytes), []byte(""))

			var orderSnapshot model.OrderSnapshot
			if err := json.Unmarshal(orderSnapshotBytes, &orderSnapshot); err != nil {
				panic(err)
			}
			lib.StructColumn(&goodsNameList, orderSnapshot, "GoodsName", "")
		}
	}

	if len(orderSnapshotStr) == 0 {
		goodsList, err := model.NewPnOrderGoodsModel(conn).FindAll(ctx,orderInfo.OrderSn)
		if err != nil {
			panic(err)
		}
		orderSnap, _ = json.Marshal(goodsList)
		lib.StructColumn(&goodsNameList, *goodsList, "goods_name", "")
	}

	if len(orderSnap) > 0 {
		return string(orderSnap), goodsNameList
	}
	return "", goodsNameList
}

func commonIsShow(orderInfo model.PnOrder) int8 {
	if orderInfo.IsShow == 1 && orderInfo.IsHide == 0 {
		return 1
	} else {
		return 0
	}
}

func cards(orderInfo model.PnOrder) []string {
	//var card []string
	card := make([]string, 0)
	if len(orderInfo.CardNo) > 0 {
		card = strings.Split(orderInfo.CardNo, ",")
	}
	return card
}
