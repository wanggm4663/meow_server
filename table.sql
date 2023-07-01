CREATE TABLE `category` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `pid` int NOT NULL DEFAULT '0' COMMENT '上级分类id',
                            `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '类目名称',
                            `name_en` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
                            `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 0 删除 1 正常',
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            PRIMARY KEY (`id`),
                            KEY `pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `product` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `cate1_id` int NOT NULL DEFAULT '0' COMMENT '一级类目id',
                           `cate2_id` int NOT NULL DEFAULT '0' COMMENT '二级类目id',
                           `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
                           `price` int NOT NULL DEFAULT '0' COMMENT '价格',
                           `img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '头图',
                           `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 0 删除 1正常',
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           PRIMARY KEY (`id`),
                           KEY `cate` (`cate1_id`,`cate2_id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `cart` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `uid` int NOT NULL DEFAULT '0' COMMENT '用户id',
                        `product_id` int NOT NULL DEFAULT '0' COMMENT '商品id',
                        `count` int NOT NULL DEFAULT '0' COMMENT '商品数量',
                        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `order` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `uid` int NOT NULL DEFAULT '0' COMMENT '用户id',
                         `product_id` int NOT NULL DEFAULT '0' COMMENT '商品id',
                         `price` int NOT NULL DEFAULT '0' COMMENT '总金额',
                         `real_price` int NOT NULL DEFAULT '0' COMMENT '实际支付金额',
                         `order_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
                         `content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '备注',
                         `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 0 删除 1 待支付 2 支付中 3 已支付 4 已完成',
                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`),
                         KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `cart` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `uid` int NOT NULL DEFAULT '0' COMMENT '用户id',
                        `product_id` int NOT NULL DEFAULT '0' COMMENT '商品id',
                        `count` int NOT NULL DEFAULT '0' COMMENT '商品数量',
                        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;