package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/MalickSidi/mvc01/services"
	"github.com/MalickSidi/mvc01/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdParam := r.URL.Query().Get("user_id")
	log.Printf("Processing user id: %v", userIdParam)

	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		// Return Bad Request
		apiErr := &utils.ApplicationError{
			Message:    "User id Must be a number\n",
			StatusCode: http.StatusBadRequest,
			Code:       "Bad_request",
		}
		jsonRes, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonRes)
		return
	}

	id := uint64(userId)
	user, apiErr := services.GetUser(id)
	if apiErr != nil {
		res, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}
