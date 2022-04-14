package dal

import (
	"context"
	"github.com/kevwan/go-stash/stash/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func SceincFormat(ctx context.Context, conn sqlx.SqlConn, m model.PnScenic) map[string]interface{} {
	esMap := make(map[string]interface{})
	//['id', 'name', 'comment_grade', 'comment_grade', 'star', 'market_price', 'price','theme_list',
	//                'is_valid', 'image_url', 'img_list', 'url', 'first_booking_date', 'longitude', 'latitude',
	//                'update_time', 'is_show', 'recommend_sort', 'country_id', 'city_id', 'in_china','city_code'
	esMap["id"] = m.Id
	esMap["name"] = m.Name
	esMap["comment_grade"] = m.CommentGrade
	esMap["star"] = m.Star
	esMap["market_price"] = m.MarketPrice
	esMap["price"] = m.Price
	esMap["theme_list"] = m.ThemeList
	esMap["is_valid"] = m.IsValid
	esMap["image_url"] = m.ImageUrl
	esMap["img_list"] = m.ImgList
	esMap["url"] = m.Url
	esMap["first_booking_date"] = m.FirstBookingDate
	esMap["longitude"] = m.Longitude
	esMap["latitude"] = m.Latitude
	esMap["update_time"] = m.UpdateTime * 1000
	esMap["is_show"] = m.IsShow
	esMap["recommend_sort"] = m.RecommendSort
	esMap["country_id"] = m.CountryId
	esMap["city_id"] = m.CityId
	esMap["in_china"] = m.InChina
	esMap["city_code"] = m.CityCode

	esMap["index_name"] = "by_scenic"
	if m.IsValid==1 && m.IsShow==1{
		esMap["common_is_show"] = 1
	}else {
		esMap["common_is_show"] = 0
	}

	return esMap
}
