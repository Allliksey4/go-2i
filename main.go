package main

import (
	dockerContainer "go-2i/api/docker/container"
)

func main() {
	dockerID := dockerContainer.Create()
	dockerContainer.Start(dockerID)
	dockerContainer.Logs(dockerID)
	dockerContainer.Remove(dockerID)
}
