package router

import (
	"fmt"
	"gorm_http_server_demo/controller"
	"net/http"
)

// Route ...
func Route(mux *http.ServeMux) {
	fmt.Println("Get mutiplexer ..")
	mux.HandleFunc("/", controller.Welcome)
	mux.HandleFunc("/create", controller.CreateUser)
	mux.HandleFunc("/read", controller.ReadUser)
	mux.HandleFunc("/update", controller.UpdateUser)
	mux.HandleFunc("/delete", controller.DeleteUser)

}
