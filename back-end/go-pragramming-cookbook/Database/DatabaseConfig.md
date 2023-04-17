# 데이터 베이스 설정

### MYSQL Settings
```
create database golangTestDB default character set utf8;
CREATE USER 'gopher'@'localhost' IDENTIFIED BY 'gopher';
grant all privileges on golangTestDB.* to 'gopher'@'localhost';
flush privileges;
```
