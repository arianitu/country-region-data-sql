CREATE TABLE IF NOT EXISTS `countries` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `code` varchar(10) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) 
) CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `regions` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `code` varchar(10) NOT NULL,
  `country_id` smallint(5) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `country_name` (`country_id`,`name`)
) CHARSET=utf8mb4;
