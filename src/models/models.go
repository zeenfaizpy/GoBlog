package models

import (
	"database/sql"
)

// Post ...
type Post struct {
	ID      int
	Title   string
	Content string
}

// GetAllPosts ...
func GetAllPosts(db *sql.DB) ([]*Post, error) {
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allPosts := make([]*Post, 0)
	for rows.Next() {
		post := new(Post)
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return allPosts, nil
}
