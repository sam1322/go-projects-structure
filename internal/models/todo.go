package models

import (
	"database/sql"
	"strings"
	"time"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Task      string `json:"task"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ExpiredAt string `json:"expired_at"`
}

type TodoRequest struct {
	// ID        string `json:"id"`
	Title string `json:"title"`
	Task  string `json:"task"`
	Done  bool   `json:"done"`
	// CreatedAt string `json:"created_at"`
	// UpdatedAt string `json:"updated_at"`
	ExpiredAt string `json:"expired_at"`
}

type TodoModel struct {
	DB *sql.DB
}

func (m *TodoModel) Insert(title string, task string, expires int) (*Todo, error) {
	query := `
			INSERT INTO todos (title,task,done,created_at,updated_at,expired_at) 
			VALUES($1,$2,$3,$4,$5,$6)
			RETURNING id, title,task,done,created_at,updated_at,expired_at`

	now := time.Now()
	expirationDate := now.AddDate(0, 0, expires)

	// Create an args slice containing the values for the placeholder parameters from
	// the movie struct. Declaring this slice immediately next to our SQL query helps to
	// make it nice and clear *what values are being used where* in the query.
	args := []interface{}{title, task, false, now, now, expirationDate}
	todo := &Todo{}
	rows := m.DB.QueryRow(query, args...)
	error := rows.Scan(&todo.ID, &todo.Title, &todo.Task, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt, &todo.ExpiredAt)
	return todo, error
}

func (m *TodoModel) Update(todoId string, title string, task string, expires int) (*Todo, error) {
	query := `
			UPDATE todos
			set title = $1, task = $2, done = $3, updated_at = $4, expired_at = $5
			WHERE id = $6
			RETURNING id,title,task,done,created_at,updated_at,expired_at`

	now := time.Now()
	expirationDate := now.AddDate(0, 0, expires)

	args := []interface{}{title, task, true, now, expirationDate, todoId}
	todo := &Todo{}
	rows := m.DB.QueryRow(query, args...)
	error := rows.Scan(&todo.ID, &todo.Title, &todo.Task, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt, &todo.ExpiredAt)
	return todo, error
}

func (m *TodoModel) UpdateCheck(todoId string, done bool) (*Todo, error) {
	query := `
			UPDATE todos
			set done = $1 , updated_at = $2
			WHERE id = $3
			RETURNING id,title,task,done,created_at,updated_at,expired_at`

	now := time.Now()
	// expirationDate := now.AddDate(0, 0, expires)

	args := []interface{}{done, now, todoId}
	todo := &Todo{}
	rows := m.DB.QueryRow(query, args...)
	error := rows.Scan(&todo.ID, &todo.Title, &todo.Task, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt, &todo.ExpiredAt)
	return todo, error
}

func (m *TodoModel) DeleteById(todoId string) (int, error) {
	todoId = strings.TrimSpace(todoId)
	query := `
			DELETE FROM todos
			WHERE id = $1;
		`

	res, err := m.DB.Exec(query, todoId)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(count), err
}

// This will return a specific snippet based on its id.
func (m *TodoModel) GetTodoById(todoId string) (*Todo, error) {
	query := `
			SELECT id, title, task, done, created_at, updated_at, expired_at FROM todos 
			WHERE id = $1`
	todo := &Todo{}
	error := m.DB.QueryRow(query, todoId).Scan(&todo.ID, &todo.Title, &todo.Task, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt, &todo.ExpiredAt)
	return todo, error

}

// This will return all todos
func (m *TodoModel) GetAllTodos() ([]*Todo, error) {
	query := `	
			SELECT id, title, task, done, created_at, updated_at, expired_at 
			FROM todos`

	rows, error := m.DB.Query(query)
	if error != nil {
		return nil, error
	}
	defer rows.Close()
	todos := []*Todo{}

	for rows.Next() {
		todo := &Todo{}
		error := rows.Scan(&todo.ID, &todo.Title, &todo.Task, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt, &todo.ExpiredAt)
		if error != nil {
			return nil, error
		}
		todos = append(todos, todo)
	}
	// When the rows.Next() loop has finished we call rows.Err() to retrieve any
	// error that was encountered during the iteration. It's important to
	// call this - don't assume that a successful iteration was completed
	// over the whole resultset.
	if err := rows.Err(); err != nil {
		return nil, err
	}
	// If everything went OK then return the todos slice
	return todos, nil

}

// This will return the 10 most recently created snippets.
func (m *TodoModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
