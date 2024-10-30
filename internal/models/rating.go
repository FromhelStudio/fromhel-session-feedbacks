package models

import "time"

type Rating struct {
	Id        string    `bson:"_id"`
	Rating    int       `bson:"rating"`
	Feedback  string    `bson:"feedback"`
	CreatedAt time.Time `bson:"createdAt"`
}

type RatingDTO struct {
	Rating   int    `json:"rating"`
	Feedback string `json:"feedback"`
}
