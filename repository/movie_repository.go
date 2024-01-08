package repository

import (
	"context"
	"log"
	"os"
	"scrape-film/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MovieRepository struct {
	client       *mongo.Client
	databaseName string
}

func NewMovieRepository() *MovieRepository {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return &MovieRepository{
		client:       client,
		databaseName: os.Getenv("DB_FILM"),
	}
}

func (r *MovieRepository) getCollection(collectionName string) *mongo.Collection {
	return r.client.Database(r.databaseName).Collection(collectionName)
}

func (r *MovieRepository) Save(movie model.Movie) error {
	_, err := r.getCollection("movie").InsertOne(context.Background(), movie)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *MovieRepository) FindAll() ([]model.Movie, error) {
	var movies []model.Movie
	cursor, err := r.getCollection("movie").Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &movies)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return movies, nil
}
