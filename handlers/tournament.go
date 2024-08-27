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
	"log"
	"net/http"
	"strconv"
)

func AddTournamentHandler() http.Handler {
	return middleware.Layout(templ.Handler(layouts.Default(pages.AddTournament())))
}

func AddTournamentSubmitHandler(w http.ResponseWriter, r *http.Request) {
	handler := middleware.Layout(templ.Handler(layouts.Default(pages.Home())))

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

	db := db.GetDB()
	coll := db.Collection("tournament")

	newTournament := models.Tournament{Name: name, TeamSize: teamSize}
	results, err := coll.InsertOne(context.TODO(), newTournament)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(results)
	if handler != nil {
		fmt.Print("hallo")
		handler.ServeHTTP(w, r)
	}
}
