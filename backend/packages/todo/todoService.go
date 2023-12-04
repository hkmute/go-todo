package todo

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoService struct {
	db *pgxpool.Pool
}

func NewTodoService(db *pgxpool.Pool) TodoService {
	return TodoService{db}
}

func (service TodoService) GetTodoList(params GetTodoListParams) ([]TodoEntity, error) {
	var todoList []TodoEntity

	limit := params.Limit
	if limit == 0 {
		limit = 10
	}

	rows, err := service.db.Query(context.Background(), "SELECT * FROM todo LIMIT $1 OFFSET $2", limit, params.Offset)
	defer rows.Close()

	if err != nil {
		return todoList, err
	}

	for rows.Next() {
		var todo TodoEntity
		rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.Created_at, &todo.Updated_at, &todo.User_id)
		todoList = append(todoList, todo)
	}
	return todoList, err
}

func (service TodoService) GetTodoCount() (int, error) {
	var count int
	err := service.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM todo").Scan(&count)
	return count, err
}

func (service TodoService) GetTodoById(id int) (TodoEntity, error) {
	var todo TodoEntity
	err := service.db.QueryRow(context.Background(), "SELECT * FROM todo WHERE id = $1", id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.Created_at, &todo.Updated_at, &todo.User_id)
	return todo, err
}

func (service TodoService) InsertTodo(newTodo NewTodo) (TodoEntity, error) {
	var todo TodoEntity
	err := service.db.QueryRow(context.Background(),
		"INSERT INTO todo (title, description, status, user_id) VALUES ($1, $2, $3, $4) RETURNING *",
		newTodo.Title, newTodo.Description, newTodo.Status, newTodo.UserId).Scan(&todo.Id, &todo.Title,
		&todo.Description, &todo.Status, &todo.Created_at, &todo.Updated_at, &todo.User_id)

	return todo, err
}

func (service TodoService) EditTodoById(id int, editTodo EditTodo) (TodoEntity, error) {
	var todo TodoEntity

	sql := "UPDATE todo SET"
	args := []interface{}{}
	i := 1

	v := reflect.ValueOf(editTodo)

	for j := 0; j < v.NumField(); j++ {
		field := v.Type().Field(j).Name
		value := v.Field(j)

		if value.Kind() == reflect.Ptr && !value.IsNil() {
			sql += fmt.Sprintf(" %s = $%d,", strings.ToLower(field), i)
			args = append(args, value.Elem().Interface())
			i++
		}
		if value.Kind() == reflect.String && value.String() != "" {
			sql += fmt.Sprintf(" %s = $%d,", strings.ToLower(field), i)
			args = append(args, value.Interface())
			i++
		}
	}

	sql += " updated_at = CURRENT_TIMESTAMP" + fmt.Sprintf(" WHERE id = $%d RETURNING *", i)
	args = append(args, id)

	err := service.db.QueryRow(context.Background(), sql, args...).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.Created_at, &todo.Updated_at, &todo.User_id)

	return todo, err
}

func (service TodoService) DeleteTodoById(id int) bool {
	commandTag, err := service.db.Exec(context.Background(), "DELETE FROM todo WHERE id = $1", id)
	if commandTag.RowsAffected() == 0 || err != nil {
		return false
	}
	return true
}
