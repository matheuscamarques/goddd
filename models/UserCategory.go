package models

// UserCategory - is related to user category many-to-many relationship

type UserCategory struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	CategoryID int64  `json:"category_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
