package services

import (
	"context"
	"time"

	"github.com/FromhelStudio/fromhel-session-feedbacks/internal/models"
	"github.com/google/uuid"
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
	collection := s.client.Database("bulletSpeel").Collection("sessions")

	ctx, cancel := context.WithTimeout(*s.ctx, 10*time.Second)
	defer cancel()

	model := models.Sessions{
		Id:           uuid.NewString(),
		Game:         "Bullet Speel",
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
