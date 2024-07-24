package main
import (
	"context"
	"fmt"
    "os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
	"log"
)
  
  func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		log.Println(err)
	}
	DATABASE_URL := os.Getenv("DBURL")
	log.Println(os.Getenv("DBURL"))
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(DATABASE_URL).SetServerAPIOptions(serverAPI)
  
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
	  panic(err)
	}
  
	defer func() {
	  if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	  }
	}()
  
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
	  panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
  }
  