package services

import (
	"go-micro/mvc/domain"
	"go-micro/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
