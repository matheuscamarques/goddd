package main

import (
	"api/domain/app"
	"api/repository"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// make a db conn to the database
	db, err := sql.Open("mysql", "root:root@/apitable")
	if err != nil {
		panic(err)
	}
	// create a new user repository
	userRepository := repository.NewUserRepository(db)
	// create a new user controller
	userController := app.NewAppUserController(userRepository)
	// create a new user router
	userRouter := chi.NewRouter()
	// add middleware to the user router
	userRouter.Use(middleware.RequestID)
	userRouter.Use(middleware.RealIP)
	userRouter.Use(middleware.Logger)
	// add the user controller to the user router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/users", userController.FindAll)
	http.ListenAndServe(":3000", r)
}
