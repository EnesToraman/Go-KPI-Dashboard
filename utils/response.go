package utils

import (
	"encoding/json"
	"net/http"

	"github.com/EnesToraman/Go-KPI-Dashboard/model"
)

func RespondWithError(err error, statusCode int, w http.ResponseWriter) string {
	body := model.ErrorBody{
		Error:  err.Error(),
		Status: statusCode,
	}
	content, err := json.Marshal(body)
	if err != nil {
		return ""
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(content)
	return string(content)
}

func RespondWithSuccess(response interface{}, w http.ResponseWriter) string {
	content, err := json.Marshal(response)
	if err != nil {
		RespondWithError(err, http.StatusInternalServerError, w)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	w.Write(content)
	return string(content)
}
