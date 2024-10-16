package docker

import (
	"fmt"
	"github.com/docker/docker/client"
	"log"
	"os/exec"
	"testing"
)

func buildImage() {
	fmt.Println("Building test image...")
	err := exec.Command("docker", "build", "-t", "docker-tui/test", ".").Run()
	if err != nil {
		log.Fatalf("Failed to build image: %v", err)
	}
	fmt.Println("Test image built!")
}

func runContainer() {
	fmt.Println("Attempting to start container...")
	err := exec.Command("docker", "run", "--rm", "--name", "docker-tui-test", "-d", "docker-tui/test").Run()
	if err != nil {
		log.Fatalf("Failed to run image: %v", err)
	}
	fmt.Println("Test container started!")
}

func stopContainer() {
	fmt.Println("Attempting to stop container...")
	err := exec.Command("docker", "stop", "docker-tui-test").Run()
	if err != nil {
		log.Fatalf("Failed to stop container: %v", err)
	}
	fmt.Println("Container stopped!")
}

func TestDocker(t *testing.T) {
	d := GetClient()
	defer func(d *client.Client) {
		err := d.Close()
		if err != nil {

		}
	}(d)

	var found bool

	buildImage()
	runContainer()
	defer stopContainer()

	containers := GetContainers()

	for _, container := range containers {
		if container.Names[0] == "/docker-tui-test" {
			fmt.Println("Found Container: " + container.Names[0])
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Docker container not found!")
	}
}
