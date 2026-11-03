package server

import "dockerwatch/docker"

type Server struct {
	dockerClient *docker.DockerClient
}


func NewServer(dockerClient *docker.DockerClient)*Server{
	return&Server{dockerClient: dockerClient,}
}