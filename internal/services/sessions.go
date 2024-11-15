package services

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/FromhelStudio/fromhel-session-feedbacks/internal/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SessionsService struct {
	client *mongo.Client
	ctx    *context.Context
}

func NewSessionService(mongoUri string, ctx context.Context) (*SessionsService, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}

	return &SessionsService{
		client: client,
		ctx:    &ctx,
	}, nil
}

func (s *SessionsService) CreateSession(session models.SessionsDTO) error {
	var collection *mongo.Collection

	gameName := strings.Trim(strings.ToLower(session.Game), " ")

	if gameName == "bulletspeel" {
		collection = s.client.Database("bulletspeel_db").Collection("sessions")
	} else {
		collection = s.client.Database("cordel_db").Collection("sessions")
	}

	ctx, cancel := context.WithTimeout(*s.ctx, 10*time.Second)
	defer cancel()

	model := models.Sessions{
		Id:           uuid.NewString(),
		Game:         gameName,
		Timespent:    session.Timespent,
		Deaths:       session.Deaths,
		ColorPicked:  session.ColorPicked,
		EnemysKilled: session.EnemysKilled,
		GameFinished: session.GameFinished,
		Money:        session.Money,
		Ammunation:   session.Ammunation,
		Items:        session.Items,
		CreatedAt:    time.Now(),
	}

	_, err := collection.InsertOne(ctx, model)
	return err
}

func (s *SessionsService) GetSession(gameName string, page int64) ([]models.Sessions, error) {
	var collection *mongo.Collection

	if gameName == "bulletspeel" {
		collection = s.client.Database("bulletspeel_db").Collection("sessions")
	} else {
		collection = s.client.Database("cordel_db").Collection("sessions")
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

	var sessions []models.Sessions
	for cursor.Next(ctx) {
		var session models.Sessions
		if err := cursor.Decode(&session); err != nil {
			log.Println(err)
			continue
		}
		sessions = append(sessions, session)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return sessions, nil
}
