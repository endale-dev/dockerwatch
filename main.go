package main 



import(
	"fmt"
	"context"
	"log"
	"github.com/moby/moby/client"
	
)


func main(){

	dockerClient, err:=client.New(client.FromEnv)

	if err!=nil{
		log.Fatal(err)
	}

	defer dockerClient.Close()

	result, err:=dockerClient.ContainerList(context.Background(), client.ContainerListOptions{All:true})

	for _,container:=range result.Items{

		fmt.Printf(
			"%s\t%s\t%s\n",
			container.ID,
			container.Status,
			container.Image,
	)
	}
}