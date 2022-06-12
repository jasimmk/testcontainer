package main

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
)

func main() {
	fmt.Println("hello")
	dockerURL := os.Getenv("DOCKER_URL")

	if dockerURL == "" {
		dockerURL = "localhost"
	}

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379:6379"},
		SkipReaper:   false,
		WaitingFor:   wait.ForListeningPort("6379"),
		AutoRemove:   true,
	}
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	fmt.Println("Container started")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(redisC)
}
