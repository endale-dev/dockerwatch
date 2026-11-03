package main

import (
	"dockerwatch/docker"
	"dockerwatch/server"
	"fmt"
	"log"
	"net/http"
)



func main(){
	
    dockerClient,err:=docker.NewDockerClient()
	if err!=nil{
		log.Fatal("Failed to Create Docker client ",err)
	}

	srv:=server.NewServer(dockerClient)
    http.HandleFunc("/", srv.AllContainersListHandler)

	fmt.Println("Server Running on port 8080")
	http.ListenAndServe(":8080", nil)
	

}