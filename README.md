# WBABEProject-14

# ๐ ๋ชฉ์ฐจ 
 - [์๊ฐ](#์๊ฐ) 
 - [์ฌ์ฉ๋ฒ](#์ฌ์ฉ๋ฒ)
 - [๋๋ ํ ๋ฆฌ ๊ตฌ์กฐ](#๋๋ ํ ๋ฆฌ-๊ตฌ์กฐ)
 - [๊ฐ๋ฐ ํ๊ฒฝ](#๊ฐ๋ฐ-ํ๊ฒฝ)
 - [์ฌ์ฉ ๊ธฐ์ ](#์ฌ์ฉ-๊ธฐ์ )
 - [ERD](#erd)
 - [์๋ฒ ์ํคํ์ฒ](#์๋ฒ-์ํคํ์ฒ) 
 - [Api ๋ช์ธ์](#api-๋ช์ธ์)
 - [ํผ๋๋ฐฑ ๊ฐ์ ](#ํผ๋๋ฐฑ-๊ฐ์ )
 - [์ฃผ์ ๊ธฐ๋ฅ](#์ฃผ์-๊ธฐ๋ฅ)
    - [ํผ์ฃผ๋ฌธ์](#ํผ์ฃผ๋ฌธ์)
        - [๋ฉ๋ด ์์ฑ](#๋ฉ๋ด-์์ฑ)
        - [๋ฉ๋ด ์์ ](#๋ฉ๋ด-์์ )
        - [๋ฉ๋ด ์ญ์ ](#๋ฉ๋ด-์ญ์ )
        - [์ถ์ฒ ๋ฉ๋ด ๋ณ๊ฒฝ](#์ถ์ฒ-๋ฉ๋ด-๋ณ๊ฒฝ)
        - [ํ์ฌ ์ฃผ๋ฌธ๋ด์ญ ์กฐํ](#ํ์ฌ-์ฃผ๋ฌธ๋ด์ญ-์กฐํ)
        - [๋ฉ๋ด๋ณ ์ฃผ๋ฌธ์์ฒญ ์ํ ๋ณ๊ฒฝ](#๋ฉ๋ด๋ณ-์ฃผ๋ฌธ์์ฒญ-์ํ-๋ณ๊ฒฝ)
    - [์ฃผ๋ฌธ์](#์ฃผ๋ฌธ์)
        - [๋ฉ๋ด ๋ฆฌ์คํธ ์กฐํ](#๋ฉ๋ด-๋ฆฌ์คํธ-์กฐํ)
        - [์ฃผ๋ฌธ ์์ฑ](#์ฃผ๋ฌธ-์์ฑ)
        - [์ฃผ๋ฌธ ๋ฉ๋ด ๋ณ๊ฒฝ](#์ฃผ๋ฌธ-๋ฉ๋ด-๋ณ๊ฒฝ)
        - [๋ฆฌ๋ทฐ ์์ฑ](#๋ฆฌ๋ทฐ-์์ฑ)
        - [๋ฉ๋ด ํ์  ๋ฆฌ๋ทฐ ์กฐํ](#๋ฉ๋ด-ํ์ -๋ฆฌ๋ทฐ-์กฐํ)
        - [ํ์ฌ ์ฃผ๋ฌธ ๋ด์ญ ์กฐํ](#ํ์ฌ-์ฃผ๋ฌธ-๋ด์ญ-์กฐํ)
 
    
    
# ์๊ฐ 

์ธํํธ ์๋์ ๊ธ์ฆํ๊ณ  ์๋ ์จ๋ผ์ธ ์ฃผ๋ฌธ ์์คํ์ ์ด๋ฏธ ์ํ์ ๋ฐ์ ๊ทธ ์ํฅ์ ๋ผ์น๊ณ  ์๋ ์ํฉ์, ๊ฐ๊น๊ฒ๋ ๋ฐฐ๋ฌ์ดํ, ๋งค์ฅ์๋ ํค์ค์คํฌ, <br> ์๋น์๋ ํจ๋๋ฅผ ์ด์ฉํ ๋ฉ๋ด ์ฃผ๋ฌธ๊น์ง ๊ทธ ์ฌ์ฉ๋ฒ์๊ฐ ์ ์  ํ๋๋์ด ๊ฐ๊ณ  ์์ต๋๋ค.<br> ์ด๋ฐ ์๋์ ํด๋น ์์คํ์ ์ดํด, ๊ฒฝํํ๊ณ  ๊ฐ ๋จ๊ณ๋ณ ํ๋ก์ธ์ค๋ฅผ ์ดํดํ์ฌ ๊ตฌํํจ์ผ๋ก์จ ์๋น์ค ๊ตฌ์ถ์ ๊ฒฝํ์ ์๊ณ , golang์ ์ดํด๋ฅผ ๋์ต๋๋ค.

1. ์ฃผ๋ฌธ์/ํผ์ฃผ๋ฌธ์์ ์ญํ ์์ ํ์์ ์ธ ๊ธฐ๋ฅ์ ๋์ถ, ๊ตฌํํฉ๋๋ค.
2. ํด๋น ์์คํ์ ๋ํด ์๊ตฌ์ฌํญ์ ์ ์ํ๊ณ  ์ฃผ๋ฌธ์์ ํผ์ฃผ๋ฌธ์ ์์ฅ์์ ํ์ํ ๊ธฐ๋ฅ์ ๋์ถํ์ฌ, ๊ธฐ๋ฅ์ ํ์ฅํ๊ณ  ์ฃผ๋ฌธ ์๋น์ค๋ฅผ ์ํ ํ๊ฒ ์ง์ํ ์ ์๋ ๊ธฐ๋ฅ์ ๊ตฌํํฉ๋๋ค.
3. ์ฃผ๋ฌธ์๋ ์ ๋ขฐ์๋ ์ฃผ๋ฌธ๊ณผ ๋ฐฐ๋ฌ๊น์ง๋ฅผ ์ํฉ๋๋ค. ๋, ํผ์ฃผ๋ฌธ์๋ ์ฃผ๋ฌธ๋ด์ญ์ ๊ด๋ฆฌํ๊ณ  ์ํ ํ ์๋น์ค๊ฐ ์ ๊ณต๋์ด์ผ ํฉ๋๋ค.

### ์ฌ์ฉ๋ฒ
```bash
$ docker start mongodb

$ git clone https://github.com/codestates/WBABEProject-14.git
$ cd go-mvc-project/
$ go mod tidy
> go run main.go
```

### ๊ฐ๋ฐ ํ๊ฒฝ
- Mac OS M1
- vscode
- MongoDB Compass
- Docker 
- PostMan
- Talen API Tester

### ์ฌ์ฉ ๊ธฐ์ 
- Go(go1.19.4 darwin/arm64)
- Gin Gonic
- MongoDB
- Swagger 

### ๋๋ ํ ๋ฆฌ ๊ตฌ์กฐ 

```
โโโ config         
โ   โโโ config.go  
โ   โโโ config.toml
โโโ controller     
โ   โโโ controller.go
โ   โโโ orderer.controller.go /* ์ฃผ๋ฌธ์ ์ปจํธ๋กค๋ฌ */
โ   โโโ taker.controller.go   /* ํผ์ฃผ๋ฌธ์ ์ปจํธ๋กค๋ฌ */
โโโ docs
โ   โโโ docs.go
โ   โโโ swagger.json
โ   โโโ swagger.yaml
โโโ go.mod
โโโ go.sum
โโโ logger
โ   โโโ logger.go
โโโ logs
โ   โโโ go-loger_2022-12-20.log
โโโ main.go
โโโ model
โ   โโโ menu.go     /* ๋ฉ๋ด ๋ชจ๋ธ */
โ   โโโ model.go    /* ๋ชจ๋ธ ์ ์ */
โ   โโโ order.go    /* ์ฃผ๋ฌธ ๋ชจ๋ธ */
โ   โโโ response.go /* ์๋ต ๋ฉ์ธ์ง ๋ชจ๋ธ */
โ   โโโ review.go   /* ๋ฆฌ๋ทฐ ๋ชจ๋ธ */
โ   โโโ swaggerModel.go /* ์ค์จ๊ฑฐ ์์ฑ์ ์์ฒญ๊ณผ ์๋ต์ ๋ํ ๋ชจ๋ธ */
โโโ route
โ   โโโ route.go
โโโ services
    โโโ orderer.service.go      /* ์ฃผ๋ฌธ์ ์๋น์ค ์ธํฐํ์ด์ค */
    โโโ orderer.service.impl.go /* ์ฃผ๋ฌธ์ ์๋น์ค ๊ตฌํ์ฒด */
    โโโ taker.service.go        /* ํผ์ฃผ๋ฌธ์ ์๋น์ค ์ธํฐํ์ด์ค */
    โโโ taker.service.impl.go   /* ํผ์ฃผ๋ฌธ์ ์๋น์ค ๊ตฌํ์ฒด */
```

### ERD
<img src="https://user-images.githubusercontent.com/103496262/211019892-79a95a76-e678-4d7b-ba73-82a58e2e87e9.png"/>

### ์๋ฒ ์ํคํ์ฒ 

<img src="https://user-images.githubusercontent.com/103496262/209453624-da34cdbd-62cb-41c3-80b8-c73a3975ba3e.png"/>


### API ๋ช์ธ์

<details>
<summary>๐ Swagger ์คํฌ๋ฆฐ์ท ๋ณด๊ธฐ</summary>
<div markdown="1">

<img src="https://user-images.githubusercontent.com/103496262/210320765-7a5e36ca-efc0-44fa-9544-08f21d8d6132.png"/>

</div>
</details>


```
ํ๋ก์ ํธ ์คํํ ์๋ ๋๋ฉ์ธ ์ผ๋ก ์ ์ 
http://localhost:8080/swagger/index.html#/ 
```

# ํผ๋๋ฐฑ ๊ฐ์ 

+ [x] [created_at, updated_at ํ ์ธํธ๋ก ๊ตฌ์ฑํ๊ธฐ](https://github.com/codestates/WBABEProject-14/commit/e756f9bf36a119a4ca6f77dc9faa1c7a7ef9a0ad)
+ [x] [์ฃผ๋ฌธ ์ํ๋ฅผ Enum, Const ํ์ฉํ๊ธฐ](https://github.com/codestates/WBABEProject-14/commit/bd302a8d54df3d60602e0704e04289952b870c67)
+ [x] [์ฌ์ฉํ์ง ์๋ ์ฃผ์ ์ฝ๋ ์ ๊ฑฐ](https://github.com/codestates/WBABEProject-14/commit/0dfa0822f2ba5ba0972c4f70096b727179eb83d1)
+ [x] [rest Api ์ ๋ง๊ฒ ๋ฆฌํฉํ ๋ง URI ์ ์์๋ง ๋ช์ ( detail ์ญ์  )](https://github.com/codestates/WBABEProject-14/commit/000e7fd298956740da9d7bdede8b3dc1a861ad85)
+ [x] [๋ฉ๋ด ์ถ๊ฐ, ๋ณ๊ฒฝ๊ณผ ๊ฐ์ ๊ฒ์ ๊ตฌ๋ถํ๊ธฐ ์ํด์  ์ผ๋ฐ์ ์ผ๋ก ์ฟผ๋ฆฌ์คํธ๋ง์ผ๋ก ์ฌ์ฉ](https://github.com/codestates/WBABEProject-14/commit/5672a10c5b346c0fa9784c1c57306be2bccefa01)
+ [x] [์ง๊ด์ ์ด์ง ๋ชปํ ๋ค์ด๋ฐ ์์ ](https://github.com/codestates/WBABEProject-14/commit/16b877c611e43f92cb02176e29c1d21e67c5fbfc)
+ [x] [ํจ์์ ์ฝ๋๊ฐ ๊ธธ์ด์ง๋ ๊ฒฝ์ฐ ํ๋์ ์ฑ์์ ๊ฐ์ง์ ์๋๋ก ๋ถ๋ฆฌ ํ์](https://github.com/codestates/WBABEProject-14/commit/4a8f002705911e9c2358c3248790f6cebdac6823)
+ [x] [๋ฉ๋ด ์ถ๊ฐ์ ๋ฉ๋ด ์ด๋ฆ๋ง์ผ๋ก ์ค๋ณต๊ฒ์ฌ ์ฒ๋ฆฌ๋ ์ฉ์ดํ์ง ๋ชปํจ](https://github.com/codestates/WBABEProject-14/commit/0c6b0285a378b5ca73bf660eed059c1e1a7bf02f)
+ [x] [CORS ํ์ฉ์ ํน์  ๋๋ฉ์ธ๋ง ์ ์ฉํ๊ธฐ](https://github.com/codestates/WBABEProject-14/blob/2c7ff2849b162ae369e5173197e121e5263c41ef/go-mvc-project/route/route.go#L34)

# ์ฃผ์ ๊ธฐ๋ฅ

# ํผ์ฃผ๋ฌธ์
## ๋ฉ๋ด ์์ฑ
[func CreateMenu()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/taker.service.impl.go#L28-L57)
### Request
`POST /api/v01/taker/menu`
```json
{
  "storenumber":100,
  "limitorder": 300,
  "menuname": "์ฐ๋",
  "orderstatus": false,
  "origin": "๊ตญ๋ด์ฐ",
  "price": 15000,
  "spice": 0
}
```

### Response
```json
{
  "message": "success"
}
```
## ๋ฉ๋ด ์์ 
[func UpdateMenu()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/taker.service.impl.go#L59-L87)
### Request
`PUT /api/v01/taker/menu/์ฐ๋`

```json
{
  "limitorder": 300,
  "orderstatus": false,
  "origin": "๊ตญ๋ด์ฐ",
  "price": 5500,
  "spice": 0
}
```
### Response 
```json
{
  "message": "success"
}
```
## ๋ฉ๋ด ์ญ์  
[func DeleteMenu()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/taker.service.impl.go#L77-L87)
### Request
`DELETE /api/v01/taker/menu`
```json
{
  "menuname": "์ฐ๋"
}
```
### Response 
```json
{
  "message": "success"
}
```
## ์ถ์ฒ ๋ฉ๋ด ๋ณ๊ฒฝ
[func UpdateMenuRecommend()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/taker.service.impl.go#L89-L112)

```bash
PATCH /api/v01/taker/menu
```
### Request
```json
{
  "menuname": "์ฐ๋"
}
```
### Response 
```json
{
  "menuList": [
    {
      "_id": "63a73a1c8d989838729bc114",
      "createdat": "2022-12-24T13:17:12.793+00:00",
      "grade": 3.5,
      "isdelete": false,
      "limitorder": 100,
      "menuname": "๋ก๋ณถ์ด",
      "orderstatus": true,
      "origin": "๊ตญ๋ด์ฐ",
      "price": 10000,
      "recommend": true,
      "reorder": 30,
      "spice": 3
    },
    {
      "_id": "63a79abf892e3345fd2f2150",
      "createdat": "2022-12-24T16:22:41.551+00:00",
      "grade": 2.5,
      "isdelete": false,
      "limitorder": 200,
      "menuname": "๊น๋ฐฅ",
      "orderstatus": true,
      "origin": "๊ตญ๋ด์ฐ",
      "price": 5000,
      "recommend": true,
      "reorder": 50,
      "spice": 0
    }
    ...
  ]
}
```
## ํ์ฌ ์ฃผ๋ฌธ๋ด์ญ ์กฐํ 
[func GetOrderList()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/taker.service.impl.go#L114-L128)


`GET /api/v01/taker/orders`

### Request
### Response 
```json
{
  "orders": [
    {
      "_id": "63a73a1c8d989838729bc114",
      "address": "์์ธ์ ๊ฐ๋จ๊ตฌ ๊ฐ๋จ๋ 1212-202 101ํธ",
      "createdat": "2022-12-24T16:17:12.793+00:00",
      "isdelete": false,
      "isexistreview": false,
      "menuname": "๋ก๋ณถ์ด",
      "ordernumber": 11,
      "phone": "010xxxxxxxx",
      "status": 0
    },
    {
      "_id": "63a79abf892e3345fd2f2150",
      "address": "์์ธ์ ์ฉ์ฐ๊ตฌ ์ฉ์ฐ๋ 2525-111 303ํธ",
      "createdat": "2022-12-24T18:11:22.440+00:00",
      "isdelete": false,
      "isexistreview": false,
      "menuname": "์ฐ๋",
      "ordernumber": 12,
      "phone": "010xxxxxxxx",
      "status": 0
    }
    ...
  ]
}
```
## ๋ฉ๋ด๋ณ ์ฃผ๋ฌธ์์ฒญ ์ํ ๋ณ๊ฒฝ
[UpdateOrderStatus](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/taker.service.impl.go#L130-L147)

`PATCH /api/v01/taker/orders/{menuname}`
### Request
```json
{
  "status": 0
}
```
### Response 
```json
{
  "message": "success"
}
```

# ์ฃผ๋ฌธ์
 
## ๋ฉ๋ด ๋ฆฌ์คํธ ์กฐํ 
[func GetAllMenu()](https://github.com/codestates/WBABEProject-14/blob/444bbb6ec3a2690a55dc284f948745caca543974/go-mvc-project/services/orderer.service.impl.go#L57-L72)

### Request 
`GET /api/v01/orderer/menu/{sort}`

`sort : [recommend, grade, reorder, createdat] ์ค ํ๋ ํํ์ฌ ์์ฒญ`

### Response 
```json
{
  "menuList": [
    {
      "_id": "63a73a1c8d989838729bc114",
      "createdat": "2022-12-24T16:17:12.793+00:00",
      "grade": 3.5,
      "isdelete": false,
      "limitorder": 100,
      "menuname": "๋ก๋ณถ์ด",
      "orderstatus": true,
      "origin": "๊ตญ๋ด์ฐ",
      "price": 150000,
      "recommend": true,
      "reorder": 30,
      "spice": 3
    },
    {
      "_id": "63a79abf892e3345fd2f2150",
      "createdat": "2022-12-24T18:42:20.191+00:00",
      "grade": 4,
      "isdelete": false,
      "limitorder": 50,
      "menuname": "์ฐ๋",
      "orderstatus": true,
      "origin": "๊ตญ๋ด์ฐ",
      "price": 5000,
      "recommend": true,
      "reorder": 50,
      "spice": 0
    },
    ...
  ]
}
```

## ์ฃผ๋ฌธ ์์ฑ
[func CreateMenu()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/orderer.service.impl.go#L35-L62)

### Request 
`POST /api/v01/orderer/order`

```json
{
  "address": "์์ธ์ ๊ฐ๋จ๊ตฌ ๊ฐ๋จ๋ 1212-202 101ํธ",
  "menuname": "๋ก๋ณถ์ด",
  "phone": "01020221225"
}
```
### Response 
```json
{
  "์ฃผ๋ฌธ๋ฒํธ": 15
}
```

## ์ฃผ๋ฌธ ๋ฉ๋ด ๋ณ๊ฒฝ
[func UpdateOrder()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/orderer.service.impl.go#L158-L182)

### Request 
`PATCH /api/v01/orderer/order/{orderId}?flag={0,1}`
```json
{
  "menuname": "์ฐ๋"
}
```

### Response 

```json
{
  "message": "success"
}
```

### OR

```json
{
  "message": "ํด๋น ์ฃผ๋ฌธ์ ๋ฐฐ๋ฌ์ค์๋๋ค. ์ ๊ท์ฃผ๋ฌธ์ผ๋ก ์ฒ๋ฆฌ๋์์ต๋๋ค.",
  "ordernumber": 17
}
```


## ๋ฆฌ๋ทฐ ์์ฑ
[func CreateReview()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/orderer.service.impl.go#L97-L156)
### Request 
`POST /api/v01/orderer/review/{orderID}`

```json
{
  "comment": "๋ง์์ด์",
  "grade": 3.5
}
```

### Response 
```json
{
  "message": "success"
}
```

## ๋ฉ๋ด ํ์  ๋ฆฌ๋ทฐ ์กฐํ
[func GetAllReview()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/orderer.service.impl.go#L81-L95)
### Request 

`GET /api/v01/orderer/detailMenu/{menuname}`

### Response 
```json
{
  "avgGrade": 3.5,
  "reviewList": [
    {
      "_id": "63a73a1c8d989838729bc114",
      "comment": "๋ง์์์ต๋๋ค.",
      "createdat": "2022-12-24T17:42:53.949+00:00",
      "grade": 3,
      "isdelete": false,
      "menuname": "๋ก๋ณถ์ด"
    },
    {
      "_id": "63a79abf892e3345fd2f2150')",
      "comment": "๋ง์์ด์.",
      "createdat": "2022-12-24T17:44:19.147+00:00",
      "grade": 4,
      "isdelete": false,
      "menuname": "๋ก๋ณถ์ด"
    },
    ...
  ]
}
```

## ํ์ฌ ์ฃผ๋ฌธ ๋ด์ญ ์กฐํ
[func Orders()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/orderer.service.impl.go#L184-L208)
### Request 
`GET /api/v01/orderer/orders`

### Response  

```json
{
  "currentOrders": [
    {
      "_id": "63a79abf892e3345fd2f2150",
      "address": "์์ธ์ ๊ฐ๋จ๊ตฌ ๊ฐ๋จ๋ 1212-202 101ํธ",
      "createdat": "2022-12-24T16:17:12.793+00:00",
      "isdelete": false,
      "isexistreview": false,
      "menuname": "๋ก๋ณถ์ด",
      "ordernumber": 1,
      "phone": "010xxxxxxxx",
      "status": 0
    },
    ...
  ],
  "pastOrders": [
    {
      "_id": "63a73a1c8d989838729bc114",
      "address": "์์ธ์ ์ฉ์ฐ๊ตฌ ์ฉ์ฐ๋ 1212-202 303ํธ",
      "createdat": "2022-12-24T16:15:44.327+00:00",
      "isdelete": false,
      "isexistreview": false,
      "menuname": "์ฐ๋",
      "ordernumber": 0,
      "phone": "010xxxxxxxx",
      "status": 5
    },
    ...
  ]
}
```
