create database if not exists rede_social;
use rede_social;

drop table if exists users;

create table users (
    id int auto_increment primary key,
    `name` varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(100) not null unique,
    `password` varchar(200) not null,
    createdat timestamp default current_timestamp()
) ENGINE=INNODB;