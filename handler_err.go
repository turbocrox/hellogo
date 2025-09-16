package main

import (
	"fmt"
	"net/http"
)

func handlerErr(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 400, "Internal Server Errr")
	fmt.Println("error  handler  called")
}
