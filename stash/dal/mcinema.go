package dal

import (
	"context"
	"github.com/kevwan/go-stash/stash/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func McinemaFormat(ctx context.Context, conn sqlx.SqlConn, m model.PnMCinema) map[string]interface{} {
	esMap:=make(map[string]interface{})
	// ['id', 'cinemaName', 'cinemaId', 'cityId','by_cityId','logo','by_cityId'
	esMap["id"]=m.Id
	esMap["cinemaName"]=m.Id
	esMap["cinemaId"]=m.Id
	esMap["cityId"]=m.Id
	esMap["by_cityId"]=m.Id
	esMap["logo"]=m.Id
	esMap["by_cityId"]=m.Id


	esMap["index_name"] = "by_scenic"
	esMap["common_is_show"] = 1
	return  esMap
}
