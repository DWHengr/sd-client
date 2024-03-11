# 服务发现-云端

### 管理页面

127.0.0.1:8282

### 相关api

**接口名称：**获取服务列表

**接口url：**/api/list

**请求方式：**GET

**数据格式：**Application/Json

**返回 response Body：**

```json
{
  "code": 0,
  "data": [
    {
      "id": "ed6c2289-13b1-4a66-b090-314f03bfd500",
      "name": "node3",
      "mac": "52:54:11:a8:6c:18",
      "domain": "www.node3.com",
      "ip": "172.16.10.3",
      "isPing": true,
      "isManuallyModify": false
    },
    {
      "id": "05c088cd-3ee8-4c6e-8037-2409d41a6ae6",
      "name": "node4",
      "mac": "52:54:11:6a:be:11",
      "domain": "WWW.node4.COM",
      "ip": "172.16.10.4",
      "isPing": true,
      "isManuallyModify": true
    }
  ]
}
```

### 相关文件描述

**config.yml文件: 程序相关配置文件** 

**templates目录: html页面相关内容**

**bind_zone_tpl.txt:  bind生成zones文件对应的模板**

**data.json:  服务列表持久化的文件**