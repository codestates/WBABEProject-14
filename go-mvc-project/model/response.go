package model

type Success struct {
	Message string `json:"message" example:"success"`
}
type Failure struct {
	Message string `json:"message" example:"false"`
}
