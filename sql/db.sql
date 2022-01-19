CREATE DATABASE  IF NOT EXISTS `backend_test` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `backend_test`;
-- MySQL dump 10.13  Distrib 8.0.27, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: backend_test
-- ------------------------------------------------------
-- Server version	8.0.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `model_catalogues`
--

DROP TABLE IF EXISTS `model_catalogues`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `model_catalogues` (
  `id` int NOT NULL AUTO_INCREMENT,
  `model_name` varchar(50) NOT NULL,
  `image_url` text,
  `description` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `model_catalogues`
--

LOCK TABLES `model_catalogues` WRITE;
/*!40000 ALTER TABLE `model_catalogues` DISABLE KEYS */;
INSERT INTO `model_catalogues` VALUES (1,'Ruri Armandhani','asdasd.jpg','I am a good model','2022-01-18 11:36:41','2022-01-19 03:08:20',NULL),(2,'Gamers','asdasd','God Gamers, Handsome','2022-01-18 13:27:56','2022-01-19 01:38:56',NULL),(3,'Arman','arman.jpg','testing desc','2022-01-19 01:50:51','2022-01-19 03:37:13','2022-01-19 03:37:14');
/*!40000 ALTER TABLE `model_catalogues` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `model_schedules`
--

DROP TABLE IF EXISTS `model_schedules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `model_schedules` (
  `id` int NOT NULL AUTO_INCREMENT,
  `model_id` int NOT NULL,
  `schedule_date` timestamp NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `model_schedules_ibfk_1` (`model_id`),
  CONSTRAINT `model_schedules_ibfk_1` FOREIGN KEY (`model_id`) REFERENCES `model_catalogues` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `model_schedules`
--

LOCK TABLES `model_schedules` WRITE;
/*!40000 ALTER TABLE `model_schedules` DISABLE KEYS */;
INSERT INTO `model_schedules` VALUES (1,1,'2022-01-18 17:00:00','2022-01-18 12:40:55','2022-01-19 01:53:38',NULL),(2,1,'2022-01-18 17:00:00','2022-01-18 12:41:07','2022-01-19 03:05:01',NULL),(3,1,'2022-01-19 08:30:00','2022-01-18 15:09:35','2022-01-18 15:09:35',NULL),(4,2,'2022-01-18 17:00:00','2022-01-18 17:32:15','2022-01-19 03:32:53','2022-01-19 03:32:54'),(5,3,'2022-01-19 17:00:00','2022-01-19 01:59:42','2022-01-19 01:59:42',NULL),(6,3,'2022-01-20 17:00:00','2022-01-19 02:08:54','2022-01-19 02:08:54',NULL);
/*!40000 ALTER TABLE `model_schedules` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-01-19 11:08:03
