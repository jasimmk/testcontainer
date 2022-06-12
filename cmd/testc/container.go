package main

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
)

func StartSampleContainer() (testcontainers.Container, error) {
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
	if err != nil {
		return nil, err
	}
	fmt.Printf("Container started: %+v", redisC)
	return redisC, nil
}
