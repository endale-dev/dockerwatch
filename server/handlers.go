package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (srv *Server)AllContainersListHandler(w http.ResponseWriter, r *http.Request){
    
	containers, err:=srv.dockerClient.ListAllContainers(r.Context());
	if err!=nil{
		fmt.Fprintf(w, "Error fetching containers list")
		return
	}
   
	w.Header().Set("Content-Type", "application/json")
	err=json.NewEncoder(w).Encode(containers);
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}