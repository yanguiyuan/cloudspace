CREATE DATABASE IF NOT EXISTS cloud_file;
USE cloud_file;
DROP TABLE IF EXISTS namespaces;
CREATE TABLE IF NOT EXISTS namespaces(
     `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
     `name` VARCHAR(25) NOT NULL,
     `create_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
     `update_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
     UNIQUE (`name`)
);
DROP TABLE IF EXISTS file_items;
CREATE TABLE IF NOT EXISTS file_items(
    `id` VARCHAR(12) PRIMARY KEY,
    `namespace_id` BIGINT UNSIGNED NOT NULL DEFAULT 1,
    `name` VARCHAR(255) NOT NULL,
    `type` VARCHAR(10) NOT NULL,
    `create_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `update_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (`namespace_id`) REFERENCES namespaces(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS file_indices;
CREATE TABLE IF NOT EXISTS file_indices(
    `parent_id` VARCHAR(12) NOT NULL,
    `child_id` VARCHAR(12) NOT NULL,
    PRIMARY KEY (`parent_id`, `child_id`),
    FOREIGN KEY (`parent_id`) REFERENCES file_items(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`child_id`) REFERENCES file_items(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);


use cloud_file;
DROP TABLE IF EXISTS user_namespaces;
CREATE TABLE IF NOT EXISTS user_namespaces
(
    `user_id`  BIGINT UNSIGNED NOT NULL,
    `namespace_id` BIGINT UNSIGNED NOT NULL,
    `authority` TINYINT  DEFAULT 0                                     NOT NULL,
    `create_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `update_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (`user_id`, `namespace_id`),
    FOREIGN KEY (`user_id`) REFERENCES auth.users(`id`),
    FOREIGN KEY (`namespace_id`) REFERENCES namespaces(`id`)
);
INSERT INTO namespaces(id, name) VALUE (0,'root_namespace');
INSERT INTO file_items(id, name, type) VALUE ('@root','root','root');