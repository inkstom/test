import requests
import html
import json
from itertools import chain
import pandas as pd

#from json import encoder
#encoder.FLOAT_REPR = lambda o: format(o, '.2f')

url = "https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT"
r = requests.get(url)
cont = json.loads(r.content.decode())

"""
page_source = r.text
price = page_source.find("4. close")
print(price)
page_source = page_source.split('\n')
#print("\nURL:", url)
#print("--------------------------------------")
# print the first five lines of the page source
#for row in page_source[:5]:
#    print(row)
#price = page_source.find("4. close")
"""

# returns JSON object as
# a dictionary

# get tuples of Meta Data, Misc info listed there
for meta_data, misc_info in cont.iteritems():
    #print meta_data, misc_info
    for time_series, dates in misc_info.iteritems():
        #print(dates)
            #print each_day.get("4.close", "Warning: no entry found for 'close'")
        #print (dates[0]["4. close"])
        once = dates.astype('int64')
        #closes = (dates[0]["4. close"])
        (dates[0]["4. close"]).astype('int64')
        #final = closes.astype('int64')
        print(once)
