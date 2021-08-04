package repository

import (
	"api/models"
	"database/sql"
)

// CategoryRepository - implements the Repository interface for Category
type CategoryRepository struct {
	conn *sql.DB
}

func (r *CategoryRepository) FindAll() []models.Category {
	var categories []models.Category
	r.conn.QueryRow("SELECT * FROM categories").Scan(&categories)
	return categories
}

func (r *CategoryRepository) Find(id int) models.Category {
	var category models.Category
	r.conn.QueryRow("SELECT * FROM categories WHERE id = ?", id).Scan(&category)
	return category
}

func (r *CategoryRepository) Create(category models.Category) error {
	_, err := r.conn.Exec("INSERT INTO categories (name, description) VALUES (?, ?)", category.Name, category.Description)
	return err
}

func (r *CategoryRepository) Update(category models.Category) error {
	_, err := r.conn.Exec("UPDATE categories SET name = ?, description = ? WHERE id = ?", category.Name, category.Description, category.ID)
	return err
}

func (r *CategoryRepository) Delete(id int) error {
	_, err := r.conn.Exec("DELETE FROM categories WHERE id = ?", id)
	return err
}
