-- +migrate Up
CREATE TABLE `person`
(
  `id`               VARCHAR(64)  NOT NULL,
  `name`        VARCHAR(64)  NOT NULL COMMENT '人员 ID',
  `age`           VARCHAR(32)  NOT NULL COMMENT '手机号',
  `gender`          VARCHAR(32)  NOT NULL COMMENT 'OpenID',
  `create_time`      DATETIME     NOT NULL,
  `update_time`      DATETIME     NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`id`)
)
  COLLATE = 'utf8mb4_general_ci'
  ENGINE = InnoDB COMMENT '人员';

-- +migrate Down
DROP TABLE `person`;
