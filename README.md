# WBABEProject-14

# 📖 목차 
 - [소개](#소개) 
 - [사용법](#사용법)
 - [디렉토리 구조](#디렉토리-구조)
 - [개발 환경](#개발-환경)
 - [사용 기술](#사용-기술)
 - [ERD](#erd)
 - [서버 아키텍처](#서버-아키텍처) 
 - [Api 명세서](#api-명세서)
 - [피드백 개선](#피드백-개선)
 - [주요 기능](#주요-기능)
    - [피주문자](#피주문자)
        - [메뉴 생성](#메뉴-생성)
        - [메뉴 수정](#메뉴-수정)
        - [메뉴 삭제](#메뉴-삭제)
        - [추천 메뉴 변경](#추천-메뉴-변경)
        - [현재 주문내역 조회](#현재-주문내역-조회)
        - [메뉴별 주문요청 상태 변경](#메뉴별-주문요청-상태-변경)
    - [주문자](#주문자)
        - [메뉴 리스트 조회](#메뉴-리스트-조회)
        - [주문 생성](#주문-생성)
        - [주문 메뉴 변경](#주문-메뉴-변경)
        - [리뷰 작성](#리뷰-작성)
        - [메뉴 평점 리뷰 조회](#메뉴-평점-리뷰-조회)
        - [현재 주문 내역 조회](#현재-주문-내역-조회)
 
    
    
# 소개 

언택트 시대에 급증하고 있는 온라인 주문 시스템은 이미 생활전반에 그 영향을 끼치고 있는 상황에, 가깝게는 배달어플, 매장에는 키오스크, <br> 식당에는 패드를 이용한 메뉴 주문까지 그 사용범위가 점점 확대되어 가고 있습니다.<br> 이런 시대에 해당 시스템을 이해, 경험하고 각 단계별 프로세스를 이해하여 구현함으로써 서비스 구축에 경험을 쌓고, golang의 이해를 돕습니다.

1. 주문자/피주문자의 역할에서 필수적인 기능을 도출, 구현합니다.
2. 해당 시스템에 대해 요구사항을 접수하고 주문자와 피주문자 입장에서 필요한 기능을 도출하여, 기능을 확장하고 주문 서비스를 원할하게 지원할수 있는 기능을 구현합니다.
3. 주문자는 신뢰있는 주문과 배달까지를 원합니다. 또, 피주문자는 주문내역을 관리하고 원할한 서비스가 제공되어야 합니다.

### 사용법
```bash
$ docker start mongodb

$ git clone https://github.com/codestates/WBABEProject-14.git
$ cd go-mvc-project/
$ go mod tidy
> go run main.go
```

### 개발 환경
- Mac OS M1
- vscode
- MongoDB Compass
- Docker 
- PostMan
- Talen API Tester

### 사용 기술
- Go(go1.19.4 darwin/arm64)
- Gin Gonic
- MongoDB
- Swagger 

### 디렉토리 구조 

```
├── config         
│   ├── config.go  
│   └── config.toml
├── controller     
│   ├── controller.go
│   ├── orderer.controller.go /* 주문자 컨트롤러 */
│   └── taker.controller.go   /* 피주문자 컨트롤러 */
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── logger
│   └── logger.go
├── logs
│   └── go-loger_2022-12-20.log
├── main.go
├── model
│   ├── menu.go     /* 메뉴 모델 */
│   ├── model.go    /* 모델 정의 */
│   ├── order.go    /* 주문 모델 */
│   ├── response.go /* 응답 메세지 모델 */
│   ├── review.go   /* 리뷰 모델 */
│   └── swaggerModel.go /* 스웨거 작성시 요청과 응답에 대한 모델 */
├── route
│   └── route.go
└── services
    ├── orderer.service.go      /* 주문자 서비스 인터페이스 */
    ├── orderer.service.impl.go /* 주문자 서비스 구현체 */
    ├── taker.service.go        /* 피주문자 서비스 인터페이스 */
    └── taker.service.impl.go   /* 피주문자 서비스 구현체 */
```

### ERD
<img src="https://user-images.githubusercontent.com/103496262/211019892-79a95a76-e678-4d7b-ba73-82a58e2e87e9.png"/>

### 서버 아키텍처 

<img src="https://user-images.githubusercontent.com/103496262/209453624-da34cdbd-62cb-41c3-80b8-c73a3975ba3e.png"/>


### API 명세서

<details>
<summary>📋 Swagger 스크린샷 보기</summary>
<div markdown="1">

<img src="https://user-images.githubusercontent.com/103496262/210320765-7a5e36ca-efc0-44fa-9544-08f21d8d6132.png"/>

</div>
</details>


```
프로젝트 실행후 아래 도메인 으로 접속 
http://localhost:8080/swagger/index.html#/ 
```

# 피드백 개선

+ [x] [created_at, updated_at 한 세트로 구성하기](https://github.com/codestates/WBABEProject-14/commit/e756f9bf36a119a4ca6f77dc9faa1c7a7ef9a0ad)
+ [x] [주문 상태를 Enum, Const 활용하기](https://github.com/codestates/WBABEProject-14/commit/bd302a8d54df3d60602e0704e04289952b870c67)
+ [x] [사용하지 않는 주석 코드 제거](https://github.com/codestates/WBABEProject-14/commit/0dfa0822f2ba5ba0972c4f70096b727179eb83d1)
+ [x] [rest Api 에 맞게 리팩토링 URI 에 자원만 명시 ( detail 삭제 )](https://github.com/codestates/WBABEProject-14/commit/000e7fd298956740da9d7bdede8b3dc1a861ad85)
+ [x] [메뉴 추가, 변경과 같은 것을 구분하기 위해선 일반적으로 쿼리스트링으로 사용](https://github.com/codestates/WBABEProject-14/commit/5672a10c5b346c0fa9784c1c57306be2bccefa01)
+ [x] [직관적이지 못한 네이밍 수정](https://github.com/codestates/WBABEProject-14/commit/16b877c611e43f92cb02176e29c1d21e67c5fbfc)
+ [x] [함수의 코드가 길어지는 경우 하나의 책임을 가질수 있도록 분리 필요](https://github.com/codestates/WBABEProject-14/commit/4a8f002705911e9c2358c3248790f6cebdac6823)
+ [x] [메뉴 추가시 메뉴 이름만으로 중복검사 처리는 용이하지 못함](https://github.com/codestates/WBABEProject-14/commit/0c6b0285a378b5ca73bf660eed059c1e1a7bf02f)
+ [x] [CORS 허용은 특정 도메인만 적용하기](https://github.com/codestates/WBABEProject-14/blob/2c7ff2849b162ae369e5173197e121e5263c41ef/go-mvc-project/route/route.go#L34)

# 주요 기능

# 피주문자
## 메뉴 생성
[func CreateMenu()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/taker.service.impl.go#L28-L57)
### Request
`POST /api/v01/taker/menu`
```json
{
  "storenumber":100,
  "limitorder": 300,
  "menuname": "우동",
  "orderstatus": false,
  "origin": "국내산",
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
## 메뉴 수정
[func UpdateMenu()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/taker.service.impl.go#L59-L87)
### Request
`PUT /api/v01/taker/menu/우동`

```json
{
  "limitorder": 300,
  "orderstatus": false,
  "origin": "국내산",
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
## 메뉴 삭제 
[func DeleteMenu()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/taker.service.impl.go#L77-L87)
### Request
`DELETE /api/v01/taker/menu`
```json
{
  "menuname": "우동"
}
```
### Response 
```json
{
  "message": "success"
}
```
## 추천 메뉴 변경
[func UpdateMenuRecommend()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/taker.service.impl.go#L89-L112)

```bash
PATCH /api/v01/taker/menu
```
### Request
```json
{
  "menuname": "우동"
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
      "menuname": "떡볶이",
      "orderstatus": true,
      "origin": "국내산",
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
      "menuname": "김밥",
      "orderstatus": true,
      "origin": "국내산",
      "price": 5000,
      "recommend": true,
      "reorder": 50,
      "spice": 0
    }
    ...
  ]
}
```
## 현재 주문내역 조회 
[func GetOrderList()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/taker.service.impl.go#L114-L128)


`GET /api/v01/taker/orders`

### Request
### Response 
```json
{
  "orders": [
    {
      "_id": "63a73a1c8d989838729bc114",
      "address": "서울시 강남구 강남동 1212-202 101호",
      "createdat": "2022-12-24T16:17:12.793+00:00",
      "isdelete": false,
      "isexistreview": false,
      "menuname": "떡볶이",
      "ordernumber": 11,
      "phone": "010xxxxxxxx",
      "status": 0
    },
    {
      "_id": "63a79abf892e3345fd2f2150",
      "address": "서울시 용산구 용산동 2525-111 303호",
      "createdat": "2022-12-24T18:11:22.440+00:00",
      "isdelete": false,
      "isexistreview": false,
      "menuname": "우동",
      "ordernumber": 12,
      "phone": "010xxxxxxxx",
      "status": 0
    }
    ...
  ]
}
```
## 메뉴별 주문요청 상태 변경
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

# 주문자
 
## 메뉴 리스트 조회 
[func GetAllMenu()](https://github.com/codestates/WBABEProject-14/blob/444bbb6ec3a2690a55dc284f948745caca543974/go-mvc-project/services/orderer.service.impl.go#L57-L72)

### Request 
`GET /api/v01/orderer/menu/{sort}`

`sort : [recommend, grade, reorder, createdat] 중 하나 택하여 요청`

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
      "menuname": "떡볶이",
      "orderstatus": true,
      "origin": "국내산",
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
      "menuname": "우동",
      "orderstatus": true,
      "origin": "국내산",
      "price": 5000,
      "recommend": true,
      "reorder": 50,
      "spice": 0
    },
    ...
  ]
}
```

## 주문 생성
[func CreateMenu()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/orderer.service.impl.go#L35-L62)

### Request 
`POST /api/v01/orderer/order`

```json
{
  "address": "서울시 강남구 강남동 1212-202 101호",
  "menuname": "떡볶이",
  "phone": "01020221225"
}
```
### Response 
```json
{
  "주문번호": 15
}
```

## 주문 메뉴 변경
[func UpdateOrder()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/orderer.service.impl.go#L158-L182)

### Request 
`PATCH /api/v01/orderer/order/{orderId}?flag={0,1}`
```json
{
  "menuname": "우동"
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
  "message": "해당 주문은 배달중입니다. 신규주문으로 처리되었습니다.",
  "ordernumber": 17
}
```


## 리뷰 작성
[func CreateReview()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/orderer.service.impl.go#L97-L156)
### Request 
`POST /api/v01/orderer/review/{orderID}`

```json
{
  "comment": "맛있어요",
  "grade": 3.5
}
```

### Response 
```json
{
  "message": "success"
}
```

## 메뉴 평점 리뷰 조회
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
      "comment": "맛있었습니다.",
      "createdat": "2022-12-24T17:42:53.949+00:00",
      "grade": 3,
      "isdelete": false,
      "menuname": "떡볶이"
    },
    {
      "_id": "63a79abf892e3345fd2f2150')",
      "comment": "맛없어요.",
      "createdat": "2022-12-24T17:44:19.147+00:00",
      "grade": 4,
      "isdelete": false,
      "menuname": "떡볶이"
    },
    ...
  ]
}
```

## 현재 주문 내역 조회
[func Orders()](https://github.com/codestates/WBABEProject-14/blob/2e16bbfe5a251f2185d21ffc72bde57a4e57e1d8/go-mvc-project/services/orderer.service.impl.go#L184-L208)
### Request 
`GET /api/v01/orderer/orders`

### Response  

```json
{
  "currentOrders": [
    {
      "_id": "63a79abf892e3345fd2f2150",
      "address": "서울시 강남구 강남동 1212-202 101호",
      "createdat": "2022-12-24T16:17:12.793+00:00",
      "isdelete": false,
      "isexistreview": false,
      "menuname": "떡볶이",
      "ordernumber": 1,
      "phone": "010xxxxxxxx",
      "status": 0
    },
    ...
  ],
  "pastOrders": [
    {
      "_id": "63a73a1c8d989838729bc114",
      "address": "서울시 용산구 용산동 1212-202 303호",
      "createdat": "2022-12-24T16:15:44.327+00:00",
      "isdelete": false,
      "isexistreview": false,
      "menuname": "우동",
      "ordernumber": 0,
      "phone": "010xxxxxxxx",
      "status": 5
    },
    ...
  ]
}
```
