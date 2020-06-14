--  -----------------
--  测试表SQL
--  -----------------
--  邮件模板表
CREATE TABLE `tb_emails` (
  `e_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `e_add_dt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `e_modify_dt` datetime DEFAULT NULL COMMENT '更新时间',
  `email_content` text COMMENT '邮件内容',
  `email_title` varchar(200) NOT NULL COMMENT '邮件标题',
  `email_name` varchar(100) NOT NULL COMMENT '邮件名',
  PRIMARY KEY (`e_id`),
  UNIQUE KEY `idx_emailName` (`email_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='邮件模板表';