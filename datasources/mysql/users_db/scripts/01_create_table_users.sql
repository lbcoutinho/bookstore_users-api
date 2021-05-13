CREATE TABLE `users`
(
    `id`           bigint       NOT NULL AUTO_INCREMENT,
    `first_name`   varchar(255) DEFAULT NULL,
    `last_name`    varchar(255) DEFAULT NULL,
    `email`        varchar(255) NOT NULL,
    `date_created` varchar(100) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
