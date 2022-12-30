package route

import (
	"fmt"
	"wba/go-mvc-procjet/docs"
	"wba/go-mvc-procjet/logger"

	ctl "wba/go-mvc-procjet/controller"

	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
)

type Router struct {
	oc *ctl.OrdererController
	tc *ctl.TakerController
}

/* 주문자, 피주문자 컨트롤러 할당 */
func NewRouter(orc *ctl.OrdererController, trc *ctl.TakerController) (*Router, error) {
	r := &Router{oc: orc, tc: trc}

	return r, nil
}

// cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
		CORS 허용을 위해서 모든 도메인을 허용한다면 보안에 이슈가 생깁니다. 
		보통 운영되는 시스템의 경우는 특정한 도메인만을 허용하고 그 이외의 요청은 거부하도록 설정합니다.
		*/
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//허용할 header 타입에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//허용할 method에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// 임의 인증을 위한 함수
func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*----------- 인증 프로세스 -----------*/
		if c == nil {
			c.Abort() // 미들웨어에서 사용, 이후 요청 중지
			return
		}
		auth := c.GetHeader("Authorization")

		if auth != "codz" {
			//로직 추가 가능 현재는 Print 로만 처리
			fmt.Println("Authorization failed")
		}
		/*--------------- END -------------*/
		c.Next()
	}
}

func (p *Router) Idx() *gin.Engine {
	// 컨피그나 상황에 맞게 gin 모드 설정
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	e := gin.Default()
	/*
	사용하지 않는 코드는 지워주시는 것이 추후에 작업할 때 헷갈리지 않습니다. 
	이러한 코드를 주석처리해두면, 추후 다른 작업자(미래의 자신 포함)가 의문을 가질수 있고 이는 곧 생산성의 저하로 이루어집니다.
	히스토리 추적은 git을 이용하면 충분합니다.
	*/
	// e.Use(gin.Logger())
	// e.Use(gin.Recovery())
	// 기존의 logger, recovery 대신 logger에서 선언한 미들웨어 사용
	e.Use(logger.GinLogger())
	e.Use(logger.GinRecovery(true))
	e.Use(CORS())

	logger.Info("start server")
	e.GET("/health")

	//swagger 핸들러 미들웨어에 등록
	e.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	docs.SwaggerInfo.Host = "localhost:8080" //swagger 정보 등록
	docs.SwaggerInfo.Title = "띵동주문이요, 온라인 주문 시스템(Online Ordering System)"
	docs.SwaggerInfo.Description = "주문자와 피주문자로 라우팅 하였습니다."
	docs.SwaggerInfo.Version = "v01"

	/*
	엔드포인트 구성에 대해서 전반적인 코멘트 드립니다.
	1. REST API의 성숙도 모델에 대해서 공부해보시면 좋을 것 같습니다.
	
	2. 일반적으로 HTTP URI에 deatil와 같은 단어는 들어가지 않습니다. 
		복수형의 단어로 구성을 하고, 동일한 URI 내에서 http method만 변경하여 행위를 표현하는 것이 일반적인 REST API의 구성 방식입니다.
		e.g.
		GET v1/orders -> 주문 목록을 조회.
		GET v1/orders/1 -> 1번 주문를 조회.
		POST v1/orders -> 주문를 생성.
		PATCH v1/orders/1 -> 1번 주문에 대해서 업데이트
		DELETE v1/orders/1 -> 1번 주문에 대해서 삭제
	*/

	/*
	그룹으로 통해 엔드포인트를 나누어서 코드를 보기에 용이합니다. 좋은 코드입니다.
	*/
	//피주문자
	taker := e.Group("api/v01/taker", liteAuth())
	{
		taker.POST("/menu", p.tc.CreateMenu)                     // 메뉴 생성
		taker.PUT("/menu/:menuname", p.tc.UpdateMenu)            // 메뉴 수정
		taker.PATCH("/menu", p.tc.UpdateMenuRecommend)           // 금일 추천 메뉴 변경
		taker.DELETE("/menu", p.tc.DeleteMenu)                   // 메뉴 삭제
		taker.GET("/orders", p.tc.GetOrderList)                  // 현재 주문내역 리스트 조회
		taker.PATCH("/orders/:menuname", p.tc.UpdateOrderStatus) // 주문별 상태 변경

	}
	//주문자
	orderer := e.Group("api/v01/orderer", liteAuth())
	{
		orderer.GET("/menu/:sort", p.oc.GetAllMenu)              // 메뉴 리스트 조회
		orderer.POST("/order", p.oc.CreateOrder)                 // 주문 생성
		orderer.POST("/review/:orderID", p.oc.CreateReview)      // 리뷰 생성
		orderer.GET("/detailMenu/:menuname", p.oc.GetMenuDetail) // 메뉴 상세보기 (평점, 리뷰 조회)
		/*
		메뉴 추가, 변경과 같은 것을 구분하기 위해선 일반적으로 쿼리스트링을 통해 받아와 사용합니다.
		혹은, 메뉴 추가와 변경에 대한 엔드포인트를 각각 생성하는 방법도 있겠습니다.
		*/
		orderer.PATCH("/order/:orderId/:flag", p.oc.UpdateOrder) // 주문 변경
		orderer.GET("/orders", p.oc.GetOrders)                   // 주문 내역 조회
	}
	return e
}
