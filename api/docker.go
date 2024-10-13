package docker

import "fmt"

type container struct {
	image string
	id    string
}

func Init() {
	fmt.Println("Hi")
}

func GetDockerContainers() []container {
	return []container{}
}
