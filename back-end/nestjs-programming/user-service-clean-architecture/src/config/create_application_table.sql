CREATE DATABASE nestJsBackendProgramming CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'nestJsBackendProgrammingUser'@'localhost' IDENTIFIED BY 'nestJsBackendProgrammingPassword';
GRANT all privileges on nestJsBackendProgramming.* to 'nestJsBackendProgrammingUser'@'localhost';
flush privileges;