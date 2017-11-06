package controllers

import (
	"net/http"

	"github.com/indaco/predix-ec-configurator/helpers"
	"github.com/indaco/predix-ec-configurator/views"
)

func NewStatic() *Static {
	return &Static{
		HomeView: views.NewMarkdownView("markdownpage"),
		Start:    views.NewView("page", "static/start"),
	}
}

type Static struct {
	HomeView *views.MarkdownView
	Start    *views.View
}

func (s *Static) Home(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	vd.Yield = PageContent{
		Body: helpers.ParseMarkdown("README.md"),
	}
	if err := s.HomeView.Render(w, vd); err != nil {
		panic(err)
	}
}
