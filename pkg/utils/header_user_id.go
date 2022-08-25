package utils

import (
	"net/http"
	"strconv"
)

func GetUserIDFromHeader(req *http.Request) int {

	headerUserID := req.Header.Get("user_id")

	userID, _ := strconv.Atoi(headerUserID)

	return userID

}
