package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Определяем endpoints
	app := &App{}
	app.FileInfoMap = make(map[string]FileInfo)
	app.FileBlockIPs = make(map[int32]map[int32][]string)
	app.FileIdLocalIdMap = make(map[int32]string)

	router.HandleFunc("/add", app.AddNewFile).Methods("POST")
	router.HandleFunc("/", app.GetFiles).Methods("GET")
	router.HandleFunc("/enable", app.EnableIPForFile).Methods("POST")
	router.HandleFunc("/getBlock", app.GetBlock).Methods("POST")

	port := 8000
	fmt.Printf("Server listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
