package repository

// import models.User

import (
	"api/models"
	"database/sql"
)

//make a implements models.User RepositoryInterface interface
type UserRepository struct {
	// conn use sqlx
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{conn}
}

func (r *UserRepository) Close() {
	r.conn.Close()
}

func (r *UserRepository) Find(id int) (models.User, error) {
	user := models.User{}
	err := r.conn.QueryRow(`SELECT * FROM users WHERE id = ?`, id).Scan(&user.ID, &user.Name)
	return user, err
}

func (r *UserRepository) FindAll() []models.User {
	var users []models.User
	r.conn.QueryRow(`SELECT * FROM users`).Scan(&users)
	return users
}

func (r *UserRepository) Create(user models.User) error {
	_, err := r.conn.Exec(`INSERT INTO users (name) VALUES (?)`, user.Name)
	return err
}

func (r *UserRepository) Update(user models.User) error {
	_, err := r.conn.Exec(`UPDATE users SET name = ? WHERE id = ?`, user.Name, user.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.conn.Exec(`DELETE FROM users WHERE id = ?`, id)
	return err
}

// make functions Exists by Name
func (r *UserRepository) Exists(user models.User) bool {
	var count int
	r.conn.QueryRow(`SELECT COUNT(*) FROM users WHERE name = ?`, user.Name).Scan(&count)
	return count > 0
}
