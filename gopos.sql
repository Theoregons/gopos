/*
 Navicat Premium Data Transfer

 Source Server         : Localhost
 Source Server Type    : MySQL
 Source Server Version : 80030 (8.0.30)
 Source Host           : localhost:3306
 Source Schema         : gopos

 Target Server Type    : MySQL
 Target Server Version : 80030 (8.0.30)
 File Encoding         : 65001

 Date: 14/04/2025 10:23:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `category` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `stock` bigint NULL DEFAULT NULL,
  `price` bigint NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of products
-- ----------------------------
INSERT INTO `products` VALUES (1, 'Raspberry', 'Sports & Outdoor', 19, 25723, '2004-09-22 04:56:38.000');
INSERT INTO `products` VALUES (2, 'Rambutan', 'Film Supplies', 90, 59166, '2009-02-12 23:42:22.000');
INSERT INTO `products` VALUES (3, 'Pluots se', 'Automotive Parts & Accessories', 80, 75832, '2018-11-17 13:39:51.000');
INSERT INTO `products` VALUES (4, 'ultra-Strawberry', 'Others', 98, 24640, '2023-12-18 03:16:48.000');
INSERT INTO `products` VALUES (5, 'haspberry mini', 'CDs & Vinyl', 18, 31389, '2015-03-04 21:40:31.000');
INSERT INTO `products` VALUES (6, 'Grape', 'Video Games', 74, 13973, '2003-06-27 00:16:13.000');
INSERT INTO `products` VALUES (7, 'Grape elite', 'Others', 51, 25859, '2006-11-22 02:22:28.000');
INSERT INTO `products` VALUES (8, 'Rambutan pro', 'Apps & Games', 96, 74491, '2001-05-14 22:59:21.000');
INSERT INTO `products` VALUES (9, 'Pluots mini', 'Industrial & Scientific Supplies', 71, 69507, '2009-10-22 14:05:47.000');
INSERT INTO `products` VALUES (10, 'Manso elite', 'Pet Supplies', 36, 69067, '2013-10-24 18:45:25.000');

-- ----------------------------
-- Table structure for transaction_items
-- ----------------------------
DROP TABLE IF EXISTS `transaction_items`;
CREATE TABLE `transaction_items`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` bigint UNSIGNED NULL DEFAULT NULL,
  `unit_price` bigint NULL DEFAULT NULL,
  `quantity` bigint NULL DEFAULT NULL,
  `transaction_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_transaction_items_product`(`product_id` ASC) USING BTREE,
  INDEX `fk_transaction_items_transaction`(`transaction_id` ASC) USING BTREE,
  CONSTRAINT `fk_transaction_items_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_transaction_items_transaction` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of transaction_items
-- ----------------------------
INSERT INTO `transaction_items` VALUES (1, 2, 12147, 29, 10);
INSERT INTO `transaction_items` VALUES (2, 4, 12089, 29, 3);
INSERT INTO `transaction_items` VALUES (3, 10, 78988, 18, 1);
INSERT INTO `transaction_items` VALUES (4, 2, 16120, 49, 13);
INSERT INTO `transaction_items` VALUES (5, 10, 19652, 18, 3);
INSERT INTO `transaction_items` VALUES (6, 6, 54598, 37, 8);
INSERT INTO `transaction_items` VALUES (7, 5, 20571, 49, 15);
INSERT INTO `transaction_items` VALUES (8, 9, 39817, 35, 5);
INSERT INTO `transaction_items` VALUES (9, 2, 59739, 16, 10);
INSERT INTO `transaction_items` VALUES (10, 1, 29849, 2, 6);
INSERT INTO `transaction_items` VALUES (11, 1, 83843, 7, 13);
INSERT INTO `transaction_items` VALUES (12, 10, 11239, 50, 5);
INSERT INTO `transaction_items` VALUES (13, 9, 57929, 48, 6);
INSERT INTO `transaction_items` VALUES (14, 4, 13645, 31, 11);
INSERT INTO `transaction_items` VALUES (15, 9, 90934, 39, 6);
INSERT INTO `transaction_items` VALUES (16, 2, 45946, 5, 15);
INSERT INTO `transaction_items` VALUES (17, 2, 96581, 25, 14);
INSERT INTO `transaction_items` VALUES (18, 4, 92540, 37, 3);
INSERT INTO `transaction_items` VALUES (19, 8, 68652, 13, 11);
INSERT INTO `transaction_items` VALUES (20, 10, 48125, 29, 8);
INSERT INTO `transaction_items` VALUES (21, 7, 42364, 33, 8);
INSERT INTO `transaction_items` VALUES (22, 10, 69177, 42, 6);
INSERT INTO `transaction_items` VALUES (23, 10, 28014, 3, 5);
INSERT INTO `transaction_items` VALUES (24, 2, 22823, 39, 13);
INSERT INTO `transaction_items` VALUES (25, 8, 38733, 31, 8);

-- ----------------------------
-- Table structure for transactions
-- ----------------------------
DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `invoice` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `customer_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `total_price` bigint NULL DEFAULT NULL,
  `total_payment` bigint NULL DEFAULT NULL,
  `total_change` bigint NULL DEFAULT NULL,
  `payment_method` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_transactions_user`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_transactions_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of transactions
-- ----------------------------
INSERT INTO `transactions` VALUES (1, 'IHJ-7971', 'Chu On Na', 263847, 323659, 87875, 'Tunai', 1, '2021-12-19 14:49:32.000');
INSERT INTO `transactions` VALUES (2, '61978553', 'Koyama Minato', 480070, 671647, 93315, 'Tunai', 2, '2017-02-16 13:42:04.000');
INSERT INTO `transactions` VALUES (3, '08304292', 'Theresa Wells', 34427, 345854, 86538, 'Tunai', 1, '2002-12-11 22:45:42.000');
INSERT INTO `transactions` VALUES (4, '9789436917380', 'Dawn Davis', 438214, 423533, 17569, 'Qris', 2, '2014-11-23 11:50:50.000');
INSERT INTO `transactions` VALUES (5, '977301234185', 'Theodore Sanchez', 109887, 724349, 48176, 'Tunai', 2, '2011-03-11 05:27:36.000');
INSERT INTO `transactions` VALUES (6, '80683742', 'Loui Chi Yuen', 389012, 366694, 87910, 'Tunai', 1, '2011-11-20 16:28:35.000');
INSERT INTO `transactions` VALUES (7, '7449989588469', 'Kono Kazuma', 15397, 463444, 75276, 'Qris', 2, '2020-07-02 04:19:33.000');
INSERT INTO `transactions` VALUES (8, 'SSP-2500', 'Xia Xiuying', 76326, 38334, 43442, 'Tunai', 2, '2023-08-22 17:40:12.000');
INSERT INTO `transactions` VALUES (9, '7023513590793', 'Chen Ziyi', 103840, 119262, 91800, 'Tunai', 2, '2006-09-03 15:01:49.000');
INSERT INTO `transactions` VALUES (10, '9794053301269', 'Sato Aoshi', 176975, 163948, 37125, 'Tunai', 1, '2016-10-01 02:44:49.000');
INSERT INTO `transactions` VALUES (11, '70932669', 'Sato Yuna', 82616, 553975, 19985, 'Qris', 2, '2011-11-12 01:13:48.000');
INSERT INTO `transactions` VALUES (12, '2866065496626', 'Fong Tsz Hin', 420275, 810341, 69174, 'Qris', 1, '2004-05-02 16:07:21.000');
INSERT INTO `transactions` VALUES (13, '03048481', 'Shimizu Ayato', 197542, 778111, 47082, 'Tunai', 1, '2004-05-30 08:33:31.000');
INSERT INTO `transactions` VALUES (14, 'WEQ-2073', 'Doris Cook', 356699, 751844, 97972, 'Qris', 1, '2004-05-11 23:53:13.000');
INSERT INTO `transactions` VALUES (15, '5263754447032', 'Kwong Wai Lam', 365808, 209736, 49003, 'Tunai', 1, '2014-01-30 03:49:45.000');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `role` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_users_email`(`email` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'Sakamoto Ayato', 'ayatos8@hotmail.com', 'ME4R7sndjd', 'Owner', '2006-01-15 23:16:19.000');
INSERT INTO `users` VALUES (2, 'Chang Anqi', 'changanqi@icloud.com', 'vPd0XHW8tk', 'Owner', '2020-09-27 09:53:21.000');

SET FOREIGN_KEY_CHECKS = 1;
