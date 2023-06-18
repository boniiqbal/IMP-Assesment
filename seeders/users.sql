-- -------------------------------------------------------------
-- TablePlus 5.3.0(486)
--
-- https://tableplus.com/
--
-- Database: imp
-- Generation Time: 2023-06-18 19:47:39.2430
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `username` varchar(255) NOT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `users` (`id`, `created_at`, `password`, `username`, `full_name`) VALUES
(1, '2023-06-18 19:41:49', '$2a$10$3pcXsyWhy5ZcsO3Vo0Ly5Ow2U55lUrRhcXFbnt.dL1E8c4cou4c.K', 'boniiqbal', 'boni'),
(2, '2023-06-18 19:44:21', '$2a$10$.Bp0vheGDvkDLq6DAg9Q7O4wTI0EVbfzzu9Jsgjwy42oU5KsrU7WS', 'boniiqbal1', 'boni'),
(3, '2023-06-18 19:44:43', '$2a$10$bzckSfO4HfvTjuMU947tJ.yAmeoxaVPoQt8a9..FsccD13QamCHXe', 'boniiqbal2', 'boni'),
(4, '2023-06-18 19:44:47', '$2a$10$JbLGkTgvDgiRQ8hqzYcspOUcHcScYeyu58GmmJj7R4X0kf9cZ9aDq', 'boniiqbal3', 'boni'),
(5, '2023-06-18 19:44:51', '$2a$10$tAr8kEksnUqujsr/09tQteuvIeerIjutQ99311mT9GhMnTinMutR2', 'boniiqbal4', 'boni'),
(6, '2023-06-18 19:44:55', '$2a$10$gGGcKW5WnA5B6Ui6/kg3POSMBn/K0tN3MdB1fE8rkPndbx.mII4gm', 'boniiqbal5', 'boni'),
(7, '2023-06-18 19:44:59', '$2a$10$qIMz9jkFiXAGtBnWHQZO1e82JmcPJhNRJZjXu7fkdLrrZxgjOljN.', 'boniiqbal6', 'boni'),
(8, '2023-06-18 19:45:03', '$2a$10$Vd6/qG0qqZgEE1ExigHz0uPX4g/vipMPKMLXYZQBGyGHcQKmUtqMS', 'boniiqbal7', 'boni'),
(9, '2023-06-18 19:45:08', '$2a$10$3wd4r/2RYTLovuWZ0TkjyOiIlyDxh8H2PCTEhFZsDr2kEraNi5Bua', 'boniiqbal8', 'boni'),
(10, '2023-06-18 19:45:12', '$2a$10$gK5umGTSV2PBf9INwD0aneANsXGKNSQo9WAWCFmHKedu4/H7cl9ve', 'boniiqbal9', 'boni'),
(11, '2023-06-18 19:45:16', '$2a$10$6Ml3SZ.3..Ji7qXspjvSluTGwCI3qPqI/z2RyKJl3f4AXu4/YaZo6', 'boniiqbal10', 'boni'),
(12, '2023-06-18 19:45:21', '$2a$10$b0xgF34ZbmNyUPiHldunR.sFyny5Ivl.rLO2eNmEzyHiX.znwmr16', 'boniiqbal11', 'boni');


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;