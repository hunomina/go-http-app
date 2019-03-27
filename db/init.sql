CREATE DATABASE IF NOT EXISTS go_http_app;

USE go_http_app;

CREATE TABLE IF NOT EXISTS user (
  id int(11) not null AUTO_INCREMENT,
  name varchar(50) not null,
  email varchar(100) not null,
  password varchar(100) not null,
  age int(3) default 0,
  PRIMARY KEY (id)
)