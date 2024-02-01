CREATE DATABASE IF NOT EXISTS auth;

USE auth;
DROP TABLE  IF EXISTS users CASCADE ;
CREATE TABLE IF NOT EXISTS users
(
    `id`          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `username`    VARCHAR(25)                                                    NOT NULL,
    `password`    VARCHAR(255)                                                    NOT NULL,
    `gender`      VARCHAR(10)  DEFAULT 'unknown'                                 NOT NULL,
    `email`       VARCHAR(25)                                                     NOT NULL,
    `role`        VARCHAR(10)  DEFAULT 'user'                                     NOT NULL,
    `create_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP                             NOT NULL,
    `update_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL
);
