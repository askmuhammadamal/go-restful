CREATE DATABASE IF NOT EXISTS go_restful_api_test;
USE go_restful_api_test;
CREATE TABLE category
(
    id   integer primary key auto_increment,
    name varchar(200) not null
) engine = InnoDB;