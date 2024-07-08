package main

import (
	"log"
	"time"

	datetimeclient "github.com/codescalersinternships/datetime-client-nabil/pkg"
)

func main() {
	myClient := datetimeclient.NewClient("http://localhost:8090/datetime", time.Duration(1)*time.Second)

	data, err := myClient.GetCurrentDate()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(data)
}
