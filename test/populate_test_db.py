import mysql.connector as mariadb
import logging

# TODO need to add datetimes eventually
logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

tables = ['interfaces', 'ips', 'reports', 'users', 'subnets', 'devices', 'sessions']
test_db_name = 'changan_test'
device_fields = ['device_id', 'device_name', 'team', 'owner', 'location']
devices = [
# DeviceID Name Team Owner Macs IPs Location
    (1, 'HTTP', 'OPS', 'ksam', 'china'),
    (2, 'DNS', 'OPS', 'ksam', 'china'),
    (3, 'SQL', 'OPS', 'jnottingham', 'russia'),
    (4, 'Router', 'Networking', 'jgem', 'LA'),
    (5, 'AD', 'SysAd', 'bwhite', 'laptop'),
    (6, 'Jenkins oh god', 'Engineering', 'rgeorge', 'LA'),
    (7, 'gitlab', 'Engineering', 'rgeorge', 'LA'),
    (8, 'esxi', 'OPS', 'ksam', 'china')
]
subnet_fields = ['subnet_id', 'subnet_name', 'ip', 'mask']
subnets = [
# SubnetID Name IP Mask
    (1, 'china office', '172.16.9.0', '24'),
    (2, 'russia office', '172.16.10.0', '24'),
    (3, 'LA office', '172.16.11.0', '24'),
    (4, 'Prod', '10.100.0.0', '16')
]
ip_fields = ['ip_id', 'ip', 'subnet_id']
ips = [
# IPID Name IP
    (1, '10.100.100.200', 4),
    (2, '10.100.100.201', 4),
    (3, '10.100.100.202', 4),
    (4, '10.100.100.100', 4),
    (5, '172.16.11.203', 3),
    (6, '172.16.11.124', 3),
    (7, '172.16.10.1', 2),
    (8, '172.16.9.1', 1),
    (9, '172.16.9.53', 1),
    (10, '172.16.10.101', 2)
]
interface_fields = ['device_id', 'ip_id', 'mac', 'interface_name']
interfaces = [
    (1, 8, 'AA:BB:CC:DD:EE:FF', 'eth0'),
    (2, 9, 'AA:BB:CC:DD:EE:EE', 'eth0'),
    (3, 7, 'AA:BB:CC:DD:EE:DD', 'eth0'),
    (4, 6, 'AA:BB:CC:DD:EE:CC', 'eth0'),
    (5, 10, 'AA:BB:CC:DD:EE:BB', 'eth0'),
    (6, 1, 'AA:BB:CC:DD:EE:AA', 'eth0'),
    (7, 2, 'AA:BB:CC:DD:FF:FF', 'eth0'),
    (8, 3, 'AA:BB:CC:DD:FF:EE', 'eth0')
]
user_fields = ['user_id', 'username', 'password', 'api_key', 'active']
users = [
    # password is testtest
    (1, 'rwhittier', '$2a$12$9TZ3IJjTaYIB9Yrct5YT.ey7DGsybJR6d8nMb.Q8coFDKwqUMTCju', 'aaa', True),
    (2, 'test', '$2a$12$9TZ3IJjTaYIB9Yrct5YT.ey7DGsybJR6d8nMb.Q8coFDKwqUMTCju', 'bbb', False)
]
report_fields = ['device_id', 'report', 'last_user_id']
reports = [
    (1, 'Is fuckity fucked', 1),
    (3, 'No way Jose', 1),
    (8, 'Is starting to get full should be addressed', 2)
]
# unsure on interfaces and how to structure it gonna have a many to many with ips, should move to a
# join table at some point
create_tables_strings = [
'''CREATE TABLE IF NOT EXISTS devices (
    device_id int(5) NOT NULL AUTO_INCREMENT,
    device_name varchar(50) DEFAULT NULL,
    team varchar(20) DEFAULT NULL,
    owner varchar(30) DEFAULT NULL,
    location varchar(250) DEFAULT NULL,
    PRIMARY KEY(device_id)
    );''',
'''CREATE TABLE IF NOT EXISTS subnets (
    subnet_id int(5) NOT NULL AUTO_INCREMENT,
    subnet_name varchar(50) DEFAULT NULL,
    ip varchar(55) DEFAULT NULL,
    mask int(5) DEFAULT NULL,
    PRIMARY KEY(subnet_id)
    );''',
'''CREATE TABLE IF NOT EXISTS ips (
    ip_id int(5) NOT NULL AUTO_INCREMENT,
    subnet_id int(5) NOT NULL,
    ip varchar(55) DEFAULT NULL,
    PRIMARY KEY(ip_id),
    FOREIGN KEY(subnet_id) REFERENCES subnets(subnet_id)
    );''',
'''CREATE TABLE IF NOT EXISTS interfaces (
    device_id int(5) NOT NULL,
    ip_id int(5) NOT NULL,
    mac varchar(17) NOT NULL,
    interface_name varchar(10) DEFAULT NULL,
    PRIMARY KEY(device_id, ip_id),
    FOREIGN KEY(device_id) REFERENCES devices(device_id),
    FOREIGN KEY(ip_id) REFERENCES ips(ip_id)
    );''',
'''CREATE TABLE IF NOT EXISTS users (
    username varchar(50) NOT NULL,
    user_id int(5) NOT NULL AUTO_INCREMENT,
    password char(60) NOT NULL,
    api_key varchar(50) DEFAULT NULL,
    active BOOL,
    PRIMARY KEY(user_id),
    UNIQUE KEY(username)
    );''',
'''CREATE TABLE IF NOT EXISTS reports (
    device_id int(5) NOT NULL,
    report varchar(3000) DEFAULT NULL,
    last_user_id int(5) DEFAULT NULL,
    FOREIGN KEY(device_id) REFERENCES devices(device_id),
    FOREIGN KEY(last_user_id) REFERENCES users(user_id)
    );''',
'''CREATE TABLE sessions (
  token CHAR(43) PRIMARY KEY,
  data BLOB NOT NULL,
  expiry TIMESTAMP(6) NOT NULL
);''',
'''CREATE INDEX sessions_expiry_idx ON sessions (expiry);'''
]

