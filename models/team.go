package models

type Team struct {
	Id   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
}
