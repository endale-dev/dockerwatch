package main



import(
	"context"
	"fmt"
	"log"
	"github.com/moby/moby/client"
)




func main(){

	dockerClient, err:=client.New(client.FromEnv)

	if err!=nil{
		log.Fatal(err)
	}

	defer dockerClient.Close()

	result,err:=dockerClient.ContainerList(context.Background(), client.ContainerListOptions{All:true})

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("ID\tSTATUS\tIMAGE")

	for _, container:=range result.Items{

		fmt.Printf(
			"%s\t%s\t%s\n",
			container.ID[:12],
			container.Status,
			container.Image,
		)
	}
}