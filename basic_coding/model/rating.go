package model

import (
	"github.com/jinzhu/gorm"
)

type UserReview struct {
	gorm.Model
	OrderId   int    `json:"order_id" validate:"required"`
	ProductId int    `json:"product_id" validate:"required"`
	UserId    int    `json:"user_id" validate:"required"`
	Rating    int    `json:"rating" validate:"required,gte=1,lte=5"`
	Review    string `json:"review"`
}
