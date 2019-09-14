package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

const (
	STATIC_DIR = "/static/"
)

type Handler struct {
	Tmpl *template.Template
}

func (h *Handler) Main(w http.ResponseWriter, r *http.Request) {
	err := h.Tmpl.ExecuteTemplate(w, "index.html", struct{}{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getRouter() *mux.Router {
	handlers := &Handler{
		Tmpl: template.Must(template.ParseGlob("./templates/*")),
	}

	// в целям упрощения примера пропущена авторизация и csrf
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.Main).Methods("GET")
	router.PathPrefix(STATIC_DIR).
		Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("./frontend/public/"))))
	router.NotFoundHandler = http.HandlerFunc(handlers.Main)

	return router
}
