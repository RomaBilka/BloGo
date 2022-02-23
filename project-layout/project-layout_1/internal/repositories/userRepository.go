package repositories

import (
	"database/sql"

	"github.com/RomaBiliak/BloGo/project-layout/project-layou_1/internal/models"
)

func NewPostgreUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) GetUserById(id models.UserId) (models.User, error) {
	user := models.User{}

	err := r.db.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil && err != sql.ErrNoRows {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}

	err := r.db.QueryRow("SELECT * FROM users WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil && err != sql.ErrNoRows {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(id models.UserId, user models.User) error {
	err := r.db.QueryRow("UPDATE users SET name = $1, email = $2  WHERE id=$3", user.Name, user.Email, id).Err()

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(id models.UserId) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r *UserRepository) CreateUser(user models.User) (models.UserId, error) {
	id := 0
	err := r.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2)  RETURNING id", user.Name, user.Email).Scan(&id)
	if err != nil {
		return 0, err
	}

	return models.UserId(id), nil
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User

	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.User{}
		if err = rows.Scan(&user.Id, &user.Name, &user.Email); err != nil && err != sql.ErrNoRows {
			return []models.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}
