package docker

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"log"
)

type Docker struct {
	Cli        *client.Client
	Containers []tuiContainer
}

// Struct with all data of a Docker container
type tuiContainer struct {
	Names   []string
	Image   string
	ID      string
	Created int64
	Status  string
	State   string
}

// Get an instance of the Docker struct

// Example:
// d := InitDocker()
func InitDocker() Docker {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	return Docker{
		Cli:        cli,
		Containers: []tuiContainer{},
	}
}

// Get current running Docker containers
// This method also saves all the containers to the Docker struct, if ever needed
// Returns an array of the []tuiContainer
func (d *Docker) GetContainers() []tuiContainer {
	containers, err := d.Cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	var containerList []tuiContainer

	for _, element := range containers {
		containerList = append(containerList, tuiContainer{
			Names:   element.Names,
			Image:   element.Image,
			ID:      element.ID,
			Created: element.Created,
			Status:  element.Status,
			State:   element.State,
		})
	}

	d.Containers = containerList
	return containerList
}
