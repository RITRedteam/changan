import requests
import populate_test_db
import populate_test_mongo
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
    url = '{}devices'.format(url_base)
    # Test add
    for i in range(1, 11):
        print(i)
        device = {"device_name": "wakanda_{}".format(i), "interfaces": [{"interface_name": "eth0", "ips": [{"ip":"10.2.{}.20".format(i)}]}]}
        r = requests.put(url, json=device, verify=False)
        device = {"device_name": "gotham_{}".format(i), "interfaces": [{"interface_name": "eth0", "ips": [{"ip":"10.2.{}.10".format(i)}]}]}
        r = requests.put(url, json=device, verify=False)
        device = {"device_name": "atlantis_{}".format(i), "interfaces": [{"interface_name": "eth0", "ips": [{"ip":"10.2.{}.30".format(i)}]}]}
        r = requests.put(url, json=device, verify=False)
        device = {"device_name": "smallville_{}".format(i), "interfaces": [{"interface_name": "eth0", "ips": [{"ip":"10.2.{}.40".format(i)}]}]}
        r = requests.put(url, json=device, verify=False)
        device = {"device_name": "krypton_{}".format(i), "interfaces": [{"interface_name": "eth0", "ips": [{"ip":"172.21.{}.10".format(i)}]}]}
        r = requests.put(url, json=device, verify=False)
        device = {"device_name": "asgard_{}".format(i), "interfaces": [{"interface_name": "eth0", "ips": [{"ip":"172.21.{}.20".format(i)}]}]}
        r = requests.put(url, json=device, verify=False)
        device = {"device_name": "themiscrya_{}".format(i), "interfaces": [{"interface_name": "eth0", "ips": [{"ip":"10.5.{}.50".format(i)}]}]}
        r = requests.put(url, json=device, verify=False)
        device = {"device_name": "router_{}".format(i), "interfaces": [{"interface_name": "eth0", "ips": [{"ip":"10.2.{}.254".format(i)}]}]}
        r = requests.put(url, json=device, verify=False)


def main():
    #populate_test_db.main()
    add_devices()
    #test_ips()

if __name__ == '__main__':
    main()
