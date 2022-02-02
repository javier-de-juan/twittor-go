package routers

import (
	"fmt"
	"net/http"
)

func responseOnError(writer http.ResponseWriter) {
	recoverReason := recover()

	if recoverReason != nil {
		errorMessage := fmt.Sprintf("%v", recoverReason)
		http.Error(writer, "Bad requestModel: "+errorMessage, http.StatusBadRequest)
		return
	}
}
