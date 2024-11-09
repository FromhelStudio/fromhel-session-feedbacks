package services

import (
	"context"
	"log"
	"time"

	"github.com/FromhelStudio/fromhel-session-feedbacks/internal/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func (s *RatingService) CreateRating(rating *models.Rating) error {
	var collection *mongo.Collection

	if rating.Game == "bulletspeel" {
		collection = s.client.Database("bulletspeel_db").Collection("ratings")
	} else {
		collection = s.client.Database("cordel_db").Collection("ratings")
	}

	rating.CreatedAt = time.Now()
	rating.Id = uuid.New().String()

	_, err := collection.InsertOne(*s.ctx, rating)
	return err
}

func (s *RatingService) GetRatings(page int64, gameName string) ([]models.Rating, error) {
	var collection *mongo.Collection

	if gameName == "bulletspeel" {
		collection = s.client.Database("bulletspeel_db").Collection("ratings")
	} else {
		collection = s.client.Database("cordel_db").Collection("ratings")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	limit := int64(10)
	skip := int64((page - 1) * 10)

	cursor, err := collection.Find(ctx, bson.M{}, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  bson.M{"createdAt": -1},
	})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var ratings []models.Rating
	for cursor.Next(ctx) {
		var rating models.Rating
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
