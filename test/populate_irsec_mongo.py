#! /usr/bin/python
import time
from pymongo import MongoClient, ASCENDING
from werkzeug.security import generate_password_hash
# TODO make usernames unique

subnet_objs = []
for i in range(1,11):
    subnet_objs.append(('Team{}'.format(i), '10.2.{}.0'.format(i) ,24))
    subnet_objs.append(('Team{} vpc'.format(i), '172.21.{}.0'.format(i), 24))
    print(i)
subnet_objs.append(('TEST', '172.16.156.0', 24))

def insert_data(db):
    subnets = db.subnets
    for sub in subnet_objs:
        subnets.insert_one(
                {
                    'subnet_name': sub[0],
                    'ip': sub[1],
                    'mask': sub[2]
                }
        )
    subnets_objs = subnets.find()

def main():
    client = MongoClient()
    client.drop_database("changan_test")
    db = client.changan_test

    insert_data(db)

if __name__ == '__main__':
    main()
