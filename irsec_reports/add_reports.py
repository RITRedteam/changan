import requests
import logging
from IPython import embed
import urllib3
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

url_base = "https://127.0.0.1:8080/api/v1/"
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

def equal_dicts(d1, d2, ignore_keys):
    ignored = set(ignore_keys)
    for k1, v1 in d1.iteritems():
        if k1 not in ignored and (k1 not in d2 or d2[k1] != v1):
            return False
    for k2, v2 in d2.iteritems():
        if k2 not in ignored and k2 not in d1:
            return False
    return True

def add_devices():
    url = '{}reports'.format(url_base)
    # Test add
    for i in range(1, 11):
        report = {'ip': '10.2.{}.254'.format(i), 'title': 'tcl stuff', 'report': 'rootkit in ios.tcl, bot in bot.tcl, recurring script in open.tcl that drops firewalls, enables all transport methods, disables enable authentication, and ensures we have vtys'}
        r = requests.put(url, json=report, verify=False)
        report = {'ip': '10.2.{}.20'.format(i), 'title': 'rootkit and botnet', 'report': 'rootkit and anomoly bot deployed to sammich and sammichproc. sammichprocess is a service that runs the sammichproc anomoly bot, botnet is also at avahi service'}
        r = requests.put(url, json=report, verify=False)
        report = {'ip': '10.2.{}.30'.format(i), 'title': 'rootkit and botnet', 'report': 'rootkit and anomoly bot deployed to sammich and sammichproc. sammichprocess is a service that runs the sammichproc anomoly bot, botnet is also at thermald'}
        r = requests.put(url, json=report, verify=False)
        report = {'ip': '172.21.{}.10'.format(i), 'title': 'rootkit and botnet', 'report': 'rootkit and anomoly bot deployed to sammich and sammichproc. sammichprocess is a service that runs the sammichproc anomoly bot, botnet is also thermald'}
        r = requests.put(url, json=report, verify=False)
        report = {'ip': '172.21.{}.20'.format(i), 'title': 'rootkit and botnet', 'report': 'rootkit and anomoly bot deployed to sammich and sammichproc. sammichprocess is a service that runs the sammichproc anomoly bot, botnet is also thermald'}
        r = requests.put(url, json=report, verify=False)
        report = {'ip': '10.2.{}.20'.format(i), 'title': 'users', 'report': '99 interns and howard'}
        r = requests.put(url, json=report, verify=False)
        report = {'ip': '10.2.{}.30'.format(i), 'title': 'users', 'report': '99 interns and howard'}
        r = requests.put(url, json=report, verify=False)
        report = {'ip': '172.21.{}.10'.format(i), 'title': 'users', 'report': '99 interns and howard'}
        r = requests.put(url, json=report, verify=False)
        report = {'ip': '171.21.{}.20'.format(i), 'title': 'users', 'report': '99 interns and howard'} # errors?
        r = requests.put(url, json=report, verify=False)

    report = {'ip': '10.2.5.254'.format(i), 'title': 'baud rate changed and credential extraction', 'report': 'changed baud rate to 52600, and found creds with mode 7.'}
    r = requests.put(url, json=report, verify=False)
    report = {'ip': '10.2.4.254', 'title': 'extracted credentials', 'report': 'found plaintext enable password'}
    r = requests.put(url, json=report, verify=False)
    report = {'ip': '10.2.7.254', 'title': 'extracted credentials', 'report': 'found plaintext enable password'}
    r = requests.put(url, json=report, verify=False)

def main():
    #populate_test_db.main()
    add_devices()
    #test_ips()

if __name__ == '__main__':
    main()
