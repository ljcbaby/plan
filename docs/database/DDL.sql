-- ----------------------------
-- Table structure for courses
-- ----------------------------
DROP TABLE IF EXISTS `courses`;
CREATE TABLE `courses`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `foreign_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `credit` float NOT NULL,
  `hours_total` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `hours_lecture` int(0) NULL DEFAULT NULL,
  `hours_practices` int(0) NULL DEFAULT NULL,
  `hours_experiment` int(0) NULL DEFAULT NULL,
  `hours_computer` int(0) NULL DEFAULT NULL,
  `hours_self` int(0) NULL DEFAULT NULL,
  `assessment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'X代表“学校组织考试”，Y代表“学院组织考试”，C代表“考查”。',
  `show_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '双语、全英文 等',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `department_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `leader_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `code`(`code`) USING BTREE,
  INDEX `name`(`name`) USING BTREE,
  INDEX `foreignName`(`foreign_name`) USING BTREE,
  INDEX `remark`(`remark`) USING BTREE,
  INDEX `showRemark`(`show_remark`) USING BTREE,
  INDEX `departmentName`(`department_name`) USING BTREE,
  INDEX `leaderName`(`leader_name`) USING BTREE,
  INDEX `credit`(`credit`) USING BTREE,
  INDEX `assessment`(`assessment`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for programs
-- ----------------------------
DROP TABLE IF EXISTS `programs`;
CREATE TABLE `programs`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `major` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `department` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `dependency_id` int(0) UNSIGNED NULL DEFAULT NULL,
  `grade` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `content` json NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `program_ibfk_1`(`dependency_id`) USING BTREE,
  CONSTRAINT `programs_ibfk_1` FOREIGN KEY (`dependency_id`) REFERENCES `programs` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
