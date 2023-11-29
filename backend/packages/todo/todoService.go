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

func (t TodoService) GetTodoList(params GetTodoListParams) ([]todoEntity, error) {
	var todoList []todoEntity

	limit := params.Limit
	if limit == 0 {
		limit = 10
	}

	rows, err := t.db.Query(context.Background(), "SELECT * FROM todo LIMIT $1 OFFSET $2", limit, params.Offset)
	defer rows.Close()

	if err != nil {
		return todoList, err
	}

	for rows.Next() {
		var todo todoEntity
		rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.Created_at, &todo.Updated_at)
		todoList = append(todoList, todo)
	}
	return todoList, err
}

func (t TodoService) GetTodoCount() (int, error) {
	var count int
	err := t.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM todo").Scan(&count)
	return count, err
}

func (t TodoService) GetTodoById(id int) (todoEntity, error) {
	var todo todoEntity
	err := t.db.QueryRow(context.Background(), "SELECT * FROM todo WHERE id = $1", id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.Created_at, &todo.Updated_at)
	return todo, err
}

func (t TodoService) InsertTodo(newTodo NewTodo) (todoEntity, error) {
	var todo todoEntity
	err := t.db.QueryRow(context.Background(),
		"INSERT INTO todo (title, description, status) VALUES ($1, $2, $3) RETURNING *",
		newTodo.Title, newTodo.Description, newTodo.Status).Scan(&todo.Id, &todo.Title,
		&todo.Description, &todo.Status, &todo.Created_at, &todo.Updated_at)

	return todo, err
}

func (t TodoService) EditTodoById(id int, editTodo EditTodo) (todoEntity, error) {
	var todo todoEntity

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

	err := t.db.QueryRow(context.Background(), sql, args...).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.Created_at, &todo.Updated_at)

	return todo, err
}

func (t TodoService) DeleteTodoById(id int) bool {
	commandTag, err := t.db.Exec(context.Background(), "DELETE FROM todo WHERE id = $1", id)
	if commandTag.RowsAffected() == 0 || err != nil {
		return false
	}
	return true
}
