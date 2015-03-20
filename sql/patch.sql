ALTER TABLE `t_user`
	ADD COLUMN `F_user_nickname` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '昵称' AFTER `F_user_realname`;


ALTER TABLE `t_user`
	ADD COLUMN `F_avatarname` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '头像名' AFTER `F_user_nickname`;

INSERT INTO `t_config_response` (`F_response_no`, `F_response_msg`) VALUES (-20, '获取头像文件失败');
INSERT INTO `t_config_response` (`F_response_no`, `F_response_msg`) VALUES (-21, '头像文件应该小于2M');
INSERT INTO `t_config_response` (`F_response_no`, `F_response_msg`) VALUES (-22, '头像文件类型应该是png,gif,jpeg');

CREATE TABLE `t_sys_avatar` (
	`F_id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
	`F_avatar_name` VARCHAR(50) NOT NULL COMMENT '头像名(2_*)',
	PRIMARY KEY (`F_id`)
)
COMMENT='系统内置头像'
COLLATE='utf8_general_ci'
ENGINE=MyISAM
;

INSERT INTO `t_sys_avatar` (`F_id`, `F_avatar_name`) VALUES (1, '2_1.jpg');
INSERT INTO `t_sys_avatar` (`F_id`, `F_avatar_name`) VALUES (2, '2_2.jpg');
INSERT INTO `t_sys_avatar` (`F_id`, `F_avatar_name`) VALUES (3, '2_3.jpg');


INSERT INTO `t_config_response` (`F_response_no`, `F_response_msg`) VALUES (-23, '新手机号码无效，已被注册');
INSERT INTO `t_config_response` (`F_response_no`, `F_response_msg`) VALUES (-24, '真实名仅允许使用汉字、26个英文字母、阿拉伯数字组成，且小于20个字符');
INSERT INTO `t_config_response` (`F_response_no`, `F_response_msg`) VALUES (-25, '昵称仅允许使用汉字、26个英文字母、阿拉伯数字组成，且小于20个字符');


INSERT INTO `dream_api_sms_v2`.`t_config_pkg` (`F_pkg`, `F_app_name`, `F_app_id`, `F_app_key`, `F_app_master_key`, `F_app_msm_template`) VALUES ('com.team.englishsquare', '英语广场', 'ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'test2');
INSERT INTO `dream_api_sms_v2`.`t_config_pkg` (`F_pkg`, `F_app_name`, `F_app_id`, `F_app_key`, `F_app_master_key`, `F_app_msm_template`) VALUES ('com.team.englishsquare.debug', '英语广场debug', 'ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'test2');

ALTER TABLE `t_user`
	ADD COLUMN `F_user_phone` VARCHAR(15) NOT NULL COMMENT '手机号码' AFTER `F_user_password`,
	ADD INDEX `F_user_phone` (`F_user_phone`);

ALTER TABLE `t_user`
	ALTER `F_user_name` DROP DEFAULT;
ALTER TABLE `t_user`
	CHANGE COLUMN `F_user_name` `F_user_name` VARCHAR(50) NOT NULL COMMENT '用户ID' FIRST;
SELECT `DEFAULT_COLLATION_NAME` FROM `information_schema`.`SCHEMATA` WHERE `SCHEMA_NAME`='dream_api_sms_v2';

ALTER TABLE `t_token`
	ALTER `F_user_name` DROP DEFAULT;
ALTER TABLE `t_token`
	CHANGE COLUMN `F_user_name` `F_user_name` VARCHAR(50) NOT NULL COMMENT '用户ID' FIRST;

ALTER TABLE `t_sms_action_valid`
	ADD COLUMN `F_last_timestamp` DATETIME NOT NULL COMMENT '最后更新时间' AFTER `F_action`;




ALTER TABLE `t_user`
	ADD COLUMN `F_user_email` VARCHAR(50) NOT NULL COMMENT 'eamil地址' AFTER `F_user_phone`;


ALTER TABLE `t_user`
	ADD COLUMN `F_status` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '1有效,0无效' AFTER `F_avatarname`,
	DROP INDEX `F_user_phone`,
	ADD INDEX `F_user_phone` (`F_user_phone`),
	ADD INDEX `F_user_email` (`F_user_email`);

CREATE TABLE `t_email_action_valid` (
	`F_action` CHAR(32) NOT NULL COMMENT '动作，(md5(email+pkg+sms))',
	`F_last_timestamp` DATETIME NOT NULL COMMENT '最后更新时间',
	UNIQUE INDEX `F_action` (`F_action`)
)
COMMENT='记录每个动作对应的eamil证码，用于安全验证。暂时的，会改为redis'
COLLATE='utf8_general_ci'
ENGINE=MyISAM
;

CREATE TABLE `t_email_rate` (
	`F_action` CHAR(32) NOT NULL COMMENT '动作，(md5(email+pkg))',
	`F_last_timestamp` DATETIME NOT NULL COMMENT '时间',
	UNIQUE INDEX `F_action` (`F_action`)
)
COMMENT='记录email发送的频率，用于限制email的频繁发送，暂时的，会改为redis'
COLLATE='utf8_general_ci'
ENGINE=MyISAM
;



CREATE TABLE `t_auth_qq` (
	`F_user_name` VARCHAR(50) NOT NULL COMMENT '用户ID',
	`F_qq_number` VARCHAR(50) NOT NULL COMMENT 'qq号码',
	UNIQUE INDEX `F_user_name` (`F_user_name`),
	UNIQUE INDEX `F_qq_number` (`F_qq_number`)
)
COMMENT='qq认证信息'
ENGINE=MyISAM
;

CREATE TABLE `t_auth_xinlangweibo` (
	`F_user_name` VARCHAR(50) NOT NULL COMMENT '用户ID',
	`F_xinlangweibo_user` VARCHAR(50) NOT NULL COMMENT '新浪微博登录名',
	UNIQUE INDEX `F_user_name` (`F_user_name`),
	UNIQUE INDEX `F_xinlangweibo_user` (`F_xinlangweibo_user`)
)
COMMENT='新浪微博认证信息'
COLLATE='utf8_general_ci'
ENGINE=MyISAM
;




INSERT INTO `t_config_response` (`F_response_no`, `F_response_msg`) VALUES (-26, '新账号已被注册');