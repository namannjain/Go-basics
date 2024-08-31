package main

import (
	"context"
	"fmt"
	"log"
	"time"

	// "goguru/database"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Factory struct {
	name        string
	mongoClient *mongo.Client
	sqlClient   *gorm.DB
}

type FactoryMethod interface {
	SetName(name string)
	GetName() string
	GetMongoClient() *mongo.Client
	GetSqlClient() *gorm.DB
}

func (g *Factory) SetName(name string) {
	g.name = name
}

func (g *Factory) GetName() string {
	return g.name
}

func (g *Factory) GetMongoClient() *mongo.Client {
	return g.mongoClient
}

func (g *Factory) GetSqlClient() *gorm.DB {
	return g.sqlClient
}

type sql struct {
	Factory
}

type mongodb struct {
	Factory
}

func mongoDbConnection() FactoryMethod {
	uri := "localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri)))
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
	return &mongodb{Factory: Factory{name: "mongo", mongoClient: client}}
}

func main() {
	sql, _ := GetDb("mongo")
	printDetails(sql)
}

func printDetails(g FactoryMethod) {
	fmt.Printf("\n")
	fmt.Printf("db: %s", g.GetName())
	if g.GetName() == "sql" {
		fmt.Printf("\n")
		fmt.Printf("client: %v", g.GetSqlClient())
		fmt.Printf("\n")
	} else {
		fmt.Printf("\n")
		fmt.Printf("client: %v", g.GetMongoClient())
		fmt.Printf("\n")
	}
}

func sqlConnection() FactoryMethod {
	dsn := `host=localhost
          user=test1
          password=password
          port=5432
          sslmode=disable
          TimeZone=Asia/Shanghai`
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error occured")
	}
	return &sql{
		Factory: Factory{
			name:      "sql",
			sqlClient: db,
		},
	}
}

func GetDb(dbType string) (FactoryMethod, error) {
	if dbType == "sql" {
		return sqlConnection(), nil
	}
	if dbType == "mongo" {
		return mongoDbConnection(), nil
	}

	return nil, fmt.Errorf("Wrong db type passed")
}
