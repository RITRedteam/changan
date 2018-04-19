CREATE DATABASE IF NOT EXISTS changan_test DEFAULT CHARACTER SET 'utf8'
DROP TABLE IF EXISTS interfaces
DROP TABLE IF EXISTS ips
DROP TABLE IF EXISTS reports
DROP TABLE IF EXISTS users
DROP TABLE IF EXISTS subnets
DROP TABLE IF EXISTS devices
DROP TABLE IF EXISTS sessions
CREATE TABLE IF NOT EXISTS devices (
    device_id int(5) NOT NULL AUTO_INCREMENT,
    device_name varchar(50) DEFAULT NULL,
    team varchar(20) DEFAULT NULL,
    owner varchar(30) DEFAULT NULL,
    location varchar(250) DEFAULT NULL,
    PRIMARY KEY(device_id)
    );
CREATE TABLE IF NOT EXISTS subnets (
    subnet_id int(5) NOT NULL AUTO_INCREMENT,
    subnet_name varchar(50) DEFAULT NULL,
    ip varchar(55) DEFAULT NULL,
    mask int(5) DEFAULT NULL,
    PRIMARY KEY(subnet_id)
    );
CREATE TABLE IF NOT EXISTS ips (
    ip_id int(5) NOT NULL AUTO_INCREMENT,
    subnet_id int(5) NOT NULL,
    ip varchar(55) DEFAULT NULL,
    PRIMARY KEY(ip_id),
    FOREIGN KEY(subnet_id) REFERENCES subnets(subnet_id)
    );
CREATE TABLE IF NOT EXISTS interfaces (
    device_id int(5) NOT NULL,
    ip_id int(5) NOT NULL,
    mac varchar(17) NOT NULL,
    interface_name varchar(10) DEFAULT NULL,
    PRIMARY KEY(device_id, ip_id),
    FOREIGN KEY(device_id) REFERENCES devices(device_id),
    FOREIGN KEY(ip_id) REFERENCES ips(ip_id)
    );
CREATE TABLE IF NOT EXISTS users (
    username varchar(50) NOT NULL,
    user_id int(5) NOT NULL AUTO_INCREMENT,
    password char(60) NOT NULL,
    api_key varchar(50) DEFAULT NULL,
    active BOOL,
    PRIMARY KEY(user_id),
    UNIQUE KEY(username)
    );
CREATE TABLE IF NOT EXISTS reports (
    device_id int(5) NOT NULL,
    report varchar(3000) DEFAULT NULL,
    last_user_id int(5) DEFAULT NULL,
    FOREIGN KEY(device_id) REFERENCES devices(device_id),
    FOREIGN KEY(last_user_id) REFERENCES users(user_id)
    );
CREATE TABLE sessions (
  token CHAR(43) PRIMARY KEY,
  data BLOB NOT NULL,
  expiry TIMESTAMP(6) NOT NULL
);
CREATE INDEX sessions_expiry_idx ON sessions (expiry);
INSERT INTO devices (device_id, device_name, team, owner, location) values(1, 'HTTP', 'OPS', 'ksam', 'china');
INSERT INTO devices (device_id, device_name, team, owner, location) values(2, 'DNS', 'OPS', 'ksam', 'china');
INSERT INTO devices (device_id, device_name, team, owner, location) values(3, 'SQL', 'OPS', 'jnottingham', 'russia');
INSERT INTO devices (device_id, device_name, team, owner, location) values(4, 'Router', 'Networking', 'jgem', 'LA');
INSERT INTO devices (device_id, device_name, team, owner, location) values(5, 'AD', 'SysAd', 'bwhite', 'laptop');
INSERT INTO devices (device_id, device_name, team, owner, location) values(6, 'Jenkins oh god', 'Engineering', 'rgeorge', 'LA');
INSERT INTO devices (device_id, device_name, team, owner, location) values(7, 'gitlab', 'Engineering', 'rgeorge', 'LA');
INSERT INTO devices (device_id, device_name, team, owner, location) values(8, 'esxi', 'OPS', 'ksam', 'china');
INSERT INTO subnets (subnet_id, subnet_name, ip, mask) values(1, 'china office', '172.16.9.0', 24);
INSERT INTO subnets (subnet_id, subnet_name, ip, mask) values(2, 'russia office', '172.16.10.0', 24);
INSERT INTO subnets (subnet_id, subnet_name, ip, mask) values(3, 'LA office', '172.16.11.0', 24);
INSERT INTO subnets (subnet_id, subnet_name, ip, mask) values(4, 'Prod', '10.100.0.0', 16);
INSERT INTO ips (ip_id, ip, subnet_id) values(1, '10.100.100.200', 4);
INSERT INTO ips (ip_id, ip, subnet_id) values(2, '10.100.100.201', 4);
INSERT INTO ips (ip_id, ip, subnet_id) values(3, '10.100.100.202', 4);
INSERT INTO ips (ip_id, ip, subnet_id) values(4, '10.100.100.100', 4);
INSERT INTO ips (ip_id, ip, subnet_id) values(5, '172.16.11.203', 3);
INSERT INTO ips (ip_id, ip, subnet_id) values(6, '172.16.11.124', 3);
INSERT INTO ips (ip_id, ip, subnet_id) values(7, '172.16.10.1', 2);
INSERT INTO ips (ip_id, ip, subnet_id) values(8, '172.16.9.1', 1);
INSERT INTO ips (ip_id, ip, subnet_id) values(9, '172.16.9.53', 1);
INSERT INTO ips (ip_id, ip, subnet_id) values(10, '172.16.10.101', 2);
INSERT INTO interfaces (device_id, ip_id, mac, interface_name) values(1, 8, 'AA:BB:CC:DD:EE:FF', 'eth0');
INSERT INTO interfaces (device_id, ip_id, mac, interface_name) values(2, 9, 'AA:BB:CC:DD:EE:EE', 'eth0');
INSERT INTO interfaces (device_id, ip_id, mac, interface_name) values(3, 7, 'AA:BB:CC:DD:EE:DD', 'eth0');
INSERT INTO interfaces (device_id, ip_id, mac, interface_name) values(4, 6, 'AA:BB:CC:DD:EE:CC', 'eth0');
INSERT INTO interfaces (device_id, ip_id, mac, interface_name) values(5, 10, 'AA:BB:CC:DD:EE:BB', 'eth0');
INSERT INTO interfaces (device_id, ip_id, mac, interface_name) values(6, 1, 'AA:BB:CC:DD:EE:AA', 'eth0');
INSERT INTO interfaces (device_id, ip_id, mac, interface_name) values(7, 2, 'AA:BB:CC:DD:FF:FF', 'eth0');
INSERT INTO interfaces (device_id, ip_id, mac, interface_name) values(8, 3, 'AA:BB:CC:DD:FF:EE', 'eth0');
INSERT INTO users (user_id, username, password, api_key, active) values(1, 'rwhittier', '$2a$12$9TZ3IJjTaYIB9Yrct5YT.ey7DGsybJR6d8nMb.Q8coFDKwqUMTCju', 'aaa', True);
INSERT INTO users (user_id, username, password, api_key, active) values(2, 'test', '$2a$12$9TZ3IJjTaYIB9Yrct5YT.ey7DGsybJR6d8nMb.Q8coFDKwqUMTCju', 'bbb', False);
INSERT INTO reports (device_id, report, last_user_id) values(1, 'Is fuckity fucked', 1);
INSERT INTO reports (device_id, report, last_user_id) values(3, 'No way Jose', 1);
INSERT INTO reports (device_id, report, last_user_id) values(8, 'Is starting to get full should be addressed', 2);
