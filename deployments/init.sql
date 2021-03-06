DROP DATABASE IF EXISTS egosla;

CREATE DATABASE egosla;

USE egosla;

CREATE TABLE watcher (`name` VARCHAR(191) UNIQUE, `keywords` VARCHAR(191));

CREATE TABLE subscription (
    `name` VARCHAR(191) UNIQUE, 
    `watcher` VARCHAR(191), 
    FOREIGN KEY (`watcher`) REFERENCES watcher (`name`)
);
