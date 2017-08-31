/*
Navicat MySQL Data Transfer

Source Server         : 1
Source Server Version : 50525
Source Host           : localhost:3306
Source Database       : _abito

Target Server Type    : MYSQL
Target Server Version : 50525
File Encoding         : 65001

Date: 2016-11-23 19:37:25
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of categories
-- ----------------------------
INSERT INTO `categories` VALUES ('1', 'auto');
INSERT INTO `categories` VALUES ('2', 'girls');

-- ----------------------------
-- Table structure for cookies
-- ----------------------------
DROP TABLE IF EXISTS `cookies`;
CREATE TABLE `cookies` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cookie` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of cookies
-- ----------------------------
INSERT INTO `cookies` VALUES ('1', 'VQXHfDjUhsKaAwrB');
INSERT INTO `cookies` VALUES ('2', 'eMLMpcBkZrdzDvAE');
INSERT INTO `cookies` VALUES ('3', 'iTDwfCrJHnmpaBfH');
INSERT INTO `cookies` VALUES ('4', 'VQrKbCElWwQcStqw');
INSERT INTO `cookies` VALUES ('5', 'zkoGWYTESzEDZfRi');

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cookieid` int(11) DEFAULT NULL,
  `postingid` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of favorites
-- ----------------------------

-- ----------------------------
-- Table structure for postings
-- ----------------------------
DROP TABLE IF EXISTS `postings`;
CREATE TABLE `postings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `categoryId` int(11) DEFAULT NULL,
  `cookieid` int(11) DEFAULT NULL,
  `caption` text,
  `content` text,
  `phonenumber` text,
  `created` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of postings
-- ----------------------------
INSERT INTO `postings` VALUES ('2', '1', '4', 'bmw', 'E36', 'старая', '9999');
INSERT INTO `postings` VALUES ('3', '1', '4', 'sky', 'nissian', 'ss', '9999');
SET FOREIGN_KEY_CHECKS=1;
