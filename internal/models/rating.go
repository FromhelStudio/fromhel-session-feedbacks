package models

type RatingDTO struct {
	Rating   int    `json:"rating"`
	Feedback string `json:"feedback"`
}
