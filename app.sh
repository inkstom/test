#!/bin/bash

stonk="MSFT"
key="C227WD9W3LUVKVV9"
ndays=$1
declare -a list

#list=$(curl -sS https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY_ADJUSTED&symbol=$stonk | grep "4. close" | awk -F "\"" '{print $4}')
list=$(curl -sS https://www.alphavantage.co/query\?apikey\=$key\&function\=TIME_SERIES_DAILY_ADJUSTED\&symbol\=$stonk | grep "4. close" | awk -F "\"" '{print $4}' | head \-$ndays)

echo $list
