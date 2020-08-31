DROP TABLE IF EXISTS `duty_auth`;
CREATE TABLE `duty_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `duty_auth` (`id`, `username`, `password`) VALUES ('1', 'test', 'test123');

DROP TABLE IF EXISTS `duty_vacation`;
CREATE TABLE `duty_vacation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '姓名',
  `remain_vacation` int(10) DEFAULT '' COMMENT '剩余调休天数',
  `remain_annual_vacation` int(10) DEFAULT '' COMMENT '剩余年休天数',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `duty_rota`;
CREATE TABLE `duty_vacation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `datetime` date DEFAULT '' COMMENT '日期',
  `billing_late` varchar(50) DEFAULT '' COMMENT '计费晚班',
  `crm_late` varchar(50) DEFAULT '' COMMENT 'crm晚班',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;