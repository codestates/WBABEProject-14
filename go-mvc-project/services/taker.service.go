package services

import "wba/go-mvc-procjet/model"

type TakerService interface {
	CreateMenu(*model.Menu) error                           //피주문자 - 메뉴 생상
	UpdateMenu(string, *model.Menu) error                   //피주문자 - 메뉴 수정
	DeleteMenu(*model.Menu) error                           //피주문자 - 메뉴 삭제
	UpdateMenuRecommend(*model.Menu) ([]*model.Menu, error) //피주문자 - 오늘의 추천 메뉴 수정
	GetOrderList() ([]*model.Order, error)                  //피주문자 - 현재 주문내역 리스트 조회
	UpdateOrderStatus(string, model.Status) error           //피주문자 - 메뉴별 주문 요청 상태변경 (접수, 접수취소, 조리중, 추가주문, 배달중)
}
