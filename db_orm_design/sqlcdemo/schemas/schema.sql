CREATE TABLE `students` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `student_no` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `age` bigint NOT NULL,
--  `date` date NOT NULL,
  PRIMARY KEY (`id`)
); 

CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`)
); 