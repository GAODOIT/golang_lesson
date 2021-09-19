package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
func main() {
	fmt.Println("hello world!")
	helloHandler := func(w http.ResponseWriter,req *http.Request) {
		io.WriteString(w,"Hello,world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
*/

/*

func helloServer(w http.ResponseWriter, req *http.Request){
	io.WriteString(w,"hello,word!")
}
func main(){
	fmt.Println("hello world!")
	//HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	http.HandleFunc("/hello", helloServer)

	//The handler is typically nil, in which case the DefaultServeMux is used.
	//equal to http.ListenAndServe(":9090",http.DefaultServeMux);
	if err := http.ListenAndServe(":9090",http.DefaultServeMux); err!=nil{
	//if err := http.ListenAndServe(":9090",nil); err!=nil{
		log.Fatal("server start error: ", err)
	}

}
*/

func helloServer(w http.ResponseWriter, req *http.Request){
	io.WriteString(w,"hello,word!")
}
func main(){
	fmt.Println("hello world!")
	//srv := &http.Server{Addr:":9090",Handler:http.DefaultServeMux}
	srv := &http.Server{Addr:":9090"}
	http.HandleFunc("/hello", helloServer)
	fmt.Println("http server start:")
	//err := http.ListenAndServe(":9090",http.DefaultServeMux);
	if err := srv.ListenAndServe(); err!=nil{
		//if err := http.ListenAndServe(":9090",nil); err!=nil{
		log.Fatal("server start error: ", err)
	}

	srv.Shutdown(context.TODO())

}