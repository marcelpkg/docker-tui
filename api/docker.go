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

type tuiContainer struct {
	Image   string
	ID      string
	Created int64
	Status  string
	State   string
}

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

func (d Docker) GetDockerContainers() {
	containers, err := d.Cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	var containerList []tuiContainer
	for _, element := range containers {
		containerList = append(containerList, tuiContainer{
			Image:   element.Image,
			ID:      element.ID,
			Created: element.Created,
			Status:  element.Status,
			State:   element.State,
		})

	}
}
