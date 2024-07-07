package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongo/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = ""

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// Connect with mongodb

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance

	fmt.Println("Collection instance is ready")
}

// Mongo helper

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a movie with id - ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with count", deleteCount)
}

func deleteAllMovie() {
	filter := bson.D{{}}
	deleteCount, err := collection.DeleteMany(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("No ofmovies deleted ", deleteCount.DeletedCount)
}

func getAllMovie() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for curr.Next(context.Background()) {
		var movie bson.M
		err := curr.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer curr.Close(context.Background())
	fmt.Println(movies)
	return movies
}

// Actual controller that we are going to use above arre for MONGODB
func GetAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "x-www-urlencode")
	allMovies := getAllMovie()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "x-www-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "x-www-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
