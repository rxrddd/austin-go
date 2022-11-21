/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50737
 Source Host           : localhost:3306
 Source Schema         : austin

 Target Server Type    : MySQL
 Target Server Version : 50737
 File Encoding         : 65001

 Date: 03/05/2022 15:35:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for message_template
-- ----------------------------
DROP TABLE IF EXISTS `message_template`;
CREATE TABLE `message_template` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `audit_status` tinyint NOT NULL DEFAULT '0' COMMENT '当前消息审核状态： 10.待审核 20.审核成功 30.被拒绝',
  `id_type` tinyint NOT NULL DEFAULT '0' COMMENT '消息的发送ID类型：10. userId 20.did 30.手机号 40.openId 50.email 60.企业微信userId',
  `send_channel` tinyint NOT NULL DEFAULT '0' COMMENT '消息发送渠道：10.IM 20.Push 30.短信 40.Email 50.公众号 60.小程序 70.企业微信',
  `template_type` tinyint NOT NULL DEFAULT '0' COMMENT '10.运营类 20.技术类接口调用',
  `msg_type` tinyint NOT NULL DEFAULT '0' COMMENT '10.通知类消息 20.营销类消息 30.验证码类消息',
  `shield_type` tinyint NOT NULL DEFAULT '0' COMMENT '10.夜间不屏蔽 20.夜间屏蔽 30.夜间屏蔽(次日早上9点发送)',
  `msg_content` varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '消息内容 占位符用{$var}表示',
  `send_account` tinyint NOT NULL DEFAULT '0' COMMENT '发送账号 一个渠道下可存在多个账号',
  `creator` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '创建者',
  `updator` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '更新者',
  `auditor` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '审核人',
  `team` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '业务方团队',
  `proposer` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '业务方',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除：0.不删除 1.删除',
  `created` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deduplication_config` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '数据去重配置',
  `template_sn` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '发送消息的模版ID',
  `sms_channel` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '短信渠道 send_channel=30的时候有用  tencent腾讯云  aliyun阿里云 yunpian云片',
  PRIMARY KEY (`id`),
  KEY `idx_channel` (`send_channel`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='消息模板信息';

-- ----------------------------
-- Records of message_template
-- ----------------------------
BEGIN;
INSERT INTO `austin`.`message_template` (`id`, `name`, `audit_status`, `flow_id`, `msg_status`, `cron_task_id`, `cron_crowd_path`, `expect_push_time`, `id_type`, `send_channel`, `template_type`, `msg_type`, `shield_type`, `msg_content`, `send_account`, `creator`, `updator`, `auditor`, `team`, `proposer`, `is_deleted`, `created`, `updated`, `deduplication_config`, `template_sn`) VALUES (1, '买一送十活动', 10, '', 10, NULL, '', '', 30, 30, 20, 20, 30, '{\"content\":\"恭喜你:{$content}\",\"url\":\"\",\"title\":\"\"}', 10, 'Java3y', 'Java3y', '3y', '公众号Java3y', '三歪', 0, 1646274112, 1646275242, '', '');
INSERT INTO `austin`.`message_template` (`id`, `name`, `audit_status`, `flow_id`, `msg_status`, `cron_task_id`, `cron_crowd_path`, `expect_push_time`, `id_type`, `send_channel`, `template_type`, `msg_type`, `shield_type`, `msg_content`, `send_account`, `creator`, `updator`, `auditor`, `team`, `proposer`, `is_deleted`, `created`, `updated`, `deduplication_config`, `template_sn`) VALUES (2, '校招信息', 10, '', 10, NULL, '', '', 50, 40, 20, 10, 0, '{\"content\":\"你已成功获取到offer 内容:{$content}\",\"url\":\"\",\"title\":\"招聘通知\"}', 1, 'Java3y', 'Java3y', '3y', '公众号Java3y', '鸡蛋', 0, 1646274195, 1646274195, '', '');
INSERT INTO `austin`.`message_template` (`id`, `name`, `audit_status`, `flow_id`, `msg_status`, `cron_task_id`, `cron_crowd_path`, `expect_push_time`, `id_type`, `send_channel`, `template_type`, `msg_type`, `shield_type`, `msg_content`, `send_account`, `creator`, `updator`, `auditor`, `team`, `proposer`, `is_deleted`, `created`, `updated`, `deduplication_config`, `template_sn`) VALUES (3, '验证码通知', 10, '', 10, NULL, '', '', 30, 30, 20, 30, 0, '{\"content\":\"{$content}\",\"url\":\"\",\"title\":\"\"}', 10, 'Java3y', 'Java3y', '3y', '公众号Java3y', '孙悟空', 0, 1646275213, 1646275213, '', '');
INSERT INTO `austin`.`message_template` (`id`, `name`, `audit_status`, `flow_id`, `msg_status`, `cron_task_id`, `cron_crowd_path`, `expect_push_time`, `id_type`, `send_channel`, `template_type`, `msg_type`, `shield_type`, `msg_content`, `send_account`, `creator`, `updator`, `auditor`, `team`, `proposer`, `is_deleted`, `created`, `updated`, `deduplication_config`, `template_sn`) VALUES (4, '微信测试通知', 10, ' ', 10, NULL, NULL, NULL, 40, 50, 20, 10, 0, '{\"content\":\"{$content}\",\"url\":\"\",\"title\":\"\"}', 2, '', '', '', '', '', 0, 1646275213, 1646275213, '', '');
INSERT INTO `austin`.`message_template` (`id`, `name`, `audit_status`, `flow_id`, `msg_status`, `cron_task_id`, `cron_crowd_path`, `expect_push_time`, `id_type`, `send_channel`, `template_type`, `msg_type`, `shield_type`, `msg_content`, `send_account`, `creator`, `updator`, `auditor`, `team`, `proposer`, `is_deleted`, `created`, `updated`, `deduplication_config`, `template_sn`) VALUES (5, '钉钉测试通知', 10, ' ', 10, NULL, NULL, NULL, 40, 70, 20, 10, 0, '{\"content\":\"钉钉测试消息:\\n内容:{$content}\",\"url\":\"\",\"title\":\"\"}', 3, '', '', '', '', '', 0, 1646275213, 1646275213, '', '');
COMMIT;

-- ----------------------------
-- Table structure for send_account
-- ----------------------------
DROP TABLE IF EXISTS `send_account`;
CREATE TABLE `send_account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `send_chanel` varchar(255) NOT NULL DEFAULT '' COMMENT '发送渠道',
  `config` varchar(2000) NOT NULL DEFAULT '' COMMENT '账户配置',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '账号名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of send_account
