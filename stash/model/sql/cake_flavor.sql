CREATE TABLE `pn_cake_flavor` (
                                  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                  `flavor_id` int NOT NULL DEFAULT '0' COMMENT '口味id',
                                  `flavor_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '口味名称',
                                  `spu_id` int NOT NULL DEFAULT '0' COMMENT '蛋糕spu_id',
                                  PRIMARY KEY (`id`) USING BTREE,
                                  KEY `idx_flavorid` (`flavor_id`) USING BTREE,
                                  KEY `idx_spuid` (`spu_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='蛋糕口味表'