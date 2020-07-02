# Activities Service API Doc

|版本|说明|
|:---|:--:|
|v1.0|初版|

---

## 通用规定

- 返回结构

```json
{"code": 0, "data":{}, "msg":"successful"}
```

---

### 新增待上线活动

| **Request**

|字段|类型|说明|必须|
|:---|:--:|:---:|:---:|
|activity_id|int|活动ID|Y|

- Method: **POST**
- Path: ```/backstage/activity```
- Header:
- Body:

```json
{"activity_id": 1}
```

| **Response**

- Body:

```json
{"code": 0, "data":{}, "msg":"successful"}
```

---

### 更新待上线活动

| **Request**

|字段|类型|说明|必须|
|:---|:--:|:---:|:---:|
|activity_id|int|活动ID|Y|

- Method: **PUT**
- Path: ```/backstage/activity```
- Header:
- Body:

```json
{"activity_id": 1}
```

| **Response**

- Body:

```json
{"code": 0, "data":{}, "msg":"successful"}
```

---

### 删除待上线活动

| **Request**

|字段|类型|说明|必须|
|:---|:--:|:---:|:---:|
|activity_id|int|活动ID|Y|

- Method: **DELETE**
- Path: ```/backstage/activity```
- Header:
- Body:

```json
{"activity_id": 1}
```

| **Response**

- Body:

```json
{"code": 0, "data":{}, "msg":"successful"}
```

---

### 提前下线活动

| **Request**

|字段|类型|说明|必须|
|:---|:--:|:---:|:---:|
|activity_id|int|活动ID|Y|

- Method: **POST**
- Path: ```/backstage/activity/offline```
- Header:
- Body:

```json
{"activity_id": 1}
```

| **Response**

- Body:

```json
{"code": 0, "data":{}, "msg":"successful"}
```

---
