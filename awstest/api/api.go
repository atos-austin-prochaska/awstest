package api

import (
	A "awstest/api/aws"
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func Run() {

	//change this
	A.ShowBuckets()
	A.ShowKeys()
	//A.ShowObjects()
	fmt.Println("test")

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
