package config

type BinLog struct {
	Data      interface{} `json:"data"`
	Database  string    `json:"database"`
	Es        int64     `json:"es"`
	ID        int       `json:"id"`
	IsDdl     bool      `json:"isDdl"`
	MysqlType MysqlType `json:"mysqlType"`
	Old       []Old     `json:"old"`
	PkNames   []string  `json:"pkNames"`
	SQL       string    `json:"sql"`
	SQLType   SQLType   `json:"sqlType"`
	Table     string    `json:"table"`
	Ts        int64     `json:"ts"`
	Type      string    `json:"type"`
}

type MysqlType struct {
	GoodsID           string `json:"goods_id"`
	SpuID             string `json:"spu_id"`
	SkuID             string `json:"sku_id"`
	ZySn              string `json:"zy_sn"`
	SkuSn             string `json:"sku_sn"`
	FinanceSn         string `json:"finance_sn"`
	GoodsName         string `json:"goods_name"`
	GoodsImg          string `json:"goods_img"`
	ZtPrice           string `json:"zt_price"`
	GoodsPrice        string `json:"goods_price"`
	GoodsPriceSell    string `json:"goods_price_sell"`
	CurrCatID         string `json:"curr_cat_id"`
	CatID             string `json:"cat_id"`
	SecondCatID       string `json:"second_cat_id"`
	ThirdCatID        string `json:"third_cat_id"`
	ZhongtaiCat       string `json:"zhongtai_cat"`
	BrandID           string `json:"brand_id"`
	ZhongtaiBrand     string `json:"zhongtai_brand"`
	UpdateTime        string `json:"update_time"`
	AddTime           string `json:"add_time"`
	IsGift            string `json:"is_gift"`
	SupplierID        string `json:"supplier_id"`
	Supplier          string `json:"supplier"`
	SuppliersNameReal string `json:"suppliers_name_real"`
	SaleCount         string `json:"sale_count"`
	CompleteComDegree string `json:"complete_com_degree"`
	OnSale            string `json:"on_sale"`
	IsVirtual         string `json:"is_virtual"`
	IsMinPrice        string `json:"is_min_price"`
	MinSaleCount      string `json:"min_sale_count"`
	GoodsRestrictHour string `json:"goods_restrict_hour"`
	GoodsRestrictNum  string `json:"goods_restrict_num"`
	GoodsStock        string `json:"goods_stock"`
	IsShow            string `json:"is_show"`
	SortTag           string `json:"sort_tag"`
	IsShowSellout     string `json:"is_show_sellout"`
	IsRecommond       string `json:"is_recommond"`
	IsNew             string `json:"is_new"`
	IsHot             string `json:"is_hot"`
}
type Old struct {
	GoodsImg   string `json:"goods_img"`
	UpdateTime string `json:"update_time"`
}
type SQLType struct {

}
