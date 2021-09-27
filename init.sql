CREATE DATABASE IF NOT EXISTS go_restful_api;
USE go_restful_api;
CREATE TABLE category
(
    id   integer primary key auto_increment,
    name varchar(200) not null
) engine = InnoDB;