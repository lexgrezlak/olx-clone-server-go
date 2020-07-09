package models

import "context"

type Posting struct {
	Title string
	Price int
}

func AllPostings() ([]*Posting, error) {
	rows, err := db.Pool.Query(context.Background(), "SELECT title, price FROM posting")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ps := make([]*Posting, 0)
	for rows.Next() {
		p := new(Posting)
		err := rows.Scan(&p.Title, &p.Price)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ps, nil
}
