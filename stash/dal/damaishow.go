package dal

import (
	"context"
	"github.com/kevwan/go-stash/stash/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func DamaiShowFormat(ctx context.Context, conn sqlx.SqlConn, m model.PnDamaiShow) map[string]interface{} {
	esMap:=make(map[string]interface{})
	//'id','project_id', 'project_name', 'city_id', 'city_name','is_has_seat','venue_id','lat','lng',
	//                'venue_name','venue_address','on_sale','show_pic','show_time'
	esMap["id"]=m.Id
	esMap["project_id"]=m.Id
	esMap["project_name"]=m.Id
	esMap["city_id"]=m.Id
	esMap["city_name"]=m.Id
	esMap["is_has_seat"]=m.Id
	esMap["venue_id"]=m.Id
	esMap["venue_name"]=m.Id
	esMap["venue_address"]=m.Id
	esMap["on_sale"]=m.Id
	esMap["show_pic"]=m.Id
	esMap["show_time"]=m.Id

	esMap["index_name"]="by_damai_show"
	esMap["common_is_show"]=0
	if m.OnSale==1 {
		esMap["common_is_show"]=1
	}

	return esMap

}
