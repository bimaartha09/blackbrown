CREATE DATABASE bayarxyz_development;

CREATE TABLE `bayarxyz_development`.`payment` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `payment_type` INT NOT NULL DEFAULT 0,
  `amount` INT NOT NULL DEFAULT 0,
  `discount_amount` INT NOT NULL DEFAULT 0,
  `source_id` VARCHAR(30) NULL,
  `description` VARCHAR(500) NULL,
  `created_at` INT NOT NULL,
  `updated_at` INT NOT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `bayarxyz_development`.`virtual_account` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NULL,
  `code` VARCHAR(20) NULL,
  `is_used` INT NOT NULL DEFAULT 0,
  `created_at` INT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `bayarxyz_development`.`coupon` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  `description` VARCHAR(500) NULL,
  `percentage` INT NOT NULL DEFAULT 0,
  `amount` INT NOT NULL DEFAULT 0,
  `expire_time` INT NULL DEFAULT 0,
  `created_at` INT NULL DEFAULT 0,
  `updated_at` INT NULL DEFAULT 0,
  PRIMARY KEY (`id`));

INSERT INTO `bayarxyz_development`.`coupon` (`name`, `description`, `percentage`) VALUES ('Diskon 5%', 'Get Discount 5%', '5');
INSERT INTO `bayarxyz_development`.`coupon` (`name`, `description`, `amount`) VALUES ('Diskon 5 ribu', 'Get Discount Rp 5000', '5000');
