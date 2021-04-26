from bs4 import BeautifulSoup
import requests

"""
url = 'https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT'
page = requests.get(url)
soup = BeautifulSoup(page.text, 'lxml')
price = soup.find('4.', close =).text
print(price)
"""

url = 'https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT'
response = requests.get(url, timeout=240)
response.status_code
response.json()

content = response.json()
key = content.keys()

print(content)
print(key)
