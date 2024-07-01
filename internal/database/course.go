package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Category    string
	Description string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name string, description string, categoryID string) (Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1,$2,$3,$4)", id, name, description, categoryID)
	if err != nil {
		return Course{}, err
	}
	return Course{ID: id, Name: name, Description: description, Category: categoryID}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id from courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []Course{}
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		categories = append(categories, Course{ID: id, Name: name, Description: description, Category: categoryID})
	}
	return categories, nil
}
