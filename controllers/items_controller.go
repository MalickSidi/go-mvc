package controllers

import (
	"net/http"
	"strconv"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("item_id")
	itemId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {

	}
}
