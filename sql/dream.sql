-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.5.40-0ubuntu1 - (Ubuntu)
-- 服务器操作系统:                      debian-linux-gnu
-- HeidiSQL 版本:                  9.1.0.4867
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 dream_api_sms_v2 的数据库结构
DROP DATABASE IF EXISTS `dream_api_sms_v2`;
CREATE DATABASE IF NOT EXISTS `dream_api_sms_v2` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `dream_api_sms_v2`;


-- 导出  表 dream_api_sms_v2.t_area 结构
DROP TABLE IF EXISTS `t_area`;
CREATE TABLE IF NOT EXISTS `t_area` (
  `F_area_id` tinyint(3) unsigned NOT NULL COMMENT '区ID',
  `F_area` varchar(50) NOT NULL COMMENT '区',
  UNIQUE KEY `F_area_id` (`F_area_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='地区表';

-- 正在导出表  dream_api_sms_v2.t_area 的数据：0 rows
DELETE FROM `t_area`;
/*!40000 ALTER TABLE `t_area` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_area` ENABLE KEYS */;


-- 导出  表 dream_api_sms_v2.t_city 结构
DROP TABLE IF EXISTS `t_city`;
CREATE TABLE IF NOT EXISTS `t_city` (
  `F_city_id` tinyint(3) unsigned NOT NULL COMMENT '城市ID',
  `F_city` varchar(50) NOT NULL COMMENT '城市',
  UNIQUE KEY `F_city_id` (`F_city_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='城市表';

-- 正在导出表  dream_api_sms_v2.t_city 的数据：0 rows
DELETE FROM `t_city`;
/*!40000 ALTER TABLE `t_city` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_city` ENABLE KEYS */;


-- 导出  表 dream_api_sms_v2.t_config_pkg 结构
DROP TABLE IF EXISTS `t_config_pkg`;
CREATE TABLE IF NOT EXISTS `t_config_pkg` (
  `F_pkg` varchar(250) NOT NULL COMMENT '包名',
  `F_app_name` varchar(250) NOT NULL COMMENT '包对应的应用名字',
  `F_app_id` varchar(250) NOT NULL COMMENT 'leancloud对应的app id',
  `F_app_key` varchar(250) NOT NULL COMMENT 'leancloud对应的app key',
  `F_app_master_key` varchar(250) NOT NULL COMMENT 'leancloud对应的master key',
  `F_app_msm_template` varchar(250) NOT NULL COMMENT 'leancloud对应的短信模板名',
  UNIQUE KEY `F_pkg` (`F_pkg`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='包相关信息';

-- 正在导出表  dream_api_sms_v2.t_config_pkg 的数据：2 rows
DELETE FROM `t_config_pkg`;
/*!40000 ALTER TABLE `t_config_pkg` DISABLE KEYS */;
INSERT INTO `t_config_pkg` (`F_pkg`, `F_app_name`, `F_app_id`, `F_app_key`, `F_app_master_key`, `F_app_msm_template`) VALUES
	('abc', '刷题吧', 'ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'template1'),
	('com.readboy.gaokao.debug', '刷题吧', 'ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'template1');
/*!40000 ALTER TABLE `t_config_pkg` ENABLE KEYS */;


-- 导出  表 dream_api_sms_v2.t_config_response 结构
DROP TABLE IF EXISTS `t_config_response`;
CREATE TABLE IF NOT EXISTS `t_config_response` (
  `F_response_no` smallint(5) NOT NULL COMMENT '响应code',
  `F_response_msg` varchar(50) NOT NULL COMMENT '响应信息'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='api的响应配置';

-- 正在导出表  dream_api_sms_v2.t_config_response 的数据：10 rows
DELETE FROM `t_config_response`;
/*!40000 ALTER TABLE `t_config_response` DISABLE KEYS */;
INSERT INTO `t_config_response` (`F_response_no`, `F_response_msg`) VALUES
	(0, '成功'),
	(-1, '失败'),
	(-2, '已注册'),
	(-3, '密码不符合规则'),
	(-4, '没有注册'),
	(-5, '用户名或密码错误'),
	(-6, '签名错误'),
	(-7, '包名不存在'),
	(-8, '现有密码错误'),
	(-9, '密码错误');
/*!40000 ALTER TABLE `t_config_response` ENABLE KEYS */;


-- 导出  表 dream_api_sms_v2.t_county 结构
DROP TABLE IF EXISTS `t_county`;
CREATE TABLE IF NOT EXISTS `t_county` (
  `F_county_id` tinyint(3) unsigned NOT NULL COMMENT '县ID',
  `F_county` varchar(50) NOT NULL COMMENT '县',
  UNIQUE KEY `F_county_id` (`F_county_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='县表';

-- 正在导出表  dream_api_sms_v2.t_county 的数据：0 rows
DELETE FROM `t_county`;
/*!40000 ALTER TABLE `t_county` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_county` ENABLE KEYS */;


-- 导出  表 dream_api_sms_v2.t_grage 结构
DROP TABLE IF EXISTS `t_grage`;
CREATE TABLE IF NOT EXISTS `t_grage` (
  `F_grade_id` tinyint(3) unsigned NOT NULL COMMENT '年级ID',
  `F_grade` varchar(50) NOT NULL COMMENT '年级',
  UNIQUE KEY `F_grade_id` (`F_grade_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='年级表';

-- 正在导出表  dream_api_sms_v2.t_grage 的数据：0 rows
DELETE FROM `t_grage`;
/*!40000 ALTER TABLE `t_grage` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_grage` ENABLE KEYS */;


-- 导出  表 dream_api_sms_v2.t_province 结构
DROP TABLE IF EXISTS `t_province`;
CREATE TABLE IF NOT EXISTS `t_province` (
  `F_province_id` tinyint(3) unsigned NOT NULL COMMENT '省份ID',
  `F_province` varchar(50) NOT NULL COMMENT '省份',
  UNIQUE KEY `F_province_id` (`F_province_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='省份表';

-- 正在导出表  dream_api_sms_v2.t_province 的数据：0 rows
DELETE FROM `t_province`;
/*!40000 ALTER TABLE `t_province` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_province` ENABLE KEYS */;


-- 导出  表 dream_api_sms_v2.t_school 结构
DROP TABLE IF EXISTS `t_school`;
CREATE TABLE IF NOT EXISTS `t_school` (
  `F_school_id` tinyint(3) unsigned NOT NULL COMMENT '学校ID',
  `F_school` varchar(50) NOT NULL COMMENT '学校',
  UNIQUE KEY `F_school_id` (`F_school_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='学校表';

-- 正在导出表  dream_api_sms_v2.t_school 的数据：0 rows
DELETE FROM `t_school`;
/*!40000 ALTER TABLE `t_school` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_school` ENABLE KEYS */;


-- 导出  表 dream_api_sms_v2.t_sms_rate 结构
DROP TABLE IF EXISTS `t_sms_rate`;
CREATE TABLE IF NOT EXISTS `t_sms_rate` (
  `F_action` char(32) NOT NULL COMMENT '动作，由(手机号码，包名，一起md5构成)',
  `F_last_timestamp` datetime NOT NULL COMMENT '时间',
  UNIQUE KEY `F_action` (`F_action`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='记录短信发送的频率，用于限制短信的频繁发送，暂时的，会改为redis';

-- 正在导出表  dream_api_sms_v2.t_sms_rate 的数据：0 rows
DELETE FROM `t_sms_rate`;
/*!40000 ALTER TABLE `t_sms_rate` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_sms_rate` ENABLE KEYS */;


-- 导出  表 dream_api_sms_v2.t_token 结构
DROP TABLE IF EXISTS `t_token`;
CREATE TABLE IF NOT EXISTS `t_token` (
  `F_user_name` varchar(50) NOT NULL COMMENT '用户名',
  `F_pkg` varchar(250) NOT NULL COMMENT '包名',
  `F_token` char(40) NOT NULL COMMENT 'token',
  `F_expire_datetime` datetime NOT NULL COMMENT 'token到期时间',
  UNIQUE KEY `F_user_name` (`F_user_name`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='token表';

-- 正在导出表  dream_api_sms_v2.t_token 的数据：0 rows
DELETE FROM `t_token`;
/*!40000 ALTER TABLE `t_token` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_token` ENABLE KEYS */;


-- 导出  表 dream_api_sms_v2.t_user 结构
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE IF NOT EXISTS `t_user` (
  `F_user_name` varchar(50) NOT NULL COMMENT '用户名',
  `F_user_password` char(40) NOT NULL COMMENT '用户密码',
  `F_gender` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '性别(1男，2女)',
  `F_grade_id` tinyint(3) unsigned zerofill DEFAULT NULL COMMENT '年级ID',
  `F_birthday` date DEFAULT NULL COMMENT '生日',
  `F_school_id` tinyint(3) unsigned DEFAULT NULL COMMENT '学校ID',
  `F_province_id` tinyint(3) unsigned DEFAULT NULL COMMENT '省ID',
  `F_city_id` tinyint(3) unsigned DEFAULT NULL COMMENT '市ID',
  `F_county_id` tinyint(3) unsigned DEFAULT NULL COMMENT '县ID',
  `F_area_id` tinyint(3) unsigned DEFAULT NULL COMMENT '区ID',
  `F_user_realname` varchar(50) DEFAULT NULL COMMENT '真实姓名',
  UNIQUE KEY `F_user_name` (`F_user_name`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='用户表';

-- 正在导出表  dream_api_sms_v2.t_user 的数据：0 rows
DELETE FROM `t_user`;
/*!40000 ALTER TABLE `t_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_user` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
