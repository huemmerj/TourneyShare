package handlers

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"github.com/huemmerj/TourneyShare/db"
	"github.com/huemmerj/TourneyShare/layouts"
	"github.com/huemmerj/TourneyShare/middleware"
	"github.com/huemmerj/TourneyShare/models"
	"github.com/huemmerj/TourneyShare/pages"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"strconv"
)

func AddTournamentHandler() http.Handler {
	return middleware.Layout(templ.Handler(layouts.Default(pages.AddTournament())))
}

func AddTournamentSubmitHandler(w http.ResponseWriter, r *http.Request) {
	db := db.GetDB()
	coll := db.Collection("tournament")

	err := r.ParseForm()
	if err != nil {
		log.Fatal("Unable to parse form")
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
	}
	teamSize, err := strconv.Atoi(r.FormValue("teamSize"))
	if err != nil {
		log.Fatal(err)
	}
	name := r.FormValue("name")

	log.Print(name)
	log.Print(teamSize)

	newTournament := models.Tournament{Name: name, TeamSize: teamSize}
	results, err := coll.InsertOne(context.TODO(), newTournament)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(results)
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var tournaments []models.Tournament
	if err = cursor.All(context.TODO(), &tournaments); err != nil {
		log.Fatal(err)
	}

	handler := middleware.Layout(templ.Handler(layouts.Default(pages.Home(tournaments))))


	if handler != nil {
		w.Header().Set("HX-Push-Url", "/new-page")
		handler.ServeHTTP(w, r)
	}
}
