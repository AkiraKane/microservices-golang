package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AkiraKane/microservices-golang/mvc/services"
	"github.com/AkiraKane/microservices-golang/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {

	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		// Handle the err and return to the client
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	// Return user to the client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
