DROP TABLE IF EXISTS `duty_auth`;
CREATE TABLE `duty_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '姓名',
  `telephone` varchar(50) DEFAULT '' COMMENT '电话',
  `group` varchar(50) DEFAULT '' COMMENT '所属组：计费:calculate  crm:crm',
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `is_administrator` tinyint(3) unsigned DEFAULT '0'  COMMENT '是否管理员，0：否，1：是',
  `created_on` timestamp   DEFAULT NOW() COMMENT '创建时间',
  `created_by` varchar(50) DEFAULT '' COMMENT '创建人',
  `modified_on` timestamp   DEFAULT NOW()  COMMENT '修改时间',
  `modified_by` varchar(50) DEFAULT '' COMMENT '修改人',
  `backup1` varchar(50) DEFAULT '' COMMENT '',
  `backup2` varchar(50) DEFAULT '' COMMENT '',
  PRIMARY KEY (`id`)
) COMMENT = '员工信息表' ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `duty_vacation`;
CREATE TABLE `duty_vacation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '姓名',
  `remain_vacation` int(10) DEFAULT 0 COMMENT '剩余调休天数',
  `remain_annual_vacation` int(10) DEFAULT 0 COMMENT '剩余年休天数',
  `update_time` timestamp   DEFAULT NOW() COMMENT '更新时间',
  `backup1` varchar(50) DEFAULT '' COMMENT '',
  `backup2` varchar(50) DEFAULT '' COMMENT '',
   PRIMARY KEY (`id`)
) COMMENT = '调休信息表' ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `duty_rota`;
CREATE TABLE `duty_rota` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `datetime` varchar(50) DEFAULT '' COMMENT '日期',
  `week` varchar(50) DEFAULT '' COMMENT '星期',
  `billing_late` varchar(50) DEFAULT '' COMMENT '计费晚班人员',
  `billing_weekend_late` varchar(50) DEFAULT '' COMMENT '计费周末晚班人员',
  `crm_late` varchar(50) DEFAULT '' COMMENT 'crm晚班人员',
  `crm_weekend_late` varchar(50) DEFAULT '' COMMENT 'crm周末晚班人员',
  `crm_duty` varchar(50) DEFAULT '' COMMENT 'crm值班人员',
  `created_on` timestamp   DEFAULT NOW() COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` timestamp   DEFAULT NOW()  COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `backup1` varchar(50) DEFAULT '' COMMENT '',
  `backup2` varchar(50) DEFAULT '' COMMENT '',
   PRIMARY KEY (`id`)
) COMMENT = '值班表' ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `duty_exchange`;
CREATE TABLE `duty_exchange` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `request_time` varchar(50) DEFAULT '' COMMENT '申请日期',
 `proposer` varchar(50) DEFAULT '' COMMENT '申请人',
 `respondent` varchar(50) DEFAULT '' COMMENT '被申请对象',
 `requested_time` varchar(50) DEFAULT '' COMMENT '被申请交换日期',
 `response` tinyint(1) COMMENT '被申请对象的回应，状态 0为拒绝、1为同意',
 `created_on` timestamp   DEFAULT NOW()  COMMENT '创建时间',
 `response_on` timestamp   DEFAULT NOW()  COMMENT '回应时间',
 `backup1` varchar(50) DEFAULT '' COMMENT '',
 `backup2` varchar(50) DEFAULT '' COMMENT '',
  PRIMARY KEY (`id`)
) COMMENT = '换班申请表' ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `duty_rest`;
CREATE TABLE `duty_rest` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `datetime` varchar(50) DEFAULT '' COMMENT '申请调休日期',
 `type` tinyint(3) unsigned DEFAULT '0' COMMENT '申请调休类型，0：上午，1：下午，2：全天',
 `proposer` varchar(50) DEFAULT '' COMMENT '申请人',
 `checker` varchar(50) DEFAULT '' COMMENT '审核人',
 `response` tinyint(1)  COMMENT '审核人的批复，状态 0为拒绝、1为同意',
 `created_on` timestamp   DEFAULT NOW() COMMENT '创建时间',
 `response_on` timestamp   DEFAULT NOW() COMMENT '审批时间',
 `backup1` varchar(50) DEFAULT '' COMMENT '',
 `backup2` varchar(50) DEFAULT '' COMMENT '',
  PRIMARY KEY (`id`)
) COMMENT = '调休申请表' ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
