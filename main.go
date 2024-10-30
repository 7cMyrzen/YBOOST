package main

import (
	db "YBOOST/DB"
	handlers "YBOOST/Handlers"

	//colors "YBOOST/Lib/Func/Colors"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDatabase()

	//g, b, re, d := colors.GetTColors()

	// Initialiser le routeur Gorilla Mux
	r := mux.NewRouter()

	staticDir := "/Static/"
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("Static"))))

	// Gestionnaire pour la route d'accueil
	r.HandleFunc("/", handlers.Home)

	// Gestionnaire pour l'API
	// r.HandleFunc("/api", handlers.API)

	// Gestionnaire pour les cocktails
	r.HandleFunc("/cocktails", handlers.Cocktails)

	// Démarrer le serveur
	http.ListenAndServe(":80", r)
}
