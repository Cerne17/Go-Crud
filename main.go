package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Printf("Start main")

	routes := httprouter.New()

	server := http.Server{Addr: "localhost:8888", Handler: routes}
}
