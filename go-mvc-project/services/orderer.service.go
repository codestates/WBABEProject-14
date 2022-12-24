package services

import "wba/go-mvc-procjet/model"

type OrdererService interface {
	CreateOrder(*model.Order) (int, error)                          //주문자 - 주문생성
	GetAllMenu(sort string) ([]model.Menu, error)                   //주문자 - 메뉴 리스트 조회
	GetAllReiview(menuName string) (float64, []model.Review, error) //주문자 - 특정 메뉴의 모든 리뷰 보기
	CreateReview(*model.Review, string) error                       //주문자 - 과거 주문 내역중, 리뷰 작성
	UpdateOrder(string, int, string) (int, error)                   //주문자 - 주문한 메뉴 변경
	GetOrders() ([]model.Order, []model.Order, error)               //주문자 - 현재 주문 내역 조회 - 구현중
}
