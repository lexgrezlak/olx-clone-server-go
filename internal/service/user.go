package service

import "context"

func (db *DB) GetUserByEmail(email string) (*User, error) {
	row := db.Pool.QueryRow(context.Background(),
		"SELECT * FROM user WHERE email=$1 LIMIT 1", email)

	var u *User
	if err := row.Scan(&u); err != nil {
		return nil, err
	}

	return u, nil
}