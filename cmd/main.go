package main

import (
	"log"
	"os"
	"time"

	datetimeclient "github.com/codescalersinternships/datetime-client-nabil/pkg"
)

func main() {
	// myClient := datetimeclient.NewClient("http://localhost:8090", time.Duration(1)*time.Second)

	// data, err := myClient.GetCurrentDate()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(data)
	os.Setenv("mybaseurl", "http://localhost:8090")
	myClient, err := datetimeclient.NewClientUsingEnv("mybaseurl", time.Duration(1)*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	data, err := myClient.GetCurrentDate()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(data)
}
