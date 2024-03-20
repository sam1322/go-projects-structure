package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"snippetbox.sam1322/internal/models"
)

func (app *Application) getAllTodos(w http.ResponseWriter, _ *http.Request) {
	todos, err := app.models.TodoModel.GetAllTodos()
	if err != nil {
		app.serverError(w, err)
		return
	}

	var response = make(map[string]interface{})
	response["todos"] = todos

	todosJSON, err := json.Marshal(response)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(todosJSON)

	// return map[string]interface{}{
	// 	"message": "It's healthy",
	// 	"users":   todos,
	// }
	// app.writeJSON(w, http.StatusOK, todos)

}

func (app *Application) getTodoById(w http.ResponseWriter, r *http.Request) {

	todoIdParam := chi.URLParam(r, "id")
	todos, err := app.models.TodoModel.GetTodoById(todoIdParam)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// var response = make(map[string]interface{})
	// response["todos"] = todos
	// todosJSON, err := json.Marshal(response)

	todosJSON, err := json.Marshal(todos)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(todosJSON)

}

func (app *Application) createTodo(w http.ResponseWriter, r *http.Request) {

	// todoIdParam := chi.URLParam(r, "id")
	todoReq := &models.TodoRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todoReq)

	// if todoReq.Task == nil {
	// 	todoReq.Task = ""
	// }

	if err != nil {
		app.serverError(w, err)
		return
	}

	todo, err := app.models.TodoModel.Insert(todoReq.Title, todoReq.Task, 7)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// // var response = make(map[string]interface{})
	// // response["todos"] = todos
	// // todosJSON, err := json.Marshal(response)

	todosJSON, err := json.Marshal(todo)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(todosJSON)

}

func (app *Application) editTodo(w http.ResponseWriter, r *http.Request) {

	// todoIdParam := chi.URLParam(r, "id")
	todo := &models.Todo{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)

	if err != nil {
		app.serverError(w, err)
		return
	}

	todo, err = app.models.TodoModel.Update(todo.ID, todo.Title, todo.Task, 7)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// // var response = make(map[string]interface{})
	// // response["todos"] = todos
	// // todosJSON, err := json.Marshal(response)

	todosJSON, err := json.Marshal(todo)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(todosJSON)

}

func (app *Application) deleteTodo(w http.ResponseWriter, r *http.Request) {

	todoIdParam := chi.URLParam(r, "id")
	// todo := &models.Todo{}

	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&todo)

	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	count, err := app.models.TodoModel.DeleteById(todoIdParam)

	if err != nil {
		app.serverError(w, err)
		return
	}
	var message string = "Deleted"

	if count > 0 {
		message = "Task deleted successfully"
	} else {
		message = "Task not found"
	}
	var response = make(map[string]interface{})
	response["message"] = message
	todosJSON, err := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(todosJSON)

}

func (app *Application) updateTodoCheck(w http.ResponseWriter, r *http.Request) {
	todoIdParam := chi.URLParam(r, "id")
	todoReq := &models.TodoRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todoReq)

	if err != nil {
		app.serverError(w, err)
		return
	}

	todo, err := app.models.TodoModel.UpdateCheck(todoIdParam, todoReq.Done)
	if err != nil {
		app.serverError(w, err)
		return
	}

	todosJSON, err := json.Marshal(todo)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(todosJSON)
}
