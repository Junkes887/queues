package utils

import (
	"fmt"
	"net/http"
)

func ErrorMessage(err error) {
	if err != nil {
		fmt.Println(err)

	}
}

func ErrorHttpStatusInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
