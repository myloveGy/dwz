/*
 Navicat Premium Data Transfer

 Source Server         : 本地MySQL
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : localhost:3306
 Source Schema         : short-url

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 12/12/2019 23:47:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `app_id` int(11) NOT NULL AUTO_INCREMENT,
  `app_name` varchar(32) NOT NULL COMMENT '应用名称',
  `app_key` varchar(100) NOT NULL COMMENT '应用KEY',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态 1 启用 2 禁用',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`app_id`),
  UNIQUE KEY `app_key` (`app_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='应用信息';

-- ----------------------------
-- Table structure for app_url
-- ----------------------------
DROP TABLE IF EXISTS `app_url`;
CREATE TABLE `app_url` (
  `id` int(11) NOT NULL,
  `app_id` int(11) NOT NULL COMMENT '关联应用',
  `url` varchar(255) NOT NULL COMMENT '地址信息',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `app_url` (`app_id`,`url`) COMMENT '应用url唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='应用url';

SET FOREIGN_KEY_CHECKS = 1;
