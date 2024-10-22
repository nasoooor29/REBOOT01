package db

import (
	"db-test/models"
	"fmt"
	"log"
	"time"
)

func CreatePostTable() error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        userid TEXT NOT NULL,
        created_at TEXT,
        FOREIGN KEY (userid) REFERENCES users(id)
        );`
	_, err := Database.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating posts table: %v", err)
	}
	return nil
}

func CreatePost(p models.Post) error {
	if _, err := ReadUser(p.UserID); err != nil {
		return err
	}

	currentTime := time.Now()
	if len(p.Categories) == 0 {
		return models.ErrMustProvideCategory
	}

	postObj, err := Database.Exec(
		"INSERT INTO posts (title, content, userid, created_at) VALUES (?, ?, ?, ?)",
		p.Title,
		p.Content,
		p.UserID,
		currentTime.Format("2006-01-02 15:04:05"),
	)

	pid, err := postObj.LastInsertId()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	for _, cata := range p.Categories {
		err := CreateCataPost(int(pid), cata.ID)
		if err != nil {
			return err
		}
	}

	return err
}

func ReadPost(id int) (*models.Post, error) {
	post := &models.Post{}
	rows, err := Database.Query(
		`SELECT title, content, userid, created_at FROM posts WHERE id = ?`,
		id,
	)
	if err != nil {
		return post, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, models.ErrNoResultFound
	}

	t := ""

	err = rows.Scan(&post.Title, &post.Content, &post.UserID, &t)
	if err != nil {
		return post, fmt.Errorf("error reading post: %v", err)
	}
	post.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", t)

	return post, nil
}

func ReadAllPosts() ([]models.Post, error) {
	rows, err := Database.Query(`SELECT * FROM posts`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []models.Post{}

	for rows.Next() {
		p := models.Post{}
		var createdAt string

		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.UserID, &createdAt)
		if err != nil {
			log.Printf("error reading p: %v", err)
			continue
		}

		p.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			log.Printf("error parsing created_at for post ID %d: %v", p.ID, err)
			continue
		}

		posts = append(posts, p)
	}
	if len(posts) == 0 {
		return nil, models.ErrNoResultFound
	}
	return posts, nil
}

func UpdatePost(p models.Post) error {
	if _, err := ReadUser(p.ID); err != nil {
		return err
	}
	if _, err := ReadPost(p.UserID); err != nil {
		return err
	}

	currentTime := time.Now()
	// if len(p.Categories) == 0 {
	// 	return models.ErrMustProvideCategory
	// }
	statement, err := Database.Prepare(
		`UPDATE posts
        SET title = ?, content = ?, userid = ?, created_at = ?
        WHERE id = ?`,
	)
	if err != nil {
		return err
	}

	_, err = statement.Exec(
		p.Title,
		p.Content,
		p.UserID,
		currentTime.Format("2006-01-02 15:04:05"),
		p.ID,
	)
	if err != nil {
		return err
	}

	return err
}

func DeletePost(id int) error {
	if _, err := ReadPost(id); err != nil {
		return err
	}

	statement, err := Database.Prepare(
		`DELETE FROM posts
        WHERE id = ?`,
	)
	if err != nil {
		return err
	}

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return err
}

func ReadPostByFunc(f func(c models.Post) bool) ([]models.Post, error) {
	posts, err := ReadAllPosts()
	if err != nil {
		return nil, err
	}
	result := []models.Post{}
	for _, v := range posts {
		wanted := f(v)
		if wanted {
			result = append(result, v)
		}
	}
	return result, nil
}