def create_database(cursor, file_obj):
    try:
        create_line = "CREATE DATABASE IF NOT EXISTS {} DEFAULT CHARACTER SET 'utf8'".format(
                test_db_name)
        cursor.execute(create_line)
        logging.info("Created TEST DB: {}".format(test_db_name))
        file_obj.write("{}\n".format(create_line))
    except mariadb.Error as err:
        logging.error("Failed creating database: {}".format(err))
        exit(1)

def create_tables(cursor, file_obj):
    for table in tables:
        drop_line = 'DROP TABLE IF EXISTS {}'.format(table)
        cursor.execute(drop_line)
        logging.debug("table dropped: {}".format(table))
        file_obj.write("{}\n".format(drop_line))
    logging.info("Dropped tables")
    for create_table in create_tables_strings:
        cursor.execute(create_table)
        logging.debug("create line ran: {}".format(create_table))
        file_obj.write("{}\n".format(create_table))

def insert_data(cursor, file_obj):
    for device in devices:
        insert_line = 'INSERT INTO devices ({}) values({}, \'{}\');'.format(', '.join(device_fields),
                device[0], '\', \''.join(device[1:]))
        logging.debug('insert line to run: {}'.format(insert_line))
        file_obj.write("{}\n".format(insert_line))
        cursor.execute(insert_line)
    for subnet in subnets:
        insert_line = 'INSERT INTO subnets ({}) values({}, \'{}\', {});'.format(
                ', '.join(subnet_fields), subnet[0], '\', \''.join(subnet[1:-1]), subnet[-1])
        logging.debug('insert line to run: {}'.format(insert_line))
        file_obj.write("{}\n".format(insert_line))
        cursor.execute(insert_line)
    for ip in ips:
        insert_line = 'INSERT INTO ips ({}) values({}, \'{}\', {});'.format(
                ', '.join(ip_fields), ip[0], ip[1], ip[2])
        logging.debug('insert line to run: {}'.format(insert_line))
        file_obj.write("{}\n".format(insert_line))
        cursor.execute(insert_line)
    for interface in interfaces:
        insert_line = 'INSERT INTO interfaces ({}) values({}, {}, \'{}\', \'{}\');'.format(
                ', '.join(interface_fields), interface[0], interface[1], interface[2],
                interface[3])
        logging.debug('insert line to run: {}'.format(insert_line))
        file_obj.write("{}\n".format(insert_line))
        cursor.execute(insert_line)
    for user in users:
        insert_line = 'INSERT INTO users ({}) values({}, \'{}\', \'{}\', \'{}\', {});'.format(
                ', '.join(user_fields), user[0], user[1], user[2], user[3], user[4])
        logging.debug('insert line to run: {}'.format(insert_line))
        file_obj.write("{}\n".format(insert_line))
        cursor.execute(insert_line)
    for report in reports:
        insert_line = 'INSERT INTO reports ({}) values({}, \'{}\', {});'.format(
                ', '.join(report_fields), report[0], report[1], report[2])
        logging.debug('insert line to run: {}'.format(insert_line))
        file_obj.write("{}\n".format(insert_line))
        cursor.execute(insert_line)


def main():
    #mariadb_connection = mariadb.conenct(user='test_user', password='test_pass', database='test')
    mariadb_connection = mariadb.connect(user='root')
    cursor = mariadb_connection.cursor()

    with open('rawsql.sql', 'w') as file_obj:
        create_database(cursor, file_obj)
        mariadb_connection.database = test_db_name
        create_tables(cursor, file_obj)
        insert_data(cursor, file_obj)
    mariadb_connection.commit()

if __name__ == '__main__':
    main()
