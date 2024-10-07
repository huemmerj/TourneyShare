package models

type Tournament struct {
	Id       string `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	TeamSize int    `bson:"team_size"`
	PublicId string `bson:"public_id"`
}
