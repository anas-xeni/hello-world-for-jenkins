package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello from Xeni")
}

func main(){
	http.HandleFunc("/",helloHandler)

	server:=&http.Server{
		Addr:":8080",
	}
	fmt.Println("Server is listening on port :8080")
	err:=server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Server error: %s\n",err)
	}

}