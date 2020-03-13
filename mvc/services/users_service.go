package services

import (
	"github.com/AkiraKane/microservices-golang/mvc/domain"
	"github.com/AkiraKane/microservices-golang/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
