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

func(d *DockerClient)Close(){
	d.client.Close()
}


func main(){
    ctx:=context.Background()
	docker,err:=NewDockerClient()
	if err!=nil{
		log.Fatal(err)
	}

	defer docker.Close()

	res, err:=docker.ListContainers(ctx)
    
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("ID    Status   Image")
	for _,container:=range res{

		id:=container.ID
		if len(id)>12{
			id=id[:12]
		}
		fmt.Printf(
			"%s\t%s\t%s\n",
			id,
			container.Status,
			container.Image,
		)
	}

}