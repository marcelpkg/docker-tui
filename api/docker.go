package docker

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"log"
)

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

func GetClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	return cli
}

// Get current running Docker containers
// This method also saves all the containers to the Docker struct, if ever needed
// Returns an array of the []tuiContainer
func GetContainers() []tuiContainer {
	d := GetClient()
	defer func(d *client.Client) {
		err := d.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(d)

	containers, err := d.ContainerList(context.Background(), container.ListOptions{})
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

	return containerList
}

// c.Stop()
func (c tuiContainer) Stop() {
	err := GetClient().ContainerStop(context.Background(), c.ID, container.StopOptions{})
	if err != nil {
		return
	}
}
