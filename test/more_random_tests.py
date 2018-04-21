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
    url = '{}device'.format(url_base)
    # Test add
    device = {
            "device_name": "SQL",
    }
    r = requests.get(url, json=device, verify=False)
    embed()

def main():
    run_test()

if __name__ == '__main__':
    main()
