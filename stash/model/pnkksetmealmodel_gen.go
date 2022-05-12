// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/kevwan/go-stash/stash/lib"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
)

var (
	pnKkSetmealFieldNames          = builder.RawFieldNames(&PnKkSetmeal{})
	pnKkSetmealRows                = strings.Join(pnKkSetmealFieldNames, ",")
	pnKkSetmealRowsExpectAutoSet   = strings.Join(stringx.Remove(pnKkSetmealFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	pnKkSetmealRowsWithPlaceHolder = strings.Join(stringx.Remove(pnKkSetmealFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	pnKkSetmealModel interface {
		Insert(ctx context.Context, data *PnKkSetmeal) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PnKkSetmeal, error)
		Update(ctx context.Context, data *PnKkSetmeal) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPnKkSetmealModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PnKkSetmeal struct {
		Id               int64              `db:"id" json:"id,string"`
		PackageCat       lib.JsonNullString `db:"packageCat" json:"packageCat"`                 // 套餐大类别
		IsKKTC           sql.NullInt64      `db:"isKKTC" json:"isKKTC"`                         // 是否为康康套餐
		PackageId        int64              `db:"packageId" json:"packageId,string"`         // 套餐ID
		PackageName      lib.JsonNullString `db:"packageName" json:"packageName"`               // 套餐名称
		PackageCode      lib.JsonNullString `db:"packageCode" json:"packageCode"`               // 套餐编号
		HospId           lib.JsonNullString `db:"hospId" json:"hospId"`                         // 医院编号
		OriginalPrice    float64            `db:"originalPrice" json:"originalPrice,string"` // kk成本
		PackageImage     lib.JsonNullString `db:"packageImage" json:"packageImage"`
		PackageType      lib.JsonNullString `db:"packageType" json:"packageType"`                    // 套餐小类型
		PackageCityName  string             `db:"packageCityName" json:"packageCityName,omitempty"` // 城市名称
		PackageCityCode  int64              `db:"packageCityCode" json:"packageCityCode,string"` // 城市id
		PackageAreaName  string             `db:"packageAreaName" json:"packageAreaName,omitempty"` // 区名称
		PackageAreaCode  int64              `db:"packageAreaCode" json:"packageAreaCode,string"` // 区id
		PackageSex       int64              `db:"packageSex" json:"packageSex,string"`            // 套餐性别 1男，2女0通用
		PayType          int64              `db:"payType" json:"payType,string"`                  // 支付类型（1在线支付，0线下支付，2同时支持在线和线下）
		Attention        lib.JsonNullString `db:"attention" json:"attention"`                         // 注意事项
		Jcyy             lib.JsonNullString `db:"jcyy" json:"jcyy"`                                   // 检查意义
		Tchyzt           int64              `db:"tchyzt" json:"tchyzt,string"`                     // 预约可选婚姻状态(0，已婚  1，未婚 2，不限)
		OnSale           int64              `db:"on_sale" json:"on_sale,string"`                   // 1:上架 0：下架
		Sort             int64              `db:"sort" json:"sort,string"`                         // 排序
		ManPrice         float64            `db:"manPrice" json:"manPrice,string"`                // 男性价格
		WomanPrice       float64            `db:"womanPrice" json:"womanPrice,string"`            // 女性价格
		WomanMarryPrice  float64            `db:"womanMarryPrice" json:"womanMarryPrice,string"` // 女性已婚价格
		IsKangkang       int64              `db:"is_kangkang" json:"is_kangkang,string"`           // 1:是 0:否
		Price            float64            `db:"price" json:"price,string"`                       // 售价
		Kkzk             sql.NullFloat64    `db:"kkzk" json:"kkzk"`
		CManPrice        float64            `db:"c_manPrice" json:"c_manPrice,string"`                // 男性成本价格
		CWomanPrice      float64            `db:"c_womanPrice" json:"c_womanPrice,omistringtempty"`            // 女性成本价格
		CWomanMarryPrice float64            `db:"c_womanMarryPrice" json:"c_womanMarryPrice,string"` // 女性已婚成本价格
		IsOnline         int64              `db:"isOnline" json:"isOnline,string"`                    // 1，上线，0下线
		UpdateTime       lib.Datetime       `db:"update_time" json:"update_time"`                         // 更新时间
		ShelfTime        int64              `db:"shelf_time" json:"shelf_time,string"`                 // 上架时间
		OffShelfTime     int64              `db:"off_shelf_time" json:"off_shelf_time,string"`         // 下架时间
		CbPrice          float64            `db:"cb_price" json:"cb_price,string"`                     // 成本价
	}
)

func newPnKkSetmealModel(conn sqlx.SqlConn) *defaultPnKkSetmealModel {
	return &defaultPnKkSetmealModel{
		conn:  conn,
		table: "`pn_kk_setmeal`",
	}
}

func (m *defaultPnKkSetmealModel) Insert(ctx context.Context, data *PnKkSetmeal) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pnKkSetmealRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.PackageCat, data.IsKKTC, data.PackageId, data.PackageName, data.PackageCode, data.HospId, data.OriginalPrice, data.PackageImage, data.PackageType, data.PackageCityName, data.PackageCityCode, data.PackageAreaName, data.PackageAreaCode, data.PackageSex, data.PayType, data.Attention, data.Jcyy, data.Tchyzt, data.OnSale, data.Sort, data.ManPrice, data.WomanPrice, data.WomanMarryPrice, data.IsKangkang, data.Price, data.Kkzk, data.CManPrice, data.CWomanPrice, data.CWomanMarryPrice, data.IsOnline, data.ShelfTime, data.OffShelfTime, data.CbPrice)
	return ret, err
}

func (m *defaultPnKkSetmealModel) FindOne(ctx context.Context, id int64) (*PnKkSetmeal, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pnKkSetmealRows, m.table)
	var resp PnKkSetmeal
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPnKkSetmealModel) Update(ctx context.Context, data *PnKkSetmeal) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pnKkSetmealRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.PackageCat, data.IsKKTC, data.PackageId, data.PackageName, data.PackageCode, data.HospId, data.OriginalPrice, data.PackageImage, data.PackageType, data.PackageCityName, data.PackageCityCode, data.PackageAreaName, data.PackageAreaCode, data.PackageSex, data.PayType, data.Attention, data.Jcyy, data.Tchyzt, data.OnSale, data.Sort, data.ManPrice, data.WomanPrice, data.WomanMarryPrice, data.IsKangkang, data.Price, data.Kkzk, data.CManPrice, data.CWomanPrice, data.CWomanMarryPrice, data.IsOnline, data.ShelfTime, data.OffShelfTime, data.CbPrice, data.Id)
	return err
}

func (m *defaultPnKkSetmealModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPnKkSetmealModel) tableName() string {
	return m.table
}
