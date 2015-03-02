<li style="color:red;font-weight:bold;">外网api地址：useracc.dream.cn , 外网api没有说明文档, 端口与内网都一样</li>

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
