import sys
import argparse
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

# Function to add "https://" if only domain is provided
def normalize_url(url):
    if not url.startswith("http://") and not url.startswith("https://"):
        url = "https://" + url
    return url.strip()

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
def send_request(url, method):
    try:
        response = requests.request(method, url)
        return str(response.status_code), method
    except requests.exceptions.RequestException as e:
        return f"Error occurred with {method} request: {e}", method

# Function to handle single URL input
def handle_single_url(url, methods):
    url = normalize_url(url)
    with ThreadPoolExecutor(max_workers=10) as executor:
        futures = [executor.submit(send_request, url, method) for method in methods]
        print(f"(+) URL: {Fore.BLUE}{url}{Style.RESET_ALL}")
        for future in futures:
            status_code, method = future.result()
            print_colored_output(status_code, method)

# Function to handle file input
def handle_file_input(urls, methods):
    with ThreadPoolExecutor(max_workers=10) as executor:
        for url in urls:
            url = normalize_url(url)
            futures = [executor.submit(send_request, url, method) for method in methods]
            print(f"(+) URL: {Fore.BLUE}{url}{Style.RESET_ALL}")
            for future in futures:
                status_code, method = future.result()
                print_colored_output(status_code, method)

# Parse command-line arguments
parser = argparse.ArgumentParser(description="HTTP Method Tester Tool")
parser.add_argument("--url", help="Test a single URL", metavar="URL")
args = parser.parse_args()

if args.url:
    handle_single_url(args.url, http_methods)
else:
    # Read URLs from stdin
    urls = [line.strip() for line in sys.stdin]
    if not urls:
        print("No URLs provided.")
        sys.exit(1)
    handle_file_input(urls, http_methods)
