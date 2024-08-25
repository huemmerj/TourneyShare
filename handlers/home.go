package handlers

import (
	"github.com/a-h/templ"
	"github.com/huemmerj/TourneyShare/layouts"
	"github.com/huemmerj/TourneyShare/middleware"
	"github.com/huemmerj/TourneyShare/pages"
	"net/http"
)

func HomeHandler() http.Handler {
	return middleware.Layout(templ.Handler(layouts.Default(pages.Home())))
}
