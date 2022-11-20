/*
SQLyog Community v13.1.6 (64 bit)
MySQL - 5.7.19 : Database - anygo
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`anygo` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `anygo`;

/*Data for the table `format_converts` */

insert  into `format_converts`(`id`,`in_format`,`out_format`,`convert_util`) values 
(6,'doc','docx',1),
(7,'doc','html',1),
(5,'doc','pdf',1),
(8,'doc','txt',1),
(2,'docx','doc',1),
(4,'docx','html',1),
(1,'docx','pdf',1),
(3,'docx','txt',1),
(9,'pdf','docx',2);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
