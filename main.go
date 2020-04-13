package main

import (
	. "chitchat/routes"
	"log"
	"net/http"
)

func main() {
	startWebServer("8080")
}

func startWebServer(port string) {
	r := NewRouter()
	// 静态资源
	assets := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r) // 通过 router.go 中定义的路由器来分发请求

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting http listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
