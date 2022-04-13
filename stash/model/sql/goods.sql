CREATE TABLE `pn_goods` (
                            `goods_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键、自增',
                            `spu_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商品spu_id',
                            `sku_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商品sku_id',
                            `zy_sn` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '资源编码',
                            `sku_sn` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品货号',
                            `finance_sn` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '财务编码',
                            `goods_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
                            `goods_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品图片',
                            `zt_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '中台成本价',
                            `goods_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '商品价格',
                            `goods_price_sell` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '商品售卖价格',
                            `set_price_type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '售价修改类型：1固定价格 2浮动倍率',
                            `set_price_rate` decimal(10,2) unsigned NOT NULL DEFAULT '1.00' COMMENT '浮动倍率',
                            `curr_cat_id` int NOT NULL DEFAULT '0' COMMENT '商品分类',
                            `cat_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商品一级分类id',
                            `second_cat_id` int unsigned NOT NULL DEFAULT '0' COMMENT '二级分类id',
                            `third_cat_id` int unsigned NOT NULL DEFAULT '0' COMMENT '三级分类id',
                            `zhongtai_cat` int NOT NULL DEFAULT '0' COMMENT '中台分类',
                            `brand_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商品品牌',
                            `zhongtai_brand` int NOT NULL DEFAULT '0' COMMENT '中台品牌id',
                            `update_time` int unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
                            `add_time` int unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
                            `is_gift` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否为礼包商品',
                            `supplier_id` int NOT NULL DEFAULT '0' COMMENT '供货商id',
                            `supplier` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '供货商别名',
                            `suppliers_name_real` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '供货商 真实名称',
                            `sale_count` int NOT NULL DEFAULT '0' COMMENT '销量',
                            `complete_com_degree` tinyint(1) NOT NULL DEFAULT '0' COMMENT '企业专区商品上下架',
                            `on_sale` tinyint(1) NOT NULL DEFAULT '1' COMMENT '上下架状态 0：下架 1：上架',
                            `is_virtual` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否虚拟卡 0：否 1：是',
                            `is_min_price` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否该spu下价格最低商品 0：否 1：是',
                            `min_sale_count` smallint NOT NULL DEFAULT '1' COMMENT '起售数量',
                            `goods_restrict_hour` int NOT NULL DEFAULT '0' COMMENT '限制购买时间',
                            `goods_restrict_num` int NOT NULL DEFAULT '0' COMMENT '限制购买数量',
                            `goods_stock` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '礼包库存',
                            `is_show` tinyint(1) NOT NULL DEFAULT '1' COMMENT '商品是否显示 0：否 1：是',
                            `sort_tag` int NOT NULL DEFAULT '1' COMMENT '商品列表搜索排序权重，值越高权重越大，商品越靠前',
                            `is_show_sellout` tinyint(1) NOT NULL DEFAULT '0' COMMENT '无货是否显示补货中 0：否 1：是',
                            `is_recommond` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否推荐 0：否 1：是',
                            `is_new` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否新品 0：否 1：是',
                            `is_hot` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否热卖 0：否 1：是',
                            `virtual_sku_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '虚拟电子券类型 0：普通电子券 1：嘉华 2：瑞幸 3：闹闹猴',
                            PRIMARY KEY (`goods_id`) USING BTREE,
                            KEY `idx_spuid` (`spu_id`) USING BTREE,
                            KEY `idx_skuid` (`sku_id`) USING BTREE,
                            KEY `idx_price` (`goods_price`) USING BTREE,
                            KEY `idx_sellprice` (`goods_price_sell`) USING BTREE,
                            KEY `idx_cat` (`cat_id`,`second_cat_id`,`third_cat_id`) USING BTREE,
                            KEY `idx_brand` (`brand_id`) USING BTREE,
                            KEY `idx_addtime` (`add_time`) USING BTREE,
                            KEY `idx_salecnt` (`sale_count`) USING BTREE,
                            KEY `idx_currcatid` (`curr_cat_id`) USING BTREE,
                            KEY `idx_goods_id` (`goods_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1350243 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC