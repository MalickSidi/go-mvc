package services

import (
	"github.com/MalickSidi/mvc01/domain"
	"github.com/MalickSidi/mvc01/utils"
)

func GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
