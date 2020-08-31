DROP TABLE IF EXISTS `duty_auth`;
CREATE TABLE `duty_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '姓名',
  `telephone` varchar(50) DEFAULT '' COMMENT '电话',
  `group` varchar(50) DEFAULT '' COMMENT '所属组：计费:calculate  crm:crm',
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) COMMENT = '员工信息表' ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `duty_auth` (`id`, `username`, `password`) VALUES ('1', 'test', 'test123');

DROP TABLE IF EXISTS `duty_vacation`;
CREATE TABLE `duty_vacation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '姓名',
  `remain_vacation` int(10) DEFAULT '' COMMENT '剩余调休天数',
  `remain_annual_vacation` int(10) DEFAULT '' COMMENT '剩余年休天数',
   PRIMARY KEY (`id`)
) COMMENT = '调休信息表' ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `duty_rota`;
CREATE TABLE `duty_vacation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `datetime` date DEFAULT '' COMMENT '日期',
  `week` varchar(50) DEFAULT '' COMMENT '星期',
  `billing_late` varchar(50) DEFAULT '' COMMENT '计费晚班人员',
  `crm_late` varchar(50) DEFAULT '' COMMENT 'crm晚班人员',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
   PRIMARY KEY (`id`)
) COMMENT = '值班表' ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;


CREATE TABLE `duty_exchange` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `request_time` date DEFAULT '' COMMENT '申请日期',
 `proposer` varchar(50) DEFAULT '' COMMENT '申请人',
 `respondent` varchar(50) DEFAULT '' COMMENT '被申请对象',
 `requested_time` date DEFAULT '' COMMENT '被申请交换日期',
 `response` tinyint(3) unsigned DEFAULT ' ' COMMENT '被申请对象的回应，状态 0为拒绝、1为同意',
 `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
 `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) COMMENT = '换班申请表' ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;