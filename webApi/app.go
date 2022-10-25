package webapi

import (
	"fmt"
	"net/http"
	"sur/application/commands"
	"sur/infrastructure/adapters"
	"sur/webApi/endpoints"

	"github.com/gorilla/mux"
)

func Run() {
	// 1. Create handlers
	dbStub := adapters.NewUrlRepositoryStub()
	createUrlHandler := endpoints.NewCreateUrl(*commands.NewCreateUrlHandler(dbStub))
	getUrlHandler := endpoints.NewGetUrl(*commands.NewReadUrlHandler(dbStub))

	// 2. Set Routers
	router := mux.NewRouter()
	router.HandleFunc("/short-url-open", createUrlHandler.Handler).Methods(http.MethodPost).Name("Short Url")
	router.HandleFunc("/{id}", getUrlHandler.Handler).Methods(http.MethodGet).Name("Short Url")
	// 3. open connection
	fmt.Println("listening on: 8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		fmt.Println(err.Error())
	}

}
