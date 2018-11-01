#!/usr/bin/python3
# coding:utf-8

import urllib.request
from bs4 import BeautifulSoup
import time
import pandas as pd
import xlwt

url = 'http://bj.lianjia.com/ershoufang/'
page = 'pg'

headers = {'User-Agent':'Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11',
    'Accept':'text/html;q=0.9,*/*;q=0.8',
    'Accept-Charset':'ISO-8859-1,utf-8;q=0.7,*;q=0.3',
    'Accept-Encoding':'gzip',
    'Connection':'close',
    'Referer':'http://www.baidu.com/link?url=_andhfsjjjKRgEWkj7i9cFmYYGsisrnm2A-TN3XZDQXxvGsM9k9ZZSnikW2Yds4s&amp;amp;wd=&amp;amp;eqid=c3435a7d00146bd600000003582bfd1f'
}

tp = []
up = []
hi = []
for i in range(1, 5):
    i = str(i)
    a = (url + page + i + '/')
    r = urllib.request.Request(url = a)
    q = urllib.request.urlopen(r)
    soup = BeautifulSoup(q, "html.parser")
    price = soup.findAll("div", "priceInfo")
    unitprice = soup.findAll("div", "unitPrice")
    info = soup.find_all('div', attrs={'class':'houseInfo'})
    for a in price:
        totalPrice = a.span.string
        tp.append(totalPrice)

    for a in unitprice:
        totalprice = a.get_text()
        up.append(totalprice)

    for a in info:
        totalInfo = a.get_text()
        hi.append(totalInfo)

    time.sleep(1)



a = dict(zip(tp, hi))
#print(a)

house = pd.DataFrame({'totalprice': tp, 'unitprice': up, 'info': hi})
print(house)


#houseinfo_split = pd.DataFrame((x.split('/') for x in house.info),index=house.index,columns=['小区','户型','面积','朝向','装修','电梯'])
#print(houseinfo_split)

book = xlwt.Workbook(encoding='utf-8', style_compression=0)
sheet = book.add_sheet('hinfo', cell_overwrite_ok=True)
for a in house.index:
    info = house.loc[a].values
    x = info.tolist()
    for j in range(len(x)):
        print(a, j, x[j])
        sheet.write(a, j, x[j])

book.save(r'./hinfo.xls')
    
