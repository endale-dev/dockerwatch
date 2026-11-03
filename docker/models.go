package docker


import(
	"github.com/moby/moby/client"
	
)

type DockerClient struct{
	client *client.Client
}


type Container struct{
	ID string `json:"id"`
	Status string `json:"status"`
	Image string `json:"image"`
}