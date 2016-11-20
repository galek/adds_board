/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50525
Source Host           : localhost:3306
Source Database       : _abito

Target Server Type    : MYSQL
Target Server Version : 50525
File Encoding         : 65001

Date: 2016-11-21 01:49:37
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` int(11) NOT NULL,
  `name` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of categories
-- ----------------------------
INSERT INTO `categories` VALUES ('1', 'auto');
INSERT INTO `categories` VALUES ('2', 'girls');

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites` (
  `id` int(11) NOT NULL,
  `cookie` text,
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
  `id` int(11) NOT NULL,
  `categoryId` int(11) DEFAULT NULL,
  `cookie` text,
  `caption` text,
  `content` text,
  `phonenumber` text,
  `created` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of postings
-- ----------------------------
INSERT INTO `postings` VALUES ('0', '1', '0', 'BMW M3 III (E46) 2003', 'Авто новое в идеальном состояние\r\nСделана на заказ', '7 963 710-45-90', '0');
INSERT INTO `postings` VALUES ('1', '1', '0', 'Nissan Skyline X (R34)', 'Торг, о цене договоримся, ОЧЕНЬ СРОЧНО надо погасить кредит!\r\nАтмосферный двигатель;\r\nДвигатель в идеальном состоянии, работает ровно, без нареканий;\r\nДвигатель Масло не ест;\r\nДефекты кузова на фото;\r\nРасходники менялись вовремя;\r\nЗамена масла в двигателе производилась каждые 5000 км.;\r\nКоробка переключает плавно без рывков;\r\n3000 км. назад была замена ремня ГРМ+ролики, колодки новые, рычаги и по мелочи;\r\nХодовая частично перебиралась;\r\nЭлектрика - работает без нареканий;\r\nСигнализация с обратной связью и автозапуском (очень удобно зимой);\r\nШтатный ксенон с корректором ближнего света;\r\nДатчик заднего хода(парктроник);\r\nЭлектропривод зеркал;\r\nТонированная вся, кроме лобового стекла;\r\nСтекла все родные;\r\nПодогрев заднего стекла и боковых зеркал;\r\nЧистый салон;\r\nКондиционер работает;\r\nРезина перед 225-35R19, зад 255-55R19;\r\nЯ хозяин уже 7 лет;\r\nСкай специально пригоняли на заказ из Владивостока;\r\nОбмен не интересует;\r\nПродавцов из автосалонов и перекупщиков просьба не беспокоить, \r\nне теряйте своё и моё драгоценное время; ', '7 916 267-86-82', '20111');
INSERT INTO `postings` VALUES ('2', '2', '0', 'test2', 'test2', '8', '2012');
SET FOREIGN_KEY_CHECKS=1;