-- ----------------------------
BEGIN;
INSERT INTO `send_account` VALUES (1, '40', '{\"host\":\"smtp.qq.com\",\"port\":25,\"username\":\"test@qq.com\",\"password\":\"tesxxxx\"}', '邮箱账号');
INSERT INTO `send_account` VALUES (2, '50', '{\"app_id\":\"app_id\",\"app_secret\":\"app_secret\",\"token\":\"weixin\"}', '微信公众号配置');
INSERT INTO `send_account` VALUES (3, '80', '{\"access_token\":\"access_token\",\"secret\":\"secret\"}', '钉钉自定义机器人');
COMMIT;

-- ----------------------------
-- Table structure for sms_record
-- ----------------------------
DROP TABLE IF EXISTS `sms_record`;
CREATE TABLE `sms_record` (
  `id` bigint NOT NULL,
  `message_template_id` bigint NOT NULL DEFAULT '0' COMMENT '消息模板ID',
  `phone` bigint NOT NULL DEFAULT '0' COMMENT '手机号',
  `supplier_id` tinyint NOT NULL DEFAULT '0' COMMENT '发送短信渠道商的ID',
  `supplier_name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '发送短信渠道商的名称',
  `msg_content` varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '短信发送的内容',
  `series_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '下发批次的ID',
  `charging_num` tinyint NOT NULL DEFAULT '0' COMMENT '计费条数',
  `report_content` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '回执内容',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '短信状态： 10.发送 20.成功 30.失败',
  `send_date` int NOT NULL DEFAULT '0' COMMENT '发送日期：20211112',
  `created` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  `request_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '唯一请求 ID',
  `biz_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '业务id',
  `send_channel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '短信渠道 tencent腾讯云  aliyun阿里云 yunpian云片',
  PRIMARY KEY (`id`),
  KEY `idx_send_date` (`send_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='短信记录信息';

-- ----------------------------
-- Records of sms_record
-- ----------------------------
BEGIN;
INSERT INTO `sms_record` VALUES (309347020626198528, 2, 0, 0, '', 'test@qq.com', '', 0, '', 0, 1651562077, 1651562077, 1651562077);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
