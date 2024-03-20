package models

import "database/sql"

type Models struct {
	TodoModel *TodoModel
}

func NewModels(db *sql.DB) *Models {
	return &Models{
		TodoModel: &TodoModel{DB: db},
	}
}
