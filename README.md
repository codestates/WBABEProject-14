# WBABEProject-14

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

