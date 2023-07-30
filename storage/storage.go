package storage

import (
	"database/sql"
	"fmt"

	"github.com/koha90/project-driver/internal/users"
	_ "github.com/lib/pq"
)

type Storage interface {
	GetAccounts() ([]*users.User, error)
	GetAccountByID(id int) (*users.User, error)
	CreateAccount(*users.User) (id int, err error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(
	user string,
	database string,
	password string,
) (*PostgresStorage, error) {
	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=disable",
		user,
		database,
		password,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) Init() error {
	return s.createUserTable()
}

func (s *PostgresStorage) createUserTable() error {
	query := `create table if not exists users (
        id serial primary key,
        username varchar(50),
        encrypted_password varchar(100),
        created_at timestamp
    )`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateUser(u *users.User) error {
	query := `INSERT INTO users
    (username, encrypted_password, created_at)
    values ($1, $2, $3)`

	_, err := s.db.Query(query, u.Username, u.EncryptedPassword, u.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) DeleteUser(id int) error {
	_, err := s.db.Query("DELETE FROM users where id = $1", id)
	return err
}

func (s *PostgresStorage) GetUsers() ([]*users.User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	users := []*users.User{}
	for rows.Next() {
		user, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func scanIntoAccount(rows *sql.Rows) (*users.User, error) {
	user := new(users.User)
	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.EncryptedPassword,
		&user.CreatedAt,
	)

	return user, err
}
