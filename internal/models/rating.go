package models

import "time"

type Rating struct {
	Id        string    `bson:"_id"`
	Game      string    `bson:"gameName"`
	Rating    int       `bson:"rating"`
	Feedback  string    `bson:"feedback"`
	CreatedAt time.Time `bson:"createdAt"`
}

type RatingDTO struct {
	Game     string `json:"gameName"`
	Rating   int    `json:"rating"`
	Feedback string `json:"feedback"`
}
