/*
 Navicat Premium Data Transfer

 Source Server         : mysql01
 Source Server Type    : MySQL
 Source Server Version : 80020
 Source Host           : 127.0.0.1:3307
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 80020
 File Encoding         : 65001

 Date: 08/01/2022 16:24:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for test_auth
-- ----------------------------
DROP TABLE IF EXISTS `test_auth`;
CREATE TABLE `test_auth`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Secret',
  `created_on` int UNSIGNED NULL DEFAULT 0 COMMENT '新建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint UNSIGNED NULL DEFAULT 0 COMMENT '是否删除 0为未删除、1为已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '认证管理' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
