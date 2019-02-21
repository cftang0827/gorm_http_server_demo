package main

// https://cizixs.com/2016/08/17/golang-http-server-side/

import (
	"fmt"
	"gorm_http_server_demo/router"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// db := model.Init()
	// defer db.Close()
	fmt.Println("Start server, binding port 8080 ...")

	fmt.Println("Test gorm ....")

	router.Route(mux)
	http.ListenAndServe(":8080", mux)
}
