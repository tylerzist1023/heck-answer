DROP DATABASE IF EXISTS ANSWERHECK;
CREATE DATABASE ANSWERHECK;

USE ANSWERHECK;

DROP TABLE IF EXISTS USER;
CREATE TABLE USER (
    username VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL,
    PRIMARY KEY (`username`)
);

DROP TABLE IF EXISTS SESSION;
CREATE TABLE SESSION (
    cookie VARCHAR(64) NOT NULL,
    username VARCHAR(64) NOT NULL,
    PRIMARY KEY (`cookie`)
);