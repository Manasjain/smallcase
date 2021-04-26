package api

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
	errors "smallcase/error"
	"smallcase/utils"
)

type StandardAPIError struct {
	DevMessage string `json:"devMessage"`
	Arg        string `json:"arg"`
}

func GenerateResponse(w http.ResponseWriter, resp interface{}, err error, httpMethod string, isRespNilAcceptable ...bool) {

	var wfErr errors.Error
	if err == nil {
		if !utils.IsNil(resp) {
			var httpStatusCode int
			switch httpMethod {
			case http.MethodGet:
				httpStatusCode = http.StatusOK
			case http.MethodPost:
				httpStatusCode = http.StatusCreated
			case http.MethodPut:
				httpStatusCode = http.StatusNoContent
			}
			log.Println("API Response Status: ", httpStatusCode)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(httpStatusCode)
			_ = json.NewEncoder(w).Encode(resp)
			return
		}

		// if err and resp both are nil, but isRespNilAccecptable is true
		if len(isRespNilAcceptable) == 1 && isRespNilAcceptable[0] {
			log.Println("API Response Status: ", http.StatusNoContent)
			w.WriteHeader(http.StatusNoContent)
			return
		}

	} else {
		var ok bool
		if wfErr, ok = err.(errors.Error); !ok {
			wfErr = errors.NewError(errors.InternalServerError, err.Error()).(errors.Error)
		}
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(wfErr.GetHTTPStatusCode())
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(wfErr)
}

func GenericRecovery(w http.ResponseWriter, tag string) {
	if r := recover(); r != nil {
		log.Println(tag+" PANIC RECOVERED: ", string(debug.Stack()))
		wfErr := errors.NewError(errors.InternalServerError, errors.InternalServerError.String()).(errors.Error)
		w.WriteHeader(wfErr.GetHTTPStatusCode())
		_ = json.NewEncoder(w).Encode(wfErr)
	}
}
