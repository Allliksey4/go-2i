package main

import (
	"go-2i/api/docker"
)

func main() {
	dockerID := docker.DockerCreate()
	docker.DockerStart(dockerID)
	docker.DockerLogs(dockerID)
	docker.DockerRemove(dockerID)
}
