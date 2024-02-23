package todo

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoService struct {
	db *pgxpool.Pool
}

func NewTodoService(db *pgxpool.Pool) TodoService {
	return TodoService{db}
}

func (service TodoService) GetTodoList(params GetTodoListParams, userId int) ([]TodoEntity, error) {
	var todoList []TodoEntity

	var rows pgx.Rows
	var err error
	if params.Status == "" {
		rows, err = service.db.Query(context.Background(), "SELECT * FROM todo WHERE user_id = $1 ORDER BY item_order", userId)
	} else {
		rows, err = service.db.Query(context.Background(), "SELECT * FROM todo WHERE user_id = $1 AND status = $2 ORDER BY item_order", userId, params.Status)
	}

	defer rows.Close()

	if err != nil {
		return todoList, err
	}

	for rows.Next() {
		var todo TodoEntity
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.Created_at, &todo.Updated_at, &todo.User_id, &todo.Item_order)
		if err != nil {
			return todoList, err
		}
		todoList = append(todoList, todo)
	}

	return todoList, err
}

func (service TodoService) GetTodoCount(userId int) (int, error) {
	var count int
	err := service.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM todo where user_id = $1", userId).Scan(&count)
	return count, err
}

func (service TodoService) GetTodoById(id int) (TodoEntity, error) {
	var todo TodoEntity
	err := service.db.QueryRow(context.Background(), "SELECT * FROM todo WHERE id = $1", id).
		Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.Created_at, &todo.Updated_at, &todo.User_id, &todo.Item_order)
	return todo, err
}

func (service TodoService) InsertTodo(newTodo NewTodo) (TodoEntity, error) {
	var todo TodoEntity
	err := service.db.QueryRow(context.Background(),
		"WITH LatestOrder AS ("+
			"SELECT COALESCE(MAX(item_order), 0) AS latest_order "+
			"FROM todo "+
			"WHERE user_id = $1 AND status = $5) "+
			"INSERT INTO todo (title, description, status, user_id, item_order) "+
			"SELECT $2, $3, $4, $1, latest_order + 1 "+
			"FROM LatestOrder "+
			"RETURNING *",
		newTodo.UserId, newTodo.Title, newTodo.Description, newTodo.Status, newTodo.Status).Scan(&todo.Id, &todo.Title,
		&todo.Description, &todo.Status, &todo.Created_at, &todo.Updated_at, &todo.User_id, &todo.Item_order)

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

func (service TodoService) ReorderTodoById(id int, reorderDto ReorderTodo) (TodoEntity, error) {
	// Initialize a TodoEntity variable to store the current todo item
	var todo TodoEntity

	// Reorder todo item and update status
	sql := "WITH UpdatedOrder AS (" +
		"UPDATE todo " +
		"SET item_order = $1, status = $2 " +
		"WHERE id = $3 " +
		"RETURNING * " +
		"), " +
		"UpdatedItems AS (" +
		"UPDATE todo AS t " +
		"SET item_order = t.item_order + 1 " +
		"FROM UpdatedOrder " +
		"WHERE t.user_id = UpdatedOrder.user_id " +
		"AND t.status = UpdatedOrder.status " +
		"AND t.item_order >= UpdatedOrder.item_order " +
		"RETURNING t.*" +
		") " +
		"SELECT * FROM UpdatedOrder;"

	// Execute the SQL query
	err := service.db.QueryRow(context.Background(), sql,
		reorderDto.ItemOrder, reorderDto.Status, id).
		Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.Created_at, &todo.Updated_at, &todo.User_id, &todo.Item_order)

	return todo, err
}

func (service TodoService) DeleteTodoById(id int) bool {
	commandTag, err := service.db.Exec(context.Background(), "DELETE FROM todo WHERE id = $1", id)
	if commandTag.RowsAffected() == 0 || err != nil {
		return false
	}
	return true
}
