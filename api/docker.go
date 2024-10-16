package docker

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"log"
)

// Struct with all data of a Docker container
type Container struct {
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
// Returns an array of the []Container
func GetContainers() []Container {
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

	var containerList []Container

	for _, element := range containers {
		containerList = append(containerList, Container{
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
func (c Container) Stop() {
	err := GetClient().ContainerStop(context.Background(), c.ID, container.StopOptions{})
	if err != nil {
		return
	}
}

func (c Container) Start() {
	err := GetClient().ContainerStart(context.Background(), c.ID, container.StartOptions{})
	if err != nil {
		return
	}
}

func (c Container) Pause() {
	err := GetClient().ContainerPause(context.Background(), c.ID)
	if err != nil {
		return
	}
}

func (c Container) Resume() {
	err := GetClient().ContainerUnpause(context.Background(), c.ID)
	if err != nil {
		return
	}
}

func (c Container) IsRunning() bool {
	if c.State == "running" {
		return true
	}
	return false
}
