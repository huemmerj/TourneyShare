package handlers

import (
	"context"
	"github.com/a-h/templ"
	"github.com/huemmerj/TourneyShare/db"
	"github.com/huemmerj/TourneyShare/layouts"
	"github.com/huemmerj/TourneyShare/middleware"
	"github.com/huemmerj/TourneyShare/models"
	"github.com/huemmerj/TourneyShare/pages"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func HomeHandler() http.Handler {
	db := db.GetDB()
	coll := db.Collection("tournament")

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var tournaments []models.Tournament
	for _, tournament := range tournaments {
		log.Printf("Tournament: %+v\n", tournament)
	}
	if err = cursor.All(context.TODO(), &tournaments); err != nil {
		log.Fatal(err)
	}
	return middleware.Layout(templ.Handler(layouts.Default(pages.Home(tournaments))))
}
