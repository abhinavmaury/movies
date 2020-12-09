package routes

import (
	"movies/handler"
	"net/http"
)

func InitRoutes() {
	apiGroup := "/api/movies"
	adminGroup := apiGroup + "/admin"
	http.HandleFunc(adminGroup+"/add", handler.AddMovie)
	http.HandleFunc(adminGroup+"/populateDB", handler.ProcessMovies)
	http.HandleFunc(adminGroup+"/edit", handler.EditMovie)
	http.HandleFunc(adminGroup+"/remove", handler.RemoveMovie)

	userGroup := apiGroup + "/user"
	http.HandleFunc(userGroup+"/view", handler.ViewMoviesList)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
