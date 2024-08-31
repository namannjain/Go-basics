package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var lock = &sync.Mutex{}

type singleConnection struct {
	connection *mongo.Client
}

var singleInstance *singleConnection

func getInstance() *singleConnection {
	lock.Lock()
	defer lock.Unlock()
	if singleInstance == nil {
		fmt.Println("Creating db connection")
		uri := "localhost:27017"
		client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri)))
		// uri := "mongodb://localhost:27017"
		// client, err := mongo.NewClient(options.Client().ApplyURI(uri))
		if err != nil {
			fmt.Println(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			fmt.Println(err)
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Connected!!!")
		singleInstance = &singleConnection{connection: client}
	} else {
		fmt.Println("DB connection has already been created!")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}
	//scanln is similar to scan but stops scanning at a new line
	// after final item there must be a new line or EOF
	fmt.Scanln()
}
