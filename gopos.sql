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

 Date: 11/04/2025 14:14:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `nama` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `jenis` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `stok` bigint NULL DEFAULT NULL,
  `harga` bigint NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of products
-- ----------------------------
INSERT INTO `products` VALUES (1, 'Pluons', 'Arts, Handicrafts & Sewing', 23, 54096, '2021-04-04 02:17:49.000');
INSERT INTO `products` VALUES (2, 'Grape', 'Musical Instrument', 25, 19058, '2020-03-16 22:29:23.000');
INSERT INTO `products` VALUES (3, 'Pluots', 'Automotive Parts & Accessories', 8, 38998, '2023-06-26 18:04:49.000');
INSERT INTO `products` VALUES (4, 'Orauge', 'Automotive Parts & Accessories', 18, 57828, '2019-12-27 09:13:14.000');
INSERT INTO `products` VALUES (5, 'Cherry elite', 'Baby', 6, 91686, '2012-11-04 05:25:03.000');

-- ----------------------------
-- Table structure for transaction_items
-- ----------------------------
DROP TABLE IF EXISTS `transaction_items`;
CREATE TABLE `transaction_items`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` bigint UNSIGNED NULL DEFAULT NULL,
  `harga_satuan` bigint NULL DEFAULT NULL,
  `quantity` bigint NULL DEFAULT NULL,
  `transaction_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_transaction_items_product`(`product_id` ASC) USING BTREE,
  INDEX `fk_transaction_items_transaction`(`transaction_id` ASC) USING BTREE,
  CONSTRAINT `fk_transaction_items_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_transaction_items_transaction` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of transaction_items
-- ----------------------------
INSERT INTO `transaction_items` VALUES (1, 5, 55184, 2, 2);
INSERT INTO `transaction_items` VALUES (2, 4, 46482, 2, 6);
INSERT INTO `transaction_items` VALUES (3, 2, 16749, 2, 8);
INSERT INTO `transaction_items` VALUES (4, 5, 67164, 7, 1);
INSERT INTO `transaction_items` VALUES (5, 3, 67765, 3, 9);
INSERT INTO `transaction_items` VALUES (6, 5, 77308, 8, 6);
INSERT INTO `transaction_items` VALUES (7, 5, 68019, 8, 1);
INSERT INTO `transaction_items` VALUES (8, 2, 54461, 2, 9);
INSERT INTO `transaction_items` VALUES (9, 2, 35794, 7, 4);
INSERT INTO `transaction_items` VALUES (10, 1, 83589, 9, 8);
INSERT INTO `transaction_items` VALUES (11, 4, 92424, 10, 1);
INSERT INTO `transaction_items` VALUES (12, 1, 76370, 9, 10);
INSERT INTO `transaction_items` VALUES (13, 4, 31604, 4, 8);
INSERT INTO `transaction_items` VALUES (14, 3, 56621, 7, 7);
INSERT INTO `transaction_items` VALUES (15, 1, 18793, 3, 8);

-- ----------------------------
-- Table structure for transactions
-- ----------------------------
DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `invoice` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `nama` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `total_harga` bigint NULL DEFAULT NULL,
  `total_bayar` bigint NULL DEFAULT NULL,
  `total_kembalian` bigint NULL DEFAULT NULL,
  `metode_pembayaran` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_transactions_user`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_transactions_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of transactions
-- ----------------------------
INSERT INTO `transactions` VALUES (1, 'e3fce61a-afb7-c439-b48d-93e0cc253777', 'Michelle Vasquez', 12980, 21652, 21156, 'To successfully                 ', 1, '2001-04-29 19:31:21.000');
INSERT INTO `transactions` VALUES (2, '64eb6a6b-ae41-4e65-d0a4-62cd96524cbd', 'Watanabe Rena', 72144, 91159, 34338, 'To start working                ', 2, '2004-03-10 20:59:46.000');
INSERT INTO `transactions` VALUES (3, '4cfc492d-bc80-94e8-495e-91aec5dc6169', 'Tong Chiu Wai', 71737, 62769, 82111, 'Actually it is                  ', 2, '2007-06-04 14:47:13.000');
INSERT INTO `transactions` VALUES (4, '4cfb15d7-49d6-b7f8-6695-1bb6fa106e3a', 'Mori Aoi', 43023, 86268, 70425, 'Navicat authorizes              ', 2, '2012-11-05 12:33:11.000');
INSERT INTO `transactions` VALUES (5, '88de89e4-d7ba-31c6-21e1-f31a0ccb638b', 'Andrew Herrera', 82687, 23912, 96459, 'A comfort zone                  ', 1, '2024-08-12 15:37:26.000');
INSERT INTO `transactions` VALUES (6, '508328d7-2872-8796-d3a3-4387d2df10ea', 'Sasaki Yamato', 38704, 20594, 52874, 'Actually it is                  ', 1, '2015-10-13 17:57:19.000');
INSERT INTO `transactions` VALUES (7, 'b2b7ec14-b3d5-9c52-65ec-c5f7abcbdaf7', 'Edith Young', 90008, 39703, 36057, 'Champions keep                  ', 1, '2022-12-27 01:34:45.000');
INSERT INTO `transactions` VALUES (8, 'd78987ee-d17b-f421-6dc3-3be5de353ec6', 'Tse Wai Lam', 43900, 63822, 91014, 'The first step                  ', 2, '2001-09-06 03:29:09.000');
INSERT INTO `transactions` VALUES (9, '98c2be8e-c648-b163-47f4-074471f3ec12', 'Li Lan', 26349, 84625, 75541, 'Difficult circumstances         ', 2, '2016-10-06 22:24:35.000');
INSERT INTO `transactions` VALUES (10, 'b01febde-5475-6b74-5136-9080758bc598', 'Zhu Yuning', 90150, 63043, 41698, 'To get a secure                 ', 2, '2003-09-18 14:08:32.000');

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
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'Kondo Momoka', 'kmomo89@icloud.com', 'To open a query                    ', 'Navicat allows                 ', '2009-09-06 00:32:29.000');
INSERT INTO `users` VALUES (2, 'Sarah Hawkins', 'sarah75@icloud.com', 'Creativity is intelligence         ', 'It collects                    ', '2003-10-04 04:00:53.000');

SET FOREIGN_KEY_CHECKS = 1;
