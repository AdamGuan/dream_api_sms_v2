CREATE TABLE `t_coin` (
	`F_user_name` VARCHAR(50) NOT NULL COMMENT '用户ID',
	`F_coin` INT(10) NOT NULL DEFAULT '0' COMMENT '用户金币',
	`F_coin_status` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '用户金币状态(1激活，0冻结)',
	UNIQUE INDEX `F_user_name` (`F_user_name`)
)
COMMENT='用户金币表'
COLLATE='utf8_general_ci'
ENGINE=MyISAM
;



CREATE TABLE `t_config_other` (
	`F_key` VARCHAR(50) NOT NULL COMMENT 'key',
	`F_value` VARCHAR(250) NOT NULL COMMENT 'value'
)
COMMENT='其它配置信息'
ENGINE=MyISAM
;

INSERT INTO `t_config_other` (`F_key`, `F_value`) VALUES ('coin', '50');

CREATE TABLE `t_ip_white_list` (
	`F_ip` CHAR(15) NOT NULL COMMENT 'IP地址',
	`F_type` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '1:IP',
	`F_status` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '1:有效,0无效'
)
COMMENT='ip白名单'
ENGINE=MyISAM
;

INSERT INTO `dream_api_sms_v2`.`t_ip_white_list` (`F_ip`) VALUES ('127.0.0.1');