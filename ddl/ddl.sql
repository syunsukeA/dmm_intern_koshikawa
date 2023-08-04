CREATE TABLE `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL UNIQUE,
  `password_hash` varchar(255) NOT NULL,
  `display_name` varchar(255),
  `avatar` text,
  `header` text,
  `note` text,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

-- ToDo: 外部キーの定義
CREATE TABLE `status` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `account_id` bigint(20) NOT NULL,
  `content` varchar(255) NOT NULL, -- ToDo: contentの上限決めないとまずそう？text型に変更したい
  `url` text,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

-- ToDo: 外部キーの定義
CREATE TABLE `relationship` (
  `follower_id` bigint(20) NOT NULL,
  `followee_id` bigint(20) NOT NULL,
  PRIMARY KEY (`follower_id`, `followee_id`)
);
