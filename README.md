# dictionary

Dictionary svr by golang.

# DB

## Tables\_in\_dictionary

```
+----------------------+
| Tables_in_dictionary |
+----------------------+
| explanations         |
| users                |
| words                |
+----------------------+
3 rows in set (0.00 sec)
```

## show create tables

```
       Table: words
Create Table: CREATE TABLE `words` (
  `wordid` int(10) NOT NULL AUTO_INCREMENT,
  `value` varchar(128) NOT NULL,
  PRIMARY KEY (`wordid`),
  UNIQUE KEY `value` (`value`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8

       Table: explanations
Create Table: CREATE TABLE `explanations` (
  `explanationid` int(10) NOT NULL AUTO_INCREMENT,
  `wordid` int(10) NOT NULL,
  `explanation` varchar(128) NOT NULL,
  `tags` varchar(128) NOT NULL DEFAULT '',
  `sentence` varchar(128) NOT NULL DEFAULT '',
  `background_img` varchar(128) NOT NULL DEFAULT '',
  `views_counter` int(10) NOT NULL DEFAULT '0',
  `thumbup_counter` int(10) NOT NULL DEFAULT '0',
  `thumbdown_counter` int(10) NOT NULL DEFAULT '0',
  `author` varchar(128) NOT NULL,
  `updatetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`explanationid`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8

       Table: users
Create Table: CREATE TABLE `users` (
  `author` varchar(128) NOT NULL,
  `avatar` varchar(128) NOT NULL DEFAULT '',
  `contribution_value` int(12) DEFAULT '0' COMMENT '贡献值 还可以增加周/月等',
  `create_words_idlist` text COMMENT '创建的词条id列表',
  `create_explanations_idlist` text COMMENT '创建的释义id列表',
  `collection_words_idlist` text COMMENT '收藏的词条id列表',
  PRIMARY KEY (`author`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
```
