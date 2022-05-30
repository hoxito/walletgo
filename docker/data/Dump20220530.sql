-- MySQL dump 10.13  Distrib 8.0.29, for Win64 (x86_64)
--
-- Host: localhost    Database: wallet
-- ------------------------------------------------------
-- Server version	8.0.29

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
-- Table structure for table `transaction`
--

DROP TABLE IF EXISTS `transaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transaction` (
  `idtransaction` varchar(50) NOT NULL,
  `origin` varchar(50) DEFAULT NULL,
  `destination` varchar(50) DEFAULT NULL,
  `amount` float DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`idtransaction`),
  KEY `fkdestination_idx` (`destination`),
  KEY `fkorigin_idx` (`origin`),
  KEY `date_idx` (`created`),
  CONSTRAINT `fkdestination` FOREIGN KEY (`destination`) REFERENCES `wallet` (`idwallet`),
  CONSTRAINT `fkorigin` FOREIGN KEY (`origin`) REFERENCES `wallet` (`idwallet`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaction`
--

LOCK TABLES `transaction` WRITE;
/*!40000 ALTER TABLE `transaction` DISABLE KEYS */;
INSERT INTO `transaction` VALUES ('0654a038-3f04-4a27-827f-c8083abae355','testwallet1','testwallet2',0,'Created','2022-05-29 19:11:14'),('0f16bcb4-8c1a-4ef5-a379-06a15431a215','testwallet1','testwallet2',10,'Created','2022-05-30 00:46:29'),('1d65a6ce-a6dc-4ea3-90c9-f8e0f952af74','testwallet1','testwallet2',0,'Created','2022-05-29 00:00:00'),('1e11202a-3bed-4bf0-8d30-bef7c10a0e05','testwallet1','testwallet2',100,'Created','2022-05-29 19:31:55'),('30df3338-786f-4476-bb4d-887879c27bf2','testwallet1','testwallet2',10,'Created','2022-05-30 00:46:28'),('30ea6bb8-2947-48aa-8a60-1d07bd7676fb','testwallet1','testwallet2',10,'Created','2022-05-30 03:53:51'),('466273e4-f770-4c64-b28c-596019381553','testwallet1','testwallet2',0,'Created','2022-05-29 19:14:19'),('5426ca58-d448-4e6f-9d2e-097dfb86e3ad','testwallet1','testwallet2',0,'Created','2022-05-29 19:11:20'),('5c6c6d90-17a7-4ede-8ecf-01e4fe946699','testwallet1','testwallet2',100,'Created','2022-05-29 19:52:34'),('612e99ad-df80-4818-a49b-5d03a6c69cce','testwallet1','testwallet2',10,'Created','2022-05-30 00:46:26'),('b800b494-5321-4779-8522-e5e5ee17b96b','testwallet1','testwallet2',700,'Created','2022-05-29 19:52:57'),('deb1aaec-ef51-4510-8a71-6159457c6db9','testwallet1','testwallet2',100,'Created','2022-05-29 19:52:05'),('dfb26608-7664-4e2f-9a67-900f9982beac','testwallet1','testwallet2',0,'Created','2022-05-29 19:14:53'),('ee898e05-bc07-428b-a3f5-1e97bccfadee','testwallet1','testwallet3',10,'Created','2022-05-30 03:53:56'),('f21d3546-60e8-413c-8619-50eb14dcbd61','testwallet1','testwallet2',10,'Created','2022-05-30 00:46:22'),('fdd5bec5-eceb-4b4c-a026-cb8fc6953e7b','testwallet1','testwallet2',10,'Created','2022-05-30 00:46:27'),('tiaad4','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tiaadd4','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tiaadsaad4','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tiaadsaafd4','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tiaadsaafsad4','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tiaadsaafsasd4','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tiaadsaafsasdad4','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tiaadsaafsasdasd4','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tiaadsaafsasdd4','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tiaadsaafsd4','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tiaadsad4','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tiaadsd4','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tid1','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tid12','testwallet1','testwallet2',100,NULL,'2022-05-28 00:00:00'),('tid162','testwallet1','testwallet2',100,NULL,'2022-05-28 00:00:00'),('tid16212','testwallet3','testwallet2',100,NULL,'2022-05-30 00:00:00'),('tid2','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tid3','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00'),('tid4','testwallet1','testwallet2',100,NULL,'2022-05-26 00:00:00');
/*!40000 ALTER TABLE `transaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `iduser` varchar(50) NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `created` date DEFAULT NULL,
  `deleted` date DEFAULT NULL,
  PRIMARY KEY (`iduser`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES ('testuser1','jose','josearanciba09@gmail.com','jose1','2022-05-25',NULL),('testuser2','santiago','santiago@gmail.com','santiago1','2022-05-25',NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wallet`
--

DROP TABLE IF EXISTS `wallet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wallet` (
  `idwallet` varchar(50) NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `iduser` varchar(100) DEFAULT NULL,
  `ballance` float DEFAULT NULL,
  `currency` varchar(45) DEFAULT NULL,
  `created` date DEFAULT NULL,
  `deleted` date DEFAULT NULL,
  `updated` date DEFAULT NULL,
  PRIMARY KEY (`idwallet`),
  KEY `fkuser_idx` (`iduser`),
  CONSTRAINT `fkuser` FOREIGN KEY (`iduser`) REFERENCES `user` (`iduser`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wallet`
--

LOCK TABLES `wallet` WRITE;
/*!40000 ALTER TABLE `wallet` DISABLE KEYS */;
INSERT INTO `wallet` VALUES ('testwallet1','jose wallet','testuser1',930,'USD','2022-05-25',NULL,NULL),('testwallet2','santy wallet','testuser2',3960,'USD','2022-05-25',NULL,NULL),('testwallet3','miguel wallet','testuser2',1910,'USD','2022-05-25',NULL,NULL);
/*!40000 ALTER TABLE `wallet` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-30  1:22:53
