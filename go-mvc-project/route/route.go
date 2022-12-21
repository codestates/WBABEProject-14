package route

import (
	"fmt"
	"wba/go-mvc-procjet/logger"

	"lecture/go-swag/docs"
	ctl "wba/go-mvc-procjet/controller"

	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
)

type Router struct {
	ct *ctl.Controller
	oc *ctl.OrdererController
	tc *ctl.TakerController
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct: ctl} //controller 포인터를 ct로 복사, 할당

	return r, nil
}

// cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
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
	// e.Use(gin.Logger())
	// e.Use(gin.Recovery())
	// 기존의 logger, recovery 대신 logger에서 선언한 미들웨어 사용
	e.Use(logger.GinLogger())
	e.Use(logger.GinRecovery(true))
	e.Use(CORS())

	logger.Info("start server")
	e.GET("/health")

	//피주문자
	taker := e.Group("api/v01/taker", liteAuth())
	//swagger 핸들러 미들웨어에 등록
	e.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	docs.SwaggerInfo.Host = "localhost:8080" //swagger 정보 등록
	{
		taker.POST("/menu", p.tc.CreateMenu)          // 메뉴 생성
		taker.PUT("/menu", p.tc.UpdateMenu)           // 메뉴 수정
		taker.DELETE("/menu", p.tc.DeleteMenu)        // 메뉴 삭제
		taker.GET("/all-order", p.tc.GetAllOrder)     // 현재 주문내역 리스트 조회
		taker.PATCH("/order", p.tc.UpdateOrderStatus) // 주문별 상태 변경

	}
	//주문자
	orderer := e.Group("api/v01/orderer", liteAuth())
	{
		orderer.GET("/menu", p.oc.CreateOrder)      // 주문 생성
		orderer.GET("/review", p.oc.GetReview)      // 리뷰 조회
		orderer.POST("/review,", p.oc.CreateReview) // 리뷰 생성
		orderer.POST("/order", p.oc.CreateOrder)    // 주문 생성
		orderer.PUT("/order", p.oc.UpdateOrder)     // 주문 변경
	}
	return e
}
