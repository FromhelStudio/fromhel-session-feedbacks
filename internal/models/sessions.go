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
