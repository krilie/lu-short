CREATE TABLE `tb_article_master`
(
    `id`          char(36)     NOT NULL,
    `created_at`  datetime(3)  NOT NULL,
    `updated_at`  datetime(3)  NOT NULL,
    `deleted_at`  datetime(3) DEFAULT NULL,
    `title`       varchar(256) NOT NULL,
    `description` varchar(512) NOT NULL,
    `content`     text         NOT NULL,
    `picture`     varchar(512) NOT NULL,
    `sort`        int(11)      NOT NULL,
    `pv`          int(11)      NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_tb_article_master_deleted_at` (`deleted_at`),
    KEY `idx_tb_article_master_sort` (`sort`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
