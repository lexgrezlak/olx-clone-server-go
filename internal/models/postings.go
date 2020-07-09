package models

import "context"

type Posting struct {
	Title string
}

func AllPostings() ([]*Posting, error) {
	rows, err := db.Pool.Query(context.Background(), "SELECT title FROM posting")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	postings := make([]*Posting, 0)
	for rows.Next() {
		posting := new(Posting)
		err := rows.Scan(&posting.Title)
		if err != nil {
			return nil, err
		}
		postings = append(postings, posting)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return postings, nil
}
