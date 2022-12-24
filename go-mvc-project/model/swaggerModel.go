package model

/* 스웨거 표시용 모델 */

/* ------- 주문자 ------- */
/* 주문 생성 요청 */
type ApiMenuCreateRequest struct {
	MenuName string `json:"menuname" bson:"menuname" example:"떡볶이"`
	Phone    string `json:"phone" bson:"phone" example:"01020221225"`
	Address  string `json:"address" bson:"address" example:"서울시 강남구 위메이드동 1212-202 101호"`
}

/* 주문 생성 응답 */
type ApiMenuCreateResponse struct {
	OrderNumber int `json:"주문번호" example:"15"`
}

/* 메뉴 리스트 조회 */
type ApiMenuListResponse struct {
	MenuList []Menu
}

/* 메뉴 상세 보기 응답 */
type ApiGetMenuDetailResponse struct {
	AvgGrade   float64 `json:"avgGrade" example:"3.5"`
	ReviewList []Review
}

/* 메뉴 변경에 대한 요청 Body */
type ApiUpdateOrderRequest struct {
	Menuname string `json:"menuname" example:"우동"`
}

/* 메뉴 변경에 대한 신규주문 응답 */
type ApiUpdateOrderResponse struct {
	Message     string `json:"message" example:"해당 주문은 배달중입니다. 신규주문으로 처리되었습니다."`
	OrderNumber int    `json:"ordernumber" example:"17"`
}

/* 주문 내역 조회 */
type ApiGetOrdersResponse struct {
	CurrentOrders []Review
	PastOrders    []Review
}

/* 리뷰 작성 요청 Body */
type ApiCreateReviewRequest struct {
	Grade   float64 `json:"grade" example:"3.5"`
	Comment string  `json:"comment" example:"맛있어요"`
}

/* ------ 피주문자 ------- */
/* 메뉴 생성 요청 Body */
type ApiCreateMenuRequest struct {
	Menuname    string `json:"menuname" example:"우동"`
	Orderstatus bool   `json:"orderstatus" example:"false"`
	LimitOrder  int    `json:"limitorder" example:"300"`
	Origin      string `json:"origin" example:"국내산"`
	Price       int    `json:"price" example:"15000"`
	Spice       int    `json:"sprice" exampel:"3"`
}

/* 메뉴  변경 요청 Body */
type ApiUpdateMenuRequest struct {
	Orderstatus bool   `json:"orderstatus" example:"false"`
	LimitOrder  int    `json:"limitorder" example:"300"`
	Origin      string `json:"origin" example:"국내산"`
	Price       int    `json:"price" example:"15000"`
	Spice       int    `json:"sprice" exampel:"3"`
}

/* 추천 메뉴 변경후 응답 */
type ApiUpdateMenuRecommend struct {
	MenuList []Menu
}

/* 주문 리스트 조회 */
type ApiGetOrderListResponse struct {
	Orders []Order
}

/* 각 메뉴별 주문 상태 변경 */
type ApiUpdateOrderStatusRequest struct {
	Status int `json:"status" example:"0"`
}
