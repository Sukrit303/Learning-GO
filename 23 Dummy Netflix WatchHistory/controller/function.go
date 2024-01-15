package controller

import (
	"BackendNetflixWatchHistory/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const url = "mongodb+srv://sukritsahu303:iKpSQEGaRKBs8G1R@cluster0.ccyew9m.mongodb.net/"
const dbName = "netflix"
const colName = "watchHistory"

var collection *mongo.Collection

func init() {

	// Client option to setup connection info
	clientOptions := options.Client().ApplyURI(url)

	// Connection to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	// Creating Collection
	collection = client.Database(dbName).Collection(colName)

	// Collection instance created
	fmt.Println("MongoDB collection created")
}

func insertAMovie(movie models.NetflixWatchHistory) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document:", inserted)
}

func updateAMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
}

// Delete
func deleteAMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count: ", d.DeletedCount)
}

// in find all we get cursor in mongoDb
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)

	}
	defer cursor.Close(context.Background())

	return movies

}

///////////////////// Routers ????????????????????

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allMovies := getAllMovies()

	data := json.NewEncoder(w).Encode(allMovies)

	fmt.Printf("%v\n", data)
}

func CreatingMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.NetflixWatchHistory

	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertAMovie(movie)

	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateAMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}


func DeleteAMovie(w http.ResponseWriter,r *http.Request ){
	w.Header().Set("Content-Type", "application/x-ww-form-urlencode")	
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteAMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"]) 
}