
-- ----------------------------
-- Table structure for blog_project
-- ----------------------------
DROP TABLE IF EXISTS `blog_project`;
CREATE TABLE `blog_project`  (
  `id` INTEGER NOT NULL PRIMARY KEY,
  `node_id` varchar(64) NOT NULL DEFAULT '',
  `name` varchar(128) NOT NULL DEFAULT '',
  `full_name` varchar(512) NOT NULL DEFAULT '',
  `description` varchar(2048) NOT NULL DEFAULT '',
  `contents_url` varchar(1024) NOT NULL DEFAULT '',
  `events_url` varchar(1024) NOT NULL DEFAULT '',
  `pushed_at` varchar(32) NOT NULL DEFAULT '',
  `updated_at` varchar(32) NOT NULL DEFAULT '',
  `author_id` bigint(20) NOT NULL,
  `author_node_id` varchar(64) NOT NULL DEFAULT '',
  `avatar_url` varchar(1024) NOT NULL DEFAULT '',
  `repos_url` varchar(1024) NOT NULL DEFAULT '',
  `push_event_id` varchar(64) NOT NULL DEFAULT ''
);

-- ----------------------------
-- Table structure for blog_category
-- ----------------------------
DROP TABLE IF EXISTS `blog_category`;
CREATE TABLE `blog_category`  (
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  `project_id` bigint(20) NOT NULL,
  `name` varchar(1024) NOT NULL DEFAULT '',
  `path` varchar(2048) NOT NULL DEFAULT '',
  `deep` int(11) NOT NULL
);

-- ----------------------------
-- Table structure for blog_content
-- ----------------------------
DROP TABLE IF EXISTS `blog_content`;
CREATE TABLE `blog_content`  (
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  `project_id` bigint(20) NOT NULL,
  `name` varchar(1024) NOT NULL DEFAULT '',
  `path` varchar(2048) NOT NULL DEFAULT '',
  `deep` int(11) NOT NULL,
  `size` bigint(20) NOT NULL,
  `download_url` varchar(2048) NOT NULL DEFAULT ''
);

