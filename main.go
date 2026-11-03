package main 


import (
	"net/http"
	"fmt"
)


func AllContainersListhandler(w http.ResponseWriter, r*http.Request){

}

func main(){

    http.HandleFunc("/", AllContainersListhandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server Running on port 8080")

}