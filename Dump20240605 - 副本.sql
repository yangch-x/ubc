
CREATE TABLE `Config` (
  `config_id` int(11) NOT NULL,
  `ubc_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `ubc_address` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `qimei_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `qimei_address` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `qimei_taxno` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `qimei_compno` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `rmb_rate` double NOT NULL,
  `euro_rate` double DEFAULT NULL,
  `jpy_rate` double DEFAULT NULL,
  `gbp_rate` double DEFAULT NULL,
  PRIMARY KEY (`ubc_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `ConfigHist` (
  `chg_id` int(11) NOT NULL AUTO_INCREMENT,
  `chg_dt` date NOT NULL,
  `config_id` int(11) NOT NULL,
  `ubc_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `ubc_address` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `qimei_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `qimei_address` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `qimei_taxno` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `qimei_compno` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `rmb_rate` double NOT NULL,
  `euro_rate` double DEFAULT NULL,
  `jpy_rate` double DEFAULT NULL,
  `gbp_rate` double DEFAULT NULL,
  PRIMARY KEY (`chg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `Customer` (
  `customer_id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_code` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `customer_email` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `customer_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `billing_contact` text COLLATE utf8_unicode_ci,
  `notify_contact` text COLLATE utf8_unicode_ci,
  `payment_term` varchar(25) COLLATE utf8_unicode_ci DEFAULT NULL,
  `ship_to` text COLLATE utf8_unicode_ci,
  `sales_person` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `ubc_merchandiser` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `country` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `discharge_loc` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `status` varchar(25) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`customer_id`),
  UNIQUE KEY `customer_code` (`customer_code`)
) ENGINE=InnoDB AUTO_INCREMENT=2401 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `HSCode` (
  `hs_id` int(11) NOT NULL AUTO_INCREMENT,
  `hs_code` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `hts_code` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `descr_en` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `descr_cn` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `custom_factors` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `notes` text COLLATE utf8_unicode_ci,
  PRIMARY KEY (`hs_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `HSCodeHist` (
  `chg_id` int(11) NOT NULL AUTO_INCREMENT,
  `chg_dt` date NOT NULL,
  `hs_code` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `hts_code` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `descr_en` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `descr_cn` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `custom_factors` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `notes` text COLLATE utf8_unicode_ci,
  PRIMARY KEY (`chg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `Invoice` (
  `invoice_id` int(11) NOT NULL AUTO_INCREMENT,
  `ship_id` int(11) DEFAULT NULL,
  `invoice_code` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `ubc_pi` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `customer_code` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `invoice_amt` double DEFAULT NULL,
  `received_amt` double DEFAULT NULL,
  `invoice_dt` date DEFAULT NULL,
  `invoice_due` date DEFAULT NULL,
  `invoice_currency` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT 'USD',
  `notes` text COLLATE utf8_unicode_ci,
  PRIMARY KEY (`invoice_id`),
  UNIQUE KEY `inv_cd` (`invoice_code`)
) ENGINE=InnoDB AUTO_INCREMENT=401185 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


CREATE TABLE `PO` (
  `po_id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_po` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `style_code` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `style_color` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `style_size` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `order_dt` date DEFAULT NULL,
  `order_qty` int(11) DEFAULT NULL,
  `ship_qty` int(11) DEFAULT NULL,
  `unit_price` double DEFAULT NULL,
  `sales_price` double DEFAULT NULL,
  `custom_price` double DEFAULT NULL,
  `cost_price` double DEFAULT NULL,
  `notes` text COLLATE utf8_unicode_ci,
  PRIMARY KEY (`po_id`),
  UNIQUE KEY `cust_po` (`customer_po`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `Packing` (
  `pack_id` int(11) NOT NULL AUTO_INCREMENT,
  `ship_id` int(11) DEFAULT NULL,
  `invoice_id` int(11) DEFAULT NULL,
  `pack_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `style_code` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `po_code` varchar(250) COLLATE utf8_unicode_ci NOT NULL,
  `carton_code` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `begin_num` int(11) NOT NULL,
  `end_num` int(11) NOT NULL,
  `carton_cnt` int(11) NOT NULL,
  `pack_color` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `pack_size` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `item_cnt` int(11) DEFAULT NULL,
  `gross_weight` double DEFAULT NULL,
  `net_weight` double DEFAULT NULL,
  `carton_size` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `line_cnt` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`pack_id`),
  UNIQUE KEY `pack_key` (`style_code`,`pack_name`,`carton_code`,`pack_color`,`pack_size`)
) ENGINE=InnoDB AUTO_INCREMENT=800115 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `PackingList` (
  `list_id` int(11) NOT NULL AUTO_INCREMENT,
  `ship_id` int(11) NOT NULL,
  `proj_id` int(11) DEFAULT NULL,
  `pack_name` varchar(255) NOT NULL,
  `carton_cnt` int(11) NOT NULL,
  `item_cnt` int(11) DEFAULT NULL,
  `meas_vol` double DEFAULT NULL,
  `gross_weight` double DEFAULT NULL,
  `net_weight` double DEFAULT NULL,
  `carton_size` varchar(255) DEFAULT NULL,
  `pack_cnt` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`list_id`)
) ENGINE=InnoDB AUTO_INCREMENT=900421 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `Projection` (
  `proj_id` int(11) NOT NULL AUTO_INCREMENT,
  `arrive_dt` date NOT NULL,
  `ubc_pi` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `fob_ldp` varchar(25) COLLATE utf8_unicode_ci NOT NULL,
  `customer_code` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `country` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `customer_po` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `master_po` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `style_code` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `style_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `fabrication` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `color` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `size` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `po_qty` int(11) DEFAULT NULL,
  `ship_qty` int(11) DEFAULT NULL,
  `sale_price` double DEFAULT NULL,
  `sale_cust_price` double DEFAULT NULL,
  `sale_currency` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT 'USD',
  `invoice_code` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `receiving` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `notes` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `cost_price` double DEFAULT NULL,
  `cost_currency` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT 'RMB',
  `rmb_inv` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `exporter` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `ubc_payable` double DEFAULT NULL,
  `pay_period` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `sales_person` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `sales_commission` double DEFAULT NULL,
  `comm_paid` double DEFAULT NULL,
  PRIMARY KEY (`proj_id`)
) ENGINE=InnoDB AUTO_INCREMENT=502000 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `Shipment` (
  `ship_id` int(11) NOT NULL AUTO_INCREMENT,
  `rmb_inv` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `master_po` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `customer_code` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `ubc_pi` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `markurl` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `orig_country` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `ship_method` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `ship_term` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `invoice_ttl` double DEFAULT NULL,
  `ship_from` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `master_bl_num` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `house_bl_num` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `exporter` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `ship_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `pack_dt` date DEFAULT NULL,
  `ship_dt` date DEFAULT NULL,
  `arrive_dt` date DEFAULT NULL,
  `notes` text COLLATE utf8_unicode_ci,
  PRIMARY KEY (`ship_id`)
) ENGINE=InnoDB AUTO_INCREMENT=300128 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `Style` (
  `style_id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_code` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `style_code` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `style_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `size_type` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT 'S-M-L',
  `fabrication` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `rmb_price` double DEFAULT NULL,
  `hs_id` int(11) DEFAULT NULL,
  `notes` text COLLATE utf8_unicode_ci,
  PRIMARY KEY (`style_id`),
  UNIQUE KEY `style_cd` (`customer_code`,`style_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `first_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `last_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `role` enum('admin','ch_user','us_user','user') COLLATE utf8_unicode_ci DEFAULT 'user',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

