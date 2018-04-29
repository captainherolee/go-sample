package main

import (
	"fmt"
 	"github.com/samalba/dockerclient"
 	"log"
)

func main() {
	fmt.Println("Start Docker Client")

	//docker, _ := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	docker, _ := dockerclient.NewDockerClient("tcp://0.0.0.0:2375", nil)
	containers, err := docker.ListContainers(false, false, "")

	if err != nil {
		log.Fatal(err)
	}

	if len(containers) > 0 {
		id := containers[0].Id
		info, _ := docker.InspectContainer(id)
		log.Println(info)
	}
}
