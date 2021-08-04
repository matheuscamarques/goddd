package app

import (
	"api/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type AppUserController struct {
	// The repository for the domain object
	userRepository *repository.UserRepository
}

func NewAppUserController(userRepository *repository.UserRepository) *AppUserController {
	return &AppUserController{userRepository}
}

func (c *AppUserController) Find(w http.ResponseWriter, r *http.Request) {
	// parse request body to get the id
	id := chi.URLParam(r, "id")
	// get the user from the repository
	// convert id to int
	idInt, _ := strconv.Atoi(id)

	user, err := c.userRepository.Find(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.userRepository.Close()
	w.Write(b)
}

func (c *AppUserController) FindAll(w http.ResponseWriter, r *http.Request) {

	users := c.userRepository.FindAll()

	b, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.userRepository.Close()
	w.Write(b)
}
