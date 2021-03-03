#==========================manage=====================================
CREATE TABLE `tb_manage`
(
    `id`         char(36)     NOT NULL primary key,
    `created_at` datetime(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` datetime(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `deleted_at` datetime(3)           DEFAULT NULL,
    `login_name` varchar(50)  NOT NULL default '',
    `phone_num`  varchar(20)  NOT NULL default '',
    `email`      varchar(100) NOT NULL default '',
    `password`   varchar(64)  NOT NULL default '',
    `picture`    varchar(500) NOT NULL default '',
    `salt`       varchar(8)   NOT NULL default '',
    INDEX `idx_phone_num` (`phone_num`) USING HASH,
    index `idx_deleted_at` (`deleted_at`) using btree,
    index `idx_updated_at` (`updated_at`) using btree,
    index `idx_created_at` (`created_at`) using btree
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
#========================================================================

#==========================visitor======================================
CREATE TABLE `tb_visitor`
(
    `id`              char(36)    NOT NULL primary key,
    `created_at`      datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at`      datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `deleted_at`      datetime(3)          DEFAULT NULL,
    `track_id`        varchar(50) NOT NULL default '',
    `last_ip`         varchar(20) NOT NULL default '',
    `last_visit_time` datetime(3) NOT NULL,
    `visit_times`     int         NOT NULL default 0,
    index `idx_deleted_at` (`deleted_at`) using btree,
    index `idx_updated_at` (`updated_at`) using btree,
    index `idx_created_at` (`created_at`) using btree
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
#======================================================================

#=============lu-short==========redirect=======================================
CREATE TABLE `tb_redirect`
(
    `id`               char(36)    NOT NULL primary key,
    `created_at`       datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at`       datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `deleted_at`       datetime(3)          DEFAULT NULL,
    `customer_id`      char(36)    not null,
    `ori_url`          char(36)    not null default '',
    `key`              varchar(50) not null default '',
    `rate_limit`       int         not null default 0,
    `times_limit_left` int         not null default 0,
    `jump_limit_left`  int         not null default 0,
    `begin_time`       datetime(3)          default null,
    `end_time`         datetime(3)          default null,
    index `idx_deleted_at` (`deleted_at`) using btree,
    index `idx_updated_at` (`updated_at`) using btree,
    index `idx_created_at` (`created_at`) using btree
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `tb_redirect_log`
(
    `id`          char(36)     NOT NULL primary key,
    `created_at`  datetime(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at`  datetime(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `deleted_at`  datetime(3)           DEFAULT NULL,
    `track_id`    varchar(50)  NOT NULL default '',
    `ip`          varchar(128) NOT NULL default '',
    `visit_time`  datetime(3)  NOT NULL,
    `device`      varchar(128) NOT NULL default '',
    `customer_id` char(36)     not null default '',
    `redirect_id` char(36)     not null default '',
    `short_url`   varchar(128) not null default '',
    `ori_url`     varchar(255) not null default '',
    index `idx_deleted_at` (`deleted_at`) using btree,
    index `idx_updated_at` (`updated_at`) using btree,
    index `idx_created_at` (`created_at`) using btree,
    index `idx_track_id` (`track_id`) using hash,
    index `idx_redirect_id` (`redirect_id`) using hash,
    index `idx_visit_time` (`visit_time`) using btree,
    index `idx_ip` (`ip`) using hash,
    index `idx_customer_id` (`customer_id`) using hash,
    index `idx_short_url` (`short_url`) using hash,
    index `idx_ori_url` (`ori_url`) using hash
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
#========================================================================

# =======================================customer========================
CREATE TABLE `tb_customer`
(
    `id`         char(36)     NOT NULL primary key,
    `created_at` datetime(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` datetime(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `deleted_at` datetime(3)           DEFAULT NULL,
    `login_name` varchar(50)  NOT NULL default '',
    `phone_num`  varchar(20)  NOT NULL default '',
    `email`      varchar(100) NOT NULL default '',
    `password`   varchar(64)  NOT NULL default '',
    `picture`    varchar(500) NOT NULL default '',
    `salt`       varchar(8)   NOT NULL default '',
    `vip`        int          not null comment '0:普通 1:一阶',
    INDEX `idx_phone_num` (`phone_num`) USING HASH,
    index `idx_deleted_at` (`deleted_at`) using btree,
    index `idx_updated_at` (`updated_at`) using btree,
    index `idx_created_at` (`created_at`) using btree
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
#=======================================================================
