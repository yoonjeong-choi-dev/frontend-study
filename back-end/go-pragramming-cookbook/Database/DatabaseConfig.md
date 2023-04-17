# 데이터 베이스 설정

### MYSQL Settings
```
create database golangTestDB default character set utf8;
CREATE USER 'gopher'@'localhost' IDENTIFIED BY 'gopher';
grant all privileges on golangTestDB.* to 'gopher'@'localhost';
flush privileges;
```

### Redis Config
* config file: /opt/homebrew/etc/redis.conf
    * 비밀번호 설정 필드: requirepass <password>
    * 현재 비밀번호 설정: requirepass localredispassword
* `redis-server /opt/homebrew/etc/redis.conf`
    * 서버 실행
* client 접속
    * `redis-cli -a localredispassword`
    * monitoring: `redis-cli -a localredispassword monitor`
