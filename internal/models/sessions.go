package models

import "time"

type Sessions struct {
	Id           string    `bson:"_id"`
	Game         string    `bson:"gameName"`
	Timespent    float32   `bson:"timespent"`
	Deaths       int       `bson:"deaths"`
	ColorPicked  int       `bson:"colorPicked"`
	EnemysKilled int       `bson:"enemysKilled"`
	GameFinished bool      `bson:"gameFinished"`
	Money        float32   `bson:"money"`
	Ammunation   int       `bson:"ammunation"`
	Items        bool      `bson:"items"`
	CreatedAt    time.Time `bson:"createdAt"`
}

type SessionsDTO struct {
	Game         string  `json:"gameName"`
	Timespent    float32 `json:"timespent"`
	Deaths       int     `json:"deaths"`
	ColorPicked  int     `json:"colorPicked"`
	EnemysKilled int     `json:"enemysKilled"`
	GameFinished bool    `json:"gameFinished"`
	Money        float32 `json:"money"`
	Ammunation   int     `json:"ammunation"`
	Items        bool    `json:"items"`
}
