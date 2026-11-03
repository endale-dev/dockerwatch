package docker


import(
	"github.com/moby/moby/client"
)

type DockerClient struct{
	client *client.Client
}