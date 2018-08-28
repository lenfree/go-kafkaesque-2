// +build !doc

package gokafkaesque_test

import (
	"fmt"

	gokafkaesque "github.com/packetloop/go-kafkaesque"
)

func ExampleNewClient() {
	config := gokafkaesque.NewConfig().
		SetURL("http://localhost:8080").
		SetRetry(3).
		Build()
	client := gokafkaesque.NewClient(config)
	a, _ := client.GetStatus()
	fmt.Println(a.GetHealth())

	// Output:
	// Ok
}