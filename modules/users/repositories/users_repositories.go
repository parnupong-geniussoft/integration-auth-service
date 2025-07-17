package repositories

import (
	"fmt"
	"go-fiber-clean-arch-example/modules/entities"

	"github.com/jmoiron/sqlx"
)

type UsersRepository interface {
	Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error)
}

type usersRepo struct {
	Db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) UsersRepository {
	return &usersRepo{
		Db: db,
	}
}

func (r *usersRepo) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	query := `
	INSERT INTO "users"(
		"username",
		"password"
	)
	VALUES ($1, $2)
	RETURNING "id", "username";
	`

	// Initail a user object
	user := new(entities.UsersRegisterRes)

	// Query part
	rows, err := r.Db.Queryx(query, req.Username, req.Password)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		if err := rows.StructScan(user); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	defer r.Db.Close()

	return user, nil
}
