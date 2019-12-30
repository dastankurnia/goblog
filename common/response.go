package common

import (
	"encoding/json"
	"net/http"

	"github.com/dastankurnia/goblog/common"
)

// ResponseData : response data struct
type ResponseData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response : http response format
func Response(w http.ResponseWriter, r *http.Request, status int, message string, payloads interface{}) {
	var res ResponseData

	clientID := r.Header.Get("clientId")
	clientKey := r.Header.Get("clientKey")
	hash := common.HmacGenerate(clientID, clientKey)
	header := common.HmacValidate(clientID, hash)

	if header != true {
		status = 401
		message = "Unauthorized"
		payloads = nil
	}

	res.Status = status
	res.Message = message
	res.Data = payloads

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&res)
	if err != nil {
		panic(err)
	}
}
