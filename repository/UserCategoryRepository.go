package repository

import (
	"api/models"
	"database/sql"
)

// UserCategoryRepository - implements RepositoryInterface
type UserCategoryRepository struct {
	conn *sql.DB
}

func (r *UserCategoryRepository) FindAll() []models.UserCategory {
	var userCategories []models.UserCategory
	r.conn.QueryRow(`SELECT * FROM user_category`).Scan(userCategories)
	return userCategories
}

func (r *UserCategoryRepository) Find(id int) models.UserCategory {
	var userCategory models.UserCategory
	r.conn.QueryRow(`SELECT * FROM user_category WHERE id = $1`, id).Scan(userCategory)
	return userCategory
}

func (r *UserCategoryRepository) Create(userCategory models.UserCategory) error {
	// insert into UserCategory (ID,UserID,CategoryID,CreatedAt,UpdatedAt) values (1,1,1,now(),now())
	_, err := r.conn.Exec(`INSERT INTO user_category (id,user_id,category_id,created_at,updated_at) VALUES ($1,$2,$3,now(),now())`, userCategory.ID, userCategory.UserID, userCategory.CategoryID)
	return err
}

func (r *UserCategoryRepository) Update(userCategory models.UserCategory) error {
	_, err := r.conn.Exec(`UPDATE user_category SET category_id = $1, updated_at = now() WHERE id = $2`, userCategory.CategoryID, userCategory.ID)
	return err
}

func (r *UserCategoryRepository) Delete(id int) error {
	_, err := r.conn.Exec(`DELETE FROM user_category WHERE id = $1`, id)
	return err
}
