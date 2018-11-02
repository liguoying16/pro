#!/usr/bin/python3
# coding:utf-8

import wechatsogou

ws_api = wechatsogou.WechatSogouAPI()

print(ws_api)

#print(ws_api.get_gzh_info('南航青年志愿者'))
#print(ws_api.search_gzh('南京航空航天大学'))
print(ws_api.get_gzh_info('BigHeart16'))
