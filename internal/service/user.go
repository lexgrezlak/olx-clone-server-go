package service

func (api *API) GetUserByEmail(email string) (*User, error) {
	row := api.Db.QueryRow(
		"SELECT * FROM user WHERE email=$1 LIMIT 1", email)

	var u *User
	if err := row.Scan(&u); err != nil {
		return nil, err
	}

	return u, nil
}