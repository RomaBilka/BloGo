package user

import (
	"database/sql"
)

func NewPostgreUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) GetUserById(id UserId) (User, error) {
	user := User{}

	err := r.db.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil && err != sql.ErrNoRows {
		return User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (User, error) {
	user := User{}

	err := r.db.QueryRow("SELECT * FROM users WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil && err != sql.ErrNoRows {
		return User{}, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(id UserId, user User) error {
	err := r.db.QueryRow("UPDATE users SET name = $1, email = $2  WHERE id=$3", user.Name, user.Email, id).Err()

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(id UserId) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r *UserRepository) CreateUser(user User) (UserId, error) {
	id := 0
	err := r.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2)  RETURNING id", user.Name, user.Email).Scan(&id)
	if err != nil {
		return 0, err
	}

	return UserId(id), nil
}

func (r *UserRepository) GetUsers() ([]User, error) {
	var users []User

	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Name, &user.Email); err != nil && err != sql.ErrNoRows {
			return []User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}
