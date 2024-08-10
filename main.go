package main
import (
	"context"
    "os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"io"
	"github.com/jackc/pgx/v5"
)
  
  func main() {
    env := os.Getenv("ENVIRONMENT")
	log.Println(os.Getenv("ENVIRONMENT"))

	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
			log.Println(err)
		}
	} 

	DATABASE_URL := os.Getenv("DBURL")
	log.Println(os.Getenv("DBURL"))
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
		  io.WriteString(w, "Connection failed!\n")
		  panic(err)
		}
	
		io.WriteString(w, "Pinged your deployment. You successfully connected to MongoDB!\n")
	})
	pgtest()

	log.Fatal(http.ListenAndServe(":8080", nil))
  }


func pgtest() {
	log.Println(os.Getenv("PG_DATABASE_URL"))
	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_DATABASE_URL"))
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var id int64
	var num int64
	var data string
	err = conn.QueryRow(context.Background(), "select id, num, data from test").Scan(&id, &num, &data)
	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	log.Println(id, num, data)
}
  