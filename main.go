package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"lesson11/handlers"
	"lesson11/middleware"
	"net/http"
	"runtime"
	"sync/atomic"
	"time"
)

var count int64

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	PORT := ":8888"
	router := mux.NewRouter()

	spa := handlers.SpaHandler{StaticPath: "build", IndexPath: "index.html"}

	//router.HandleFunc("/user/{name}", handlers.GetUserHandler).Methods("Get").Schemes("https")
	subRouter := router.PathPrefix("/user").Subrouter()
	subRouter.Use(middleware.Auth)
	subRouter.Path("/{name}").HandlerFunc(handlers.GetUserHandler).Methods("Get").Schemes("https")
	router.PathPrefix("/").Handler(spa)
	//router.Use(middleware.Auth)

	server := &http.Server{
		Addr:         PORT,
		Handler:      router,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		fmt.Println(err)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
	fmt.Printf("Request from client %s\n", r.RemoteAddr)
	atomic.AddInt64(&count, 1)
}
