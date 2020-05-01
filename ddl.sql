CREATE TABLE `sys_project`
(
    `id`           bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `project_name` varchar(50)     DEFAULT NULL COMMENT '项目名',
    `project_desc` varchar(100)    DEFAULT NULL COMMENT '项目描述',
    `create_time`  datetime        DEFAULT NULL COMMENT '创建时间',
    `create_user`  varchar(50)     DEFAULT '' COMMENT '创建人',
    `update_time`  datetime        DEFAULT NULL COMMENT '更新时间',
    `update_user`  varchar(50)     DEFAULT '' COMMENT '更新人',
    `ts`           timestamp  NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '时间戳',
    `yn`           tinyint(4)      DEFAULT '1' COMMENT '是否有效， 1-正常  0-删除',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='项目';



CREATE TABLE `sys_table`
(
    `id`                bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `project_id`        bigint(20)      DEFAULT NULL COMMENT '项目表ID',
    `table_name`        varchar(50)     DEFAULT NULL COMMENT '表名',
    `table_comment`     varchar(255)    DEFAULT NULL COMMENT '表注释',
    `is_sys_table`      tinyint(1)      DEFAULT '1' COMMENT '是否系统表',
    `table_category`    varchar(50)     DEFAULT NULL COMMENT '表分类',
    `table_mode`        varchar(50)     DEFAULT NULL COMMENT '表模式：single主表、sub主子表、ext扩展表',
    `parent_table_name` varchar(100)    DEFAULT NULL COMMENT '从属于哪张表',
    `status`            tinyint(1)      DEFAULT '1' COMMENT '表状态： 1-新建未同步 2-修改未同步 3-删除未同步 10-已同步',
    `enabled`           tinyint(1)      DEFAULT '1' COMMENT '启用状态： 1-启用 2-禁用',
    `create_time`       datetime        DEFAULT NULL COMMENT '创建时间',
    `create_user`       varchar(50)     DEFAULT '' COMMENT '创建人',
    `update_time`       datetime        DEFAULT NULL COMMENT '更新时间',
    `update_user`       varchar(50)     DEFAULT '' COMMENT '更新人',
    `ts`                timestamp  NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '时间戳',
    `yn`                tinyint(4)      DEFAULT '1' COMMENT '是否有效， 1-正常  0-删除',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='表定义';



CREATE TABLE `sys_table_display`
(
    `id`              bigint(20)  NOT NULL AUTO_INCREMENT COMMENT '主键',
    `table_id`        bigint(20)       DEFAULT NULL COMMENT '表ID',
    `table_name`      varchar(50) NOT NULL COMMENT '表名',
    `need_tree`       tinyint(1)       DEFAULT '2' COMMENT '是否需要树表 1-是 2-否',
    `need_async`      tinyint(1)       DEFAULT '2' COMMENT '是否需要异步加载 1-是 2-否',
    `need_pagination` tinyint(1)       DEFAULT '0' COMMENT '是否需要分页',
    `need_search`     tinyint(1)       DEFAULT '0' COMMENT '是否分页',
    `create_time`     datetime         DEFAULT NULL COMMENT '创建时间',
    `create_user`     varchar(50)      DEFAULT '' COMMENT '创建人',
    `update_time`     datetime         DEFAULT NULL COMMENT '更新时间',
    `update_user`     varchar(50)      DEFAULT '' COMMENT '更新人',
    `ts`              timestamp   NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '时间戳',
    `yn`              tinyint(4)       DEFAULT '1' COMMENT '是否有效， 1-正常  0-删除',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='表显示定义';

CREATE TABLE `sys_column`
(
    `id`                    bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `table_id`              bigint(20)      DEFAULT NULL COMMENT '表ID',
    `table_name`            varchar(50)     DEFAULT NULL COMMENT '表名',
    `column_name`           varchar(255)    DEFAULT NULL COMMENT '列名',
    `column_comment`        varchar(255)    DEFAULT NULL COMMENT '列注释',
    `column_default`        varchar(255)    DEFAULT NULL COMMENT '默认值',
    `column_type`           varchar(255)    DEFAULT NULL COMMENT '列类型',
    `primary_key`           varchar(255)    DEFAULT NULL COMMENT '是否主键 1-是 2-否',
    `column_property`       tinyint(2)      DEFAULT '1' COMMENT '列属性，1-固定字段，2-扩展字段',
    `column_length`         int(11)         DEFAULT NULL COMMENT '列长度',
    `column_precision`      int(11)         DEFAULT '0' COMMENT '列精度',
    `nullable`              tinyint(4)      DEFAULT '1' COMMENT '是否为空',
    `ref_table_id`          bigint(20)      DEFAULT NULL COMMENT '引用的表ID',
    `ref_table_name`        varchar(100)    DEFAULT NULL COMMENT '引用的表名',
    `ref_column_id`         bigint(20)      DEFAULT NULL COMMENT '引用表的字段ID',
    `ref_column_name`       varchar(50)     DEFAULT NULL COMMENT '引用表的字段名',
    `ref_column_value_id`   varchar(50)     DEFAULT NULL COMMENT '引用表的值的字段Id，多个',
    `ref_column_value_name` varchar(50)     DEFAULT NULL COMMENT '引用表的值的字段名，多个',
    `ref_dict_category`     varchar(50)     DEFAULT NULL COMMENT '引用字典的类型',
    `ref_dict_code`         varchar(50)     DEFAULT NULL COMMENT '引用字典的编码',
    `create_time`           datetime        DEFAULT NULL COMMENT '创建时间',
    `create_user`           varchar(50)     DEFAULT '' COMMENT '创建人',
    `update_time`           datetime        DEFAULT NULL COMMENT '更新时间',
    `update_user`           varchar(50)     DEFAULT '' COMMENT '更新人',
    `ts`                    timestamp  NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '时间戳',
    `yn`                    tinyint(4)      DEFAULT '1' COMMENT '是否有效， 1-正常  0-删除',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='列定义';

CREATE TABLE `sys_column_display`
(
    `id`             bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `column_id`      bigint(20)      DEFAULT NULL COMMENT '列ID',
    `column_name`    varchar(255)    DEFAULT NULL COMMENT '列名',
    `show_name`      varchar(255)    DEFAULT NULL COMMENT '显示名',
    `display_type`   varchar(50)     DEFAULT NULL COMMENT '显示组件类型：input、select、checkbox等',
    `display_length` int(11)         DEFAULT NULL COMMENT '显示长度',
    `display_format` varchar(255)    DEFAULT NULL COMMENT '显示格式',
    `show_seq`       int(11)         DEFAULT NULL COMMENT '显示排序，越小越靠前',
    `showable`       tinyint(1)      DEFAULT '1' COMMENT '是否显示 1-是 2-否',
    `editable`       tinyint(1)      DEFAULT '1' COMMENT '是否可编辑 1-是 2-否',
    `orderable`      tinyint(1)      DEFAULT '2' COMMENT '是否可排序 1-是 2-否',
    `order_type`     tinyint(1)      DEFAULT '1' COMMENT '排序类型 1-no 2-asc 3-desc',
    `create_time`    datetime        DEFAULT NULL COMMENT '创建时间',
    `create_user`    varchar(50)     DEFAULT '' COMMENT '创建人',
    `update_time`    datetime        DEFAULT NULL COMMENT '更新时间',
    `update_user`    varchar(50)     DEFAULT '' COMMENT '更新人',
    `ts`             timestamp  NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '时间戳',
    `yn`             tinyint(4)      DEFAULT '1' COMMENT '是否有效， 1-正常  0-删除',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='列显示定义';

