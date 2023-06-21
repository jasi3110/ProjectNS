package main

import (
	"OnlineShop/routers"
	"OnlineShop/utls"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	//initialising mux for api
	r := routers.InitRoutes()
	// initialising negroni for handling url
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(corsMiddleware))
	n.With()
	n.UseHandler(r)

	//Read IP and Port from an external JSOn file
	serverip, serverport := utls.LoadConfiguration()
	//Creating an HTTP Server using IP and Port
	server := &http.Server{
		Addr:    serverip + ":" + serverport,
		Handler: n,
	}
	log.Println("Listening on:" + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("error is founded :",err)
	}
}
func corsMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle preflight OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}

	// Call the next handler
	next(w,r)
} 