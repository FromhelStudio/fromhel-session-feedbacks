package services

import (
	"context"
	"log"
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
	client *mongo.Client
	ctx    *context.Context
}

func NewRatingService(mongoUri string, ctx context.Context) (*RatingService, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}

	return &RatingService{
		client: client,
		ctx:    &ctx,
	}, nil
}

func (s *RatingService) CreateRating(rating *Rating) error {
	collection := s.client.Database("bulletSpeel").Collection("ratings")

	rating.CreatedAt = time.Now()
	rating.Id = uuid.New().String()

	_, err := collection.InsertOne(*s.ctx, rating)
	return err
}

func (s *RatingService) GetRatings() ([]Rating, error) {
	collection := s.client.Database("bulletSpeel").Collection("ratings")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var ratings []Rating
	for cursor.Next(ctx) {
		var rating Rating
		if err := cursor.Decode(&rating); err != nil {
			log.Println(err)
			continue
		}
		ratings = append(ratings, rating)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return ratings, nil
}

func (s *RatingService) Close() error {
	return s.client.Disconnect(*s.ctx)
}
