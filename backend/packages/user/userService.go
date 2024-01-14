package user

import (
	"context"
	"fmt"
	"go-todo/packages/util"
	"reflect"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) UserService {
	return UserService{db}
}

func (service UserService) GetUser(params GetUserParams) (AppUserEntity, error) {
	var user AppUserEntity

	sql := "SELECT id, username"
	if params.WithPassword == true {
		sql += ", password"
	}
	sql += " FROM app_user WHERE"

	args := []interface{}{}
	i := 1

	v := reflect.ValueOf(params)

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

	sql = strings.TrimSuffix(sql, ",")

	var err error
	if params.WithPassword == true {
		err = service.db.QueryRow(context.Background(), sql, args...).Scan(&user.Id, &user.Username, &user.Password)
	} else {
		err = service.db.QueryRow(context.Background(), sql, args...).Scan(&user.Id, &user.Username)
	}

	return user, err
}

func (service UserService) GetUserToken(params GetUserParams) (string, error) {
	token, err := util.GenerateJwt(*params.Id)

	return token, err
}

func (service UserService) CreateUser(newUser NewUser) (AppUserEntity, error) {
	var user AppUserEntity
	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		return user, err
	}

	err = service.db.QueryRow(context.Background(), "INSERT INTO app_user (username, password) VALUES ($1, $2) RETURNING id, username",
		newUser.Username, hashedPassword).Scan(&user.Id, &user.Username)

	return user, err
}

func (u UserService) EditUser(id int, editUser EditUser) (AppUserEntity, error) {
	var user AppUserEntity
	return user, nil
}

func (u UserService) DeleteUser(id int) bool {
	return true
}
