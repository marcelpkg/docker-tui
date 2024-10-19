package docker

import (
	"context"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
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
// Returns an array of the []Container
func GetContainers(showAll bool) []Container {
	d := GetClient()
	defer d.Close()

    containers, err := d.ContainerList(context.Background(), container.ListOptions{
        All: showAll,
    })
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

func (c Container) Restart() {
	err := GetClient().ContainerRestart(context.Background(), c.ID, container.StopOptions{})
	if err != nil {
		return
	}
}

func (c Container) IsRunning() bool {
	return c.State == "running"
}

func (c Container) Rename(name string) {
	err := GetClient().ContainerRename(context.Background(), c.ID, name)
	if err != nil {
		return
	}
}

func (c Container) Attach() {
	GetClient().ContainerAttach(context.Background(), c.ID, container.AttachOptions{})
}
