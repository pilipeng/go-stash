CREATE TABLE `pn_goods_city` (
                                 `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                 `spu_id` int NOT NULL DEFAULT '0' COMMENT '商品sku',
                                 `zt_city` int NOT NULL DEFAULT '0' COMMENT '中台城市id',
                                 `city_code` int NOT NULL DEFAULT '0' COMMENT '城市编码',
                                 `region_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地区名称',
                                 PRIMARY KEY (`id`) USING BTREE,
                                 KEY `idx_spuid` (`spu_id`) USING BTREE,
                                 KEY `idx_citycode` (`city_code`) USING BTREE,
                                 KEY `zt_city` (`zt_city`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14452015 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='商品城市关联表'