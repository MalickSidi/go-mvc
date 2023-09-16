package services

import (
	"net/http"

	"github.com/MalickSidi/mvc01/domain"
	"github.com/MalickSidi/mvc01/utils"
)

func GetItem(itemId uint64) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Not implementd yet",
		StatusCode: http.StatusNoContent,
		Code:       "not_implemnted",
	}
}
