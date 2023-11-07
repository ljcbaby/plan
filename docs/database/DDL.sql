-- ----------------------------
-- Table structure for course
-- ----------------------------
DROP TABLE IF EXISTS `course`;
CREATE TABLE `course`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `foreignName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `credit` int(0) NOT NULL,
  `hoursTotal` json NOT NULL,
  `hoursLecture` int(0) NOT NULL,
  `hoursPractices` int(0) NOT NULL,
  `hoursExperiment` int(0) NOT NULL,
  `hoursComputer` int(0) NOT NULL,
  `hoursSelf` int(0) NOT NULL,
  `assessment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'X代表“学校组织考试”，Y代表“学院组织考试”，C代表“考查”。',
  `showRemark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '双语、全英文 等',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `departmentName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `leaderName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `name`(`name`) USING BTREE,
  INDEX `foreignName`(`foreignName`) USING BTREE,
  INDEX `remark`(`remark`) USING BTREE,
  INDEX `showRemark`(`showRemark`) USING BTREE,
  INDEX `departmentName`(`departmentName`) USING BTREE,
  INDEX `leaderName`(`leaderName`) USING BTREE,
  INDEX `credit`(`credit`) USING BTREE,
  INDEX `assessment`(`assessment`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for program
-- ----------------------------
DROP TABLE IF EXISTS `program`;
CREATE TABLE `program`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `major` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `department` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `dependencyId` int(0) UNSIGNED NOT NULL,
  `grade` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `content` json NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `program_ibfk_1`(`dependencyId`) USING BTREE,
  CONSTRAINT `program_ibfk_1` FOREIGN KEY (`dependencyId`) REFERENCES `program` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
