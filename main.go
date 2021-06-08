package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	// fmt.Printf("%v \n %T\n", pageID, pageID)
	fileName := "files/" + pageID + ".html"
	if pageID != "1" {
		fileName = "files/404.html"
	}
	http.ServeFile(w, r, fileName)
}
func servePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%T\t%v", vars, vars)
	pageID := vars["id"]
	fileName := "files/" + pageID + ".html"
	http.ServeFile(w, r, fileName)
}
func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	rtr.HandleFunc("/page/{id:[0-9]+}", servePage)
	http.Handle("/", rtr)
	http.ListenAndServe(PORT, nil)
}
