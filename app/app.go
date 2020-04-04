package app

import (
	"net/http"
)

func ApiServer() {
	http.HandleFunc("/generator/", GeneratorHandler)
	http.ListenAndServe(":8080", nil)
}
