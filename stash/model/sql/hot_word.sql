CREATE TABLE `pn_hot_word` (
                               `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
                               `type` tinyint unsigned NOT NULL COMMENT '应用页面 （1首页 2商城首页3 门票首页 4演出首页 5电影首页）',
                               `word` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '热词',
                               `r_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '跳转url',
                               `sort_num` int unsigned NOT NULL COMMENT ' 排序号 （倒序）',
                               `is_show` tinyint(1) NOT NULL COMMENT '是否显示',
                               `add_time` int unsigned NOT NULL COMMENT '添加时间',
                               `update_time` int unsigned NOT NULL COMMENT '修改时间',
                               PRIMARY KEY (`id`) USING BTREE,
                               KEY `idx_type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC