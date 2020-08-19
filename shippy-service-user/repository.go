package main

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	pb "github.com/wkcw/shippy/shippy-service-user/proto/user"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Company  string
	Password string
}

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

func MarshalUserCollection(users []*pb.User) []*User {
	u := make([]*User, len(users))
	for _, val := range users {
		u = append(u, MarshalUser(val))
	}
	return u
}

func MarshalUser(user *pb.User) *User {
	return &User{
		ID:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Company:  user.Company,
		Password: user.Password,
	}
}

func UnmarshalUserCollection(users []*User) []*pb.User {
	u := make([]*pb.User, len(users))
	for _, val := range users {
		u = append(u, UnmarshalUser(val))
	}
	return u
}

func UnmarshalUser(user *User) *pb.User {
	return &pb.User{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Company:  user.Company,
		Password: user.Password,
	}
}

func (r *PostgresRepository) GetAll(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	if err := r.db.GetContext(ctx, users, "select * from users"); err != nil {
		return users, err
	}

	return users, nil
}

func (r *PostgresRepository) Get(ctx context.Context, id string) (*User, error) {
	var user *User
	if err := r.db.GetContext(ctx, &user, "select * from users where id = &1", id); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresRepository) Create(ctx context.Context, user *User) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	user.ID = id.String()
	query := "insert into users (id, name, email, company, password) values ($1, $2, $3, $4, $5)"
	_, err = r.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Company, user.Password)
	return err
}

func (r *PostgresRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := "select * from users where email = $1"
	var user *User
	if err := r.db.GetContext(ctx, &user, query, email); err != nil {
		return nil, err
	}
	return user, nil
}
