package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Println("Hello")
}

func TestWithRedis(t *testing.T) {
	t.Log("Starting test")
	c, err := StartSampleContainer()
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	t.Log(c)
}
