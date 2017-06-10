---------------------------------------------------------------------------------
--                      Copyright (C) 2017 wystan
--
--        filename: pg.sql
--     description: 
--         created: 2017-06-10 11:57:46
--          author: wystan
-- 
---------------------------------------------------------------------------------

drop database if exists bitmain;
create database bitmain;
\c bitmain

-- user for connection
drop user if exists bitmain;
create user bitmain with superuser createdb login password 'bitmain';

-- tables
create table students (
    id char(5) not null,
    classNumber int not null,
    score int not null,
    primary key(id)
);
create table classes (
    classNumber int not null,
    teacher varchar(20) not null,
    primary key(classNumber)
);

---------------------------------------------------------------------------------
