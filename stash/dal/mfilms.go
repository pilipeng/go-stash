package dal

import (
	"context"
	"github.com/kevwan/go-stash/stash/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func MFilmsFormat(ctx context.Context,conn sqlx.SqlConn,m model.PnMFilms)map[string]interface{}  {
	esMap:=make(map[string]interface{})
	// ['id', 'filmsId', 'name', 'isShow','picAddr','grade','releaseDate'
	esMap["id"]=m.Id
	esMap["filmsId"]=m.FilmsId
	esMap["name"]=m.Name
	esMap["isShow"]=m.IsShow
	esMap["picAddr"]=m.PicAddr
	esMap["grade"]=m.Grade
	esMap["releaseDate"]=m.ReleaseDate
	esMap["alias_city_code"]=1

	esMap["index_name"]="by_m_films"
	if(m.IsShow==1) {
		esMap["common_is_show"] = 1
	}else{
		esMap["common_is_show"] = 0
	}
	return esMap
}
