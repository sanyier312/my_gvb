-- MySQL dump 10.13  Distrib 8.0.32, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: gvb_db
-- ------------------------------------------------------
-- Server version	8.0.32

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `article`
--

DROP TABLE IF EXISTS `article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `article` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(100) NOT NULL,
  `cid` bigint unsigned NOT NULL,
  `desc` varchar(200) DEFAULT NULL,
  `content` longtext,
  `img` varchar(100) DEFAULT NULL,
  `comment_count` bigint NOT NULL DEFAULT '0',
  `read_count` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_article_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article`
--

LOCK TABLES `article` WRITE;
/*!40000 ALTER TABLE `article` DISABLE KEYS */;
INSERT INTO `article` VALUES (1,'2023-04-28 20:00:35.851','2023-04-28 20:00:35.851',NULL,'分类一下的第一篇文章',1,'分类一下的第一篇文章','分类一下的第三篇文章','http://rtkf04ah5.hd-bkt.clouddn.com/FjdgrLyY-blxqZN7HqwtMkGuL-86',0,1),(2,'2023-04-28 20:04:06.581','2023-04-28 20:04:06.581',NULL,'分类一下的第二篇文章',1,'分类二下的第二篇文章','分类一下的第二篇文章','http://rtkf04ah5.hd-bkt.clouddn.com/FjdgrLyY-blxqZN7HqwtMkGuL-86',0,2),(3,'2023-04-28 20:04:35.192','2023-04-28 20:04:35.192',NULL,'分类二下的第一篇文章',2,'分类二下的第一篇文章','分类二下的第一篇文章','http://rtkf04ah5.hd-bkt.clouddn.com/FpozmVW2PmxtgaqvzTDOSf0hDTLe',0,1),(4,'2023-04-28 20:04:46.382','2023-04-28 20:04:46.382',NULL,'分类三下的第一篇文章',3,'分类三下的第一篇文章','分类三下的第一篇文章','http://rtkf04ah5.hd-bkt.clouddn.com/FjdgrLyY-blxqZN7HqwtMkGuL-86',0,2),(5,'2023-04-28 20:04:56.344','2023-04-28 20:04:56.344',NULL,'分类四下的第一篇文章',4,'分类四下的第一篇文章','分类四下的第一篇文章','http://rtkf04ah5.hd-bkt.clouddn.com/FpozmVW2PmxtgaqvzTDOSf0hDTLe',0,1),(6,'2023-04-28 20:05:08.680','2023-04-28 20:05:08.680',NULL,'分类五下的第一篇文章',5,'分类五下的第一篇文章','分类五下的第一篇文章','http://rtkf04ah5.hd-bkt.clouddn.com/FjdgrLyY-blxqZN7HqwtMkGuL-86',0,5),(7,'2023-04-28 20:09:16.700','2023-04-28 20:09:16.700',NULL,'分类六的第一篇文章',6,'分类六的第一篇文章','<p>分类六的第一篇文章</p>','http://rtkf04ah5.hd-bkt.clouddn.com/FjdgrLyY-blxqZN7HqwtMkGuL-86',0,6),(8,'2023-04-28 20:28:02.705','2023-04-28 20:28:02.705',NULL,'分类七的第一篇文章',7,'分类七的第一篇文章','<p>分类七的第一篇文章</p>','http://rtkf04ah5.hd-bkt.clouddn.com/FnemSb2czi9mGawfG9zZ91sQXsBr',1,10),(9,'2023-05-02 08:47:46.168','2023-05-02 08:47:46.168','2023-05-02 10:01:17.410','测试文章',7,'测试','<p>测试</p>','http://rtkf04ah5.hd-bkt.clouddn.com/FjdgrLyY-blxqZN7HqwtMkGuL-86',0,0),(10,'2023-05-02 09:10:30.688','2023-05-02 09:10:30.688',NULL,'测试萨达',4,'鼎折覆餗','<p>大厦是大法官梵蒂冈</p>','http://rtkf04ah5.hd-bkt.clouddn.com/Ft2aXSjcBIDw0ysFqgvXa0Pafknk',1,9),(11,'2023-05-02 09:58:41.301','2023-05-02 09:58:41.301','2023-05-02 09:59:29.466','scsdfcs',4,'sdfsdfsdfs','<p>vdgvdgdfgvd</p>','',0,0),(12,'2023-05-02 10:00:38.318','2023-05-04 21:32:36.508',NULL,'dfgdfgd',5,'dfhdfh钢结构','<p>gfhfghfghfg</p>','http://rtkf04ah5.hd-bkt.clouddn.com/FksQv5ItZaAdkH2syBFni1obefbk',1,8),(13,'2023-05-04 21:33:27.587','2023-05-04 21:33:27.587','2023-05-04 21:33:34.231','阿达萨达',10,'asdasd','<p>asdasd</p>','',0,0),(14,'2023-05-04 22:28:29.679','2023-05-06 08:21:50.323',NULL,'njininn',5,'uiohouinini','<p>iuhihuiuhuhuh不错吧fgfjh</p>','http://rtkf04ah5.hd-bkt.clouddn.com/Ft2aXSjcBIDw0ysFqgvXa0Pafknk',0,11),(15,'2023-05-06 08:24:08.359','2023-05-06 08:58:59.513',NULL,'test1',5,'神鼎飞丹砂fdsf','<p>水电费hjk</p>','http://rtkf04ah5.hd-bkt.clouddn.com/Ft2aXSjcBIDw0ysFqgvXa0Pafknk',0,10),(16,'2023-05-06 08:59:28.003','2023-05-06 09:23:59.469',NULL,'gfg',9,'gdfgdfghfhg','<p>dgdfgdhfghf</p>','http://rtkf04ah5.hd-bkt.clouddn.com/FnemSb2czi9mGawfG9zZ91sQXsBr',0,3),(17,'2023-05-06 11:20:59.295','2023-05-06 12:07:39.217',NULL,'fxhbhfh',2,'dfgdfgd','<p>dfgdfgdfgdfgd</p>','http://rtkf04ah5.hd-bkt.clouddn.com/FnemSb2czi9mGawfG9zZ91sQXsBr',0,6),(18,'2023-05-06 12:17:08.025','2023-05-06 12:17:08.025','2023-05-06 12:17:14.650','打算',7,'大厦','<p>阿斯顿撒旦</p>','',0,0),(19,'2023-05-06 12:17:27.633','2023-05-06 12:17:27.633','2023-05-07 18:09:32.204','啊实打实',1,'大厦','<p>asdasd</p>','',0,0),(20,'2023-05-07 18:12:01.039','2023-05-07 18:12:07.325','2023-05-07 18:12:10.499','fsdfsf',7,'sdfsdf','<p>sdfsdfsfs</p>','http://rtkf04ah5.hd-bkt.clouddn.com/Ft2aXSjcBIDw0ysFqgvXa0Pafknk',0,1);
/*!40000 ALTER TABLE `article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1,'分类一'),(2,'分类二'),(3,'分类三'),(4,'分类四'),(5,'分类五'),(6,'分类六'),(7,'分类七'),(9,'测试');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `article_id` bigint unsigned DEFAULT NULL,
  `title` longtext,
  `username` longtext,
  `content` varchar(500) NOT NULL,
  `status` tinyint DEFAULT '2',
  PRIMARY KEY (`id`),
  KEY `idx_comment_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (1,'2023-04-28 20:35:19.654','2023-04-28 20:52:45.613','2023-04-28 20:52:47.863',0,0,'','张三','第一条评论',1),(2,'2023-04-28 20:53:05.629','2023-04-28 20:53:14.361','2023-05-06 09:23:43.975',0,0,'','张三','第一条评论',1),(3,'2023-05-02 10:14:42.327','2023-05-02 10:18:34.759',NULL,7,12,'','dddd','评论测试',1),(4,'2023-05-02 10:28:35.012','2023-05-06 09:23:37.053',NULL,3,10,'','sanyier','哈哈哈',1),(5,'2023-05-02 10:28:53.924','2023-05-06 08:22:01.448',NULL,3,10,'','sanyier','哈哈哈哈哈哈哈',1),(6,'2023-05-05 18:18:56.950','2023-05-06 08:21:59.070',NULL,1,8,'','admin','菜市场发送到',1),(7,'2023-05-08 15:51:28.472','2023-05-08 15:51:28.472',NULL,1,7,'','admin','交电费发送到积分',2);
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `profile`
--

DROP TABLE IF EXISTS `profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `profile` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL,
  `desc` varchar(200) DEFAULT NULL,
  `qqchat` varchar(200) DEFAULT NULL,
  `wechat` varchar(100) DEFAULT NULL,
  `weibo` varchar(200) DEFAULT NULL,
  `bili` varchar(200) DEFAULT NULL,
  `email` varchar(200) DEFAULT NULL,
  `img` varchar(200) DEFAULT NULL,
  `avatar` varchar(200) DEFAULT NULL,
  `icp_record` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `profile`
--

LOCK TABLES `profile` WRITE;
/*!40000 ALTER TABLE `profile` DISABLE KEYS */;
INSERT INTO `profile` VALUES (1,'Qiu&&Xue','skdjsd','4545','454544','4545','4545','454454','http://rtkf04ah5.hd-bkt.clouddn.com/FjdgrLyY-blxqZN7HqwtMkGuL-86','http://rtkf04ah5.hd-bkt.clouddn.com/Ft2aXSjcBIDw0ysFqgvXa0Pafknk','66666');
/*!40000 ALTER TABLE `profile` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(20) NOT NULL,
  `password` varchar(500) NOT NULL,
  `role` bigint DEFAULT '2',
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'2023-04-28 19:53:28.539','2023-05-09 22:14:12.071',NULL,'admin','$2a$10$ujTETZ2UenHhlKf1zleUXuWeJSem/iqyQDb8/ZnpuE7lKEXq5iwZG',1),(2,'2023-04-28 19:53:54.071','2023-05-07 18:12:54.672',NULL,'1234','$2a$10$iCubkB33q6R0Pv8nukP5zu.vmJUZTwpenypPQsVwF9t9/DuBJ6raO',2),(3,'2023-04-28 19:53:59.023','2023-04-28 19:53:59.023',NULL,'sanyier','$2a$10$8CYUgAfAKTvy4Q1/zERklegubo6PtsEM/qubLBR9xHWOELHcu6HXW',2),(4,'2023-04-28 19:54:08.848','2023-04-28 19:54:08.848','2023-05-06 08:48:07.287','ferry','$2a$10$Q2BxNlv27ED6J/MlzHvRnOXUb9ZiEIYLKokufMq6.IRirzKUbjhxa',2),(5,'2023-05-02 08:39:51.822','2023-05-02 08:39:51.822',NULL,'卤煮1234','$2a$10$t1sq4xY9Id9QBhHMIBmsRO99EIdBt//Shk3c2JYlp11iQ1HwRZ/PS',2),(6,'2023-05-02 08:41:40.009','2023-05-02 08:41:40.009','2023-05-06 08:50:50.806','哈哈哈哈','$2a$10$XEHuK2s8zXT./U9K.h4s3OcEI/W05FblTzFnu3yNzjTGqnyYGeYia',2),(7,'2023-05-02 10:09:21.233','2023-05-02 10:09:21.233','2023-05-06 09:25:45.799','dddd','$2a$10$sLMX./zDtbjirlT.aA3lFuQx1.67jErz1tDceAj0amO49N.jIabAq',2),(8,'2023-05-02 17:20:09.440','2023-05-02 17:20:09.440','2023-05-06 08:45:25.142','jnknkjnj','$2a$10$csoxgfedLmVi8e9MxKyszuTuqGUdmp4M4GjC4cHQg1ng.a3acnuBy',2);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-05-11 17:08:03
