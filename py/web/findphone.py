#!/usr/bin/python3

# -*- coding: utf-8 -*-
import urllib.request
import requests
url = "https://item.jd.com/3888278.html"
try:
    r = requests.get(url)
    r.raise_for_status()
    r.encoding = r.apparent_encoding
    print(r.text[:1000])
except:
    print("爬取失败")
