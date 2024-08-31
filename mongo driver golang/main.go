package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type manager struct {
	Connection *mongo.Client
	Ctx        context.Context
	Cancel     context.CancelFunc
}

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
}

type Manager interface {
	Insert(interface{}) error
	GetAll() ([]User, error)
	DeleteData(primitive.ObjectID) error
	UpdateData(User) error
}

var Mgr Manager

func init() {
	connectDb()
}

func main() {
	u := User{Name: "naman", Email: "naman@gmail.com"}
	Mgr.Insert(u)

	//get all records in db
	data, err := Mgr.GetAll()
	fmt.Println(data, err)

	//delete record from db
	id := "641e08889d85ada518e83ed1"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	u.ID = objectId
	u.Name = "test"
	u.Email = "test@gmail.com"
	err = Mgr.UpdateData(u)
	fmt.Println(err)
}

// context package is a  powerful tool to manage operations like timeouts, cancelation, deadlines,, etc.
// Among these operations, context with timeout is mainly used when we want to make an external request, such as a network request or a database request
func connectDb() {
	//mongo documentation says u can use 27108, 27019 also
	url := "localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", url)))
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
		return
	}

	Mgr = &manager{Connection: client, Ctx: ctx, Cancel: cancel}
	fmt.Println("Connected")
}

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	//cancelFunc to cancel the context
	defer cancel()

	//client provides a method to close a mongoDB connection
	defer func() {
		//client.Disconnect method also has deadline
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func (mgr *manager) Insert(data interface{}) error {
	orgCollection := mgr.Connection.Database("goguru123").Collection("collectiongoguru")
	result, err := orgCollection.InsertOne(context.TODO(), data)
	fmt.Println(result.InsertedID)
	return err
}

func (mgr *manager) GetAll() (data []User, err error) {

	orgCollection := mgr.Connection.Database("goguru123").Collection("collectiongoguru")

	// Pass these options to the Find method
	findOptions := options.Find()

	cur, err := orgCollection.Find(context.TODO(), bson.M{}, findOptions)
	for cur.Next(context.TODO()) {
		var d User
		err := cur.Decode(&d)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, d)
	} // close for

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return data, nil
}

func (mgr *manager) DeleteData(id primitive.ObjectID) error {
	orgCollection := mgr.Connection.Database("goguru123").Collection("collectiongoguru")

	filter := bson.D{{"_id", id}}
	_, err := orgCollection.DeleteOne(context.TODO(), filter)
	return err
}

func (mgr *manager) UpdateData(data User) error {
	orgCollection := mgr.Connection.Database("goguru123").Collection("collectiongoguru")

	filter := bson.D{{"_id", data.ID}}
	update := bson.D{{"$set", data}}

	_, err := orgCollection.UpdateOne(context.TODO(), filter, update)

	return err
}
