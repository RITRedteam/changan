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

def run_test():
    url = '{}devices'.format(url_base)
    # Test add
    device = {
            "device_name": "test_device",
            "team": "test_team",
            "owner": "test_owner",
            "location": "test_location",
            "interfaces": [{
                "mac": "ff:ff:ff:ff:ff:ff",
                "interface_name": "eth0",
                "ips": [{
                    "ip":"172.16.10.7"
                    }]
            }]
    }
    r = requests.put(url, json=device, verify=False)

def main():
    run_test()

if __name__ == '__main__':
    main()
