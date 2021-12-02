package api

import (
	A "awstest/api/aws"
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("testing"))
}

func accessKeys(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("access keys go here"))
}

func buckets(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buckets go here"))
}

func objects(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("objects go here"))
}

func Run() {

	//change this
	A.ShowBuckets()
	A.ShowKeys()
	//A.ShowObjects()
	fmt.Println("test")

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/ack", accessKeys)
	mux.HandleFunc("/buk", buckets)
	mux.HandleFunc("/obj", objects)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
