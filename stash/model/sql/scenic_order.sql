CREATE TABLE `pn_scenic_order` (
                                   `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                   `order_sn` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单号',
                                   `ctrip_order_sn` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '携程订单号',
                                   `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
                                   `user_name` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名（手机号）',
                                   `goods_id` int NOT NULL DEFAULT '0' COMMENT '商品id',
                                   `goods_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
                                   `goods_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片',
                                   `resource_id` int NOT NULL DEFAULT '0' COMMENT '资源id',
                                   `resource_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '资源名称',
                                   `count` int NOT NULL DEFAULT '0' COMMENT '数量',
                                   `amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '总金额（分）组合价',
                                   `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '博影销售单价',
                                   `ctrip_price` decimal(10,2) NOT NULL COMMENT '携程价',
                                   `purchase_tickets_name` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '购票人',
                                   `purchase_tickets_mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '购票人手机号',
                                   `pay_type` tinyint NOT NULL DEFAULT '0' COMMENT '支付方式 1微信 2支付宝',
                                   `order_status` tinyint NOT NULL DEFAULT '1' COMMENT ' 0：未付款  1：已付款  2：已扣点未补差价 3：已退款（退卡退差价） 4：已退款（退卡未退差价）5:已取消6:已确认(第三方收到订单)|已付款   7: 已发码|已发货  8：已完成(虚拟，不存在)',
                                   `send_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '信息发送状态 0：未发送  1：短信发送   2：微信发送 3：失败信息发送',
                                   `refundDesc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '退款描述',
                                   `order_time` int NOT NULL DEFAULT '0' COMMENT '订单时间',
                                   `pay_time` int NOT NULL DEFAULT '0' COMMENT '支付时间',
                                   `pay_card_num` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '支付卡号（多卡用逗号分隔）',
                                   `pay_card_sport` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '支付点数/次数（多卡用逗号分隔）',
                                   `pay_card_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '卡类型（一个订单只能使用同类型的卡）',
                                   `pay_card_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '卡名称（多卡用逗号分隔）',
                                   `pay_card_money` decimal(10,2) DEFAULT '0.00' COMMENT '卡总抵扣金额',
                                   `transaction_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '微信/支付宝 补差订单号',
                                   `surplus_money` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '补差金额',
                                   `order_maizuo_sn` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卖座订单号',
                                   `is_bindingcard_pay` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否绑定卡支付 0：否 1：是',
                                   `sms_send_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '短信发送状态 0：失败 1：成功',
                                   `wx_send_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '微信发送状态 0：失败 1：成功',
                                   `price_response` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '第三方初始价格',
                                   `original_price` int NOT NULL DEFAULT '0' COMMENT '传给票务的价格',
                                   `card_response` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卡支付信息',
                                   `open_time_desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
                                   `venueaddress` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '场地地址',
                                   `select_timer` date NOT NULL COMMENT '选择时间',
                                   `receiver_address_info` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细地址',
                                   `platform` int NOT NULL DEFAULT '0' COMMENT '1:手机  2：PC',
                                   `passenger_info_list` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '携程旅客信息列表',
                                   `is_hide` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除 1已删除 0正常',
                                   `third_order_info` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '第三方订单信息(\r\n  ProductUseMsg 产品设置使用方法（入园使用方法，如使用 身份证、携程订单号入园）\r\nVendorVoucher 供应商凭证（所有入园凭证拼接）(凭证码：136090505574，电子票：https://t.ctrip.cn/?GTWQmyG/kW8z 。)\r\nImageShortUrl 二维码短链\r\nVoucherFileUrl 文件凭证链接\r\nVoucherNO  凭证码\r\nVoucherCode //二维码辅助码\r\nAdmissionCertificate 入园凭证(数字凭证码)\r\nImageUrl //二维码图片地址 二维码优先展示此数据)',
                                   `third_order_status` int NOT NULL COMMENT '携程订单状态\r\n订单状态：0:已提交（分销无该状态） 1:确认\r\n中 2:已确认（和供应商确认资源） 3:待付款\r\n（分销无该状态） 4:已付款 5:成交(部分退) \r\n6:退订（用户已付款取消）7:成交 8:取消（用户\r\n未付款取消，分销无该状态）9:取消中（分销无该\r\n状态） 10:退订中\r\n订单若无退订，订单最终状态为：成交',
                                   `third_order_id` bigint unsigned NOT NULL COMMENT '携程订单id',
                                   `wx_response` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '微信回调信息',
                                   `order_note` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单备注',
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `idx_ordersn` (`order_sn`) USING BTREE,
                                   KEY `idx_ctrip_order_id` (`third_order_id`),
                                   FULLTEXT KEY `card` (`pay_card_num`)
) ENGINE=InnoDB AUTO_INCREMENT=2765 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='携程门票订单表'