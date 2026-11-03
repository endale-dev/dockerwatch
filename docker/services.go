package docker

import (
	"context"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)



func NewDockerClient()(*DockerClient, error){
	c,err:=client.New(client.FromEnv)
    
	if err!=nil{
		return nil, err
	}

	return &DockerClient{client:c,},nil
}

func (d *DockerClient)ListAllContainers(ctx context.Context)([]container.Summary, error){

	res,err:=d.client.ContainerList(ctx, client.ContainerListOptions{All:true})

	if err!=nil{
		return nil,err
	}

	return res.Items,nil
}