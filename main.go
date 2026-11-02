package main 



import(
	"fmt"
	"context"
	"log"
	"github.com/moby/moby/client"
	"github.com/moby/moby/api/types/container"
	
)



type DockerClient struct{
	client *client.Client
}


func NewDockerClient()(*DockerClient, error){
	c,err:=client.New(client.FromEnv)

	if err!=nil{
		return nil,err
	}

	return &DockerClient{
		client:c,
	}, nil
}


func (d *DockerClient) ListContainers(ctx context.Context) ([]container.Summary, error){

	result,err:=d.client.ContainerList(ctx,client.ContainerListOptions{All:true})
	if err!=nil{
		return nil, err
	}

	return result.Items, nil
} 


func main(){
    var ctx context.Context =context.Background()
	docker,err:=NewDockerClient()
	if err!=nil{
		log.Fatal(err)
	}

	res, err:=docker.ListContainers(ctx)
    

	fmt.Println("ID    Status   Image")
	for _,container:=range res{
		fmt.Printf(
			"%s\t%s\t%s\n",
			container.ID[:12],
			container.Status,
			container.Image,
		)
	}

}