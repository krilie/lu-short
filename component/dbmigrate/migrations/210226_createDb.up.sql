SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `tb_user_master`
(
    `id`         char(36)     NOT NULL primary key,
    `created_at` datetime(3)  NOT NULL,
    `updated_at` datetime(3)  NOT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `login_name` varchar(50)  NOT NULL,
    `phone_num`  varchar(20)  NOT NULL,
    `email`      varchar(100) NOT NULL,
    `password`   varchar(64)  NOT NULL,
    `picture`    varchar(500) NOT NULL,
    `salt`       varchar(8)   NOT NULL,
    INDEX `idx_phone_num` (`phone_num`) USING HASH,
    index `idx_deleted_at` (`deleted_at`) using btree,
    index `idx_updated_at` (`updated_at`) using btree,
    index `idx_created_at` (`created_at`) using btree
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
