package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Rating struct {
	Id        string    `bson:"_id"`
	Rating    int       `bson:"rating"`
	Feedback  string    `bson:"feedback"`
	CreatedAt time.Time `bson:"createdAt"`
}

type RatingService struct {
	mongoUri *string
}

func NewRatingService(mongoUri *string) *RatingService {
	return &RatingService{mongoUri: mongoUri}
}

func (s *RatingService) CreateRating(rating Rating) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(*s.mongoUri))

	if err != nil {
		panic(err)
	}

	collection := client.Database("bulletSpeel").Collection("ratings")

	rating.CreatedAt = time.Now()
	rating.Id = uuid.New().String()
	collection.InsertOne(
		context.Background(),
		rating,
	)

	return nil
}

func (s *RatingService) GetRatings() ([]Rating, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(*s.mongoUri))

	if err != nil {
		panic(err)
	}

	collection := client.Database("bulletSpeel").Collection("ratings")

	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		panic(err)
	}

	var ratings []Rating

	for cursor.Next(context.Background()) {
		var rating Rating
		cursor.Decode(&rating)
		ratings = append(ratings, rating)
	}

	return ratings, nil
}
