package controllers

import (
	"github.com/huemmerj/TourneyShare/models"

	"context"
	"github.com/huemmerj/TourneyShare/db"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"math/rand"
	"time"
)

func GetTournament(id string) models.Tournament {
	db := db.GetDB()
	coll := db.Collection("tournament")

	var tournament models.Tournament

	// Create a filter for the tournament with the given ID
	filter := bson.M{"public_id": id}

	err := coll.FindOne(context.TODO(), filter).Decode(&tournament)
	if err != nil {
		log.Printf("error finding tournament: %v", err)
	}
	return tournament
}

func CreateTournament(tournament models.Tournament) {
	db := db.GetDB()
	coll := db.Collection("tournament")

	results, err := coll.InsertOne(context.TODO(), tournament)

	if err != nil {
		log.Fatal(err)
	}
	log.Print(results)
}
func GenerateRandomID() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	id := make([]byte, 5)
	for i := range id {
		id[i] = letters[rng.Intn(len(letters))]
	}
	return string(id)
}
func GenerateUniqueID() (string, error) {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	db := db.GetDB()
	coll := db.Collection("tournament")
	for {
		// Generate a random ID of 5 letters
		id := make([]byte, 6)
		for i := range id {
			id[i] = letters[rng.Intn(len(letters))]
		}
		generatedID := string(id)

		// Check if the generated ID already exists in the database
		count, err := coll.CountDocuments(context.TODO(), bson.M{"public_id": generatedID})
		if err != nil {
			log.Printf("Error checking ID in the database: %v", err)
			return "", err
		}

		// If the ID is unique, return it
		if count == 0 {
			return generatedID, nil
		}

		// Otherwise, loop again to generate a new ID
	}
}
