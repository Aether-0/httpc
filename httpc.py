import sys
import requests
from concurrent.futures import ThreadPoolExecutor
from colorama import Fore, Style

# Banner
BANNER = """
\033[94m+----------------------------------------+
|        __  __________________  ______   |
|       / / / /_  __/_  __/ __ \/ ____/   |
|      / /_/ / / /   / / / /_/ / /        |
|     / __  / / /   / / / ____/ /___      |
|    /_/ /_/ /_/   /_/ /_/    \____/      |
+----------------------------------------+
\033[0m
"""

AUTHOR_INFO = """
\033[95mAuthor   : Aether
Telegram : @a37h3r
Github   : https://github.com/Aether-0
\033[0m
"""

print(BANNER)
print(AUTHOR_INFO)

# Check if URL argument is provided
if len(sys.argv) < 2:
    print("Usage: {} <URL>".format(sys.argv[0]))
    sys.exit(1)

# Function to add "https://" if only domain is provided
def normalize_url(url):
    if not url.startswith("http://") and not url.startswith("https://"):
        url = "https://" + url
    return url

url = normalize_url(sys.argv[1])

# List of HTTP methods to test
http_methods = [
    "OPTIONS", "GET", "HEAD", "POST", "PUT", "DELETE", "TRACE", "TRACK",
    "DEBUG", "PURGE", "CONNECT", "PROPFIND", "PROPPATCH", "MKCOL", "COPY",
    "MOVE", "LOCK", "UNLOCK", "VERSION-CONTROL", "REPORT", "CHECKOUT", "CHECKIN",
    "UNCHECKOUT", "MKWORKSPACE", "UPDATE", "LABEL", "MERGE", "BASELINE-CONTROL",
    "MKACTIVITY", "ORDERPATCH", "ACL", "PATCH", "SEARCH", "ARBITRARY", "BIND",
    "LINK", "MKCALENDAR", "MKREDIRECTREF", "PRI", "QUERY", "REBIND", "UNBIND",
    "UNLINK", "UPDATEREDIRECTREF"
]

# Function to print colored output based on status code
def print_colored_output(status_code, method):
    if status_code.startswith('2'):
        color = Fore.GREEN  # Green
    elif status_code.startswith('3'):
        color = Fore.BLUE  # Blue
    elif status_code.startswith('4'):
        color = Fore.RED  # Red
    elif status_code.startswith('5'):
        color = Fore.YELLOW  # Yellow
    else:
        color = Fore.WHITE  # Default color
    
    print(f"{color}[{method}]  {status_code}{Style.RESET_ALL}")

# Function to send request and get response code
def send_request(method):
    try:
        response = requests.request(method, url)
        return str(response.status_code), method
    except requests.exceptions.RequestException as e:
        return f"Error occurred with {method} request: {e}", method

# Use ThreadPoolExecutor for concurrent execution
with ThreadPoolExecutor(max_workers=10) as executor:
    futures = [executor.submit(send_request, method) for method in http_methods]
    for future in futures:
        status_code, method = future.result()
        print_colored_output(status_code, method)
