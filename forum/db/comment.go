package db

import (
	"db-test/models"
	"fmt"
	"log"
	"time"
)

func CreateCommentTable() error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        userid INTEGER NOT NULL,
        postid INTEGER NOT NULL,
        content TEXT NOT NULL,
        created_at TEXT,
        FOREIGN KEY (userid) REFERENCES users(id),
        FOREIGN KEY (postid) REFERENCES posts(id)
    );`

	_, err := Database.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating comments table: %v", err)
	}

	return nil
}

func CreateComment(c models.Comment) (int, error) {
	if _, err := ReadUser(c.UserID); err != nil {
		return -1, err
	}

	if _, err := ReadPost(c.PostID); err != nil {
		return -1, err
	}
	a := time.Now().Format("2006-01-02 15:04:05")
	o, err := Database.Exec(
		"INSERT INTO comments (userid, postid, content, created_at) VALUES (?, ?, ?, ?)",
		c.UserID,
		c.PostID,
		c.Content,
		a,
	)
	if err != nil {
		return -1, err
	}

	id, err := o.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil

}

func ReadComment(id int) (*models.Comment, error) {
	comment := &models.Comment{}
	rows, err := Database.Query(
		`SELECT id, userid, postid, content, created_at FROM comments WHERE id = ?`,
		id,
	)
	if err != nil {
		return comment, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var createdAt string
	if !rows.Next() {
		return nil, models.ErrNoResultFound
	}

	err = rows.Scan(
		&comment.ID,
		&comment.UserID,
		&comment.PostID,
		&comment.Content,
		&createdAt,
	)
	if err != nil {
		return comment, fmt.Errorf("error reading comment: %v", err)
	}

	comment.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	return comment, nil
}

func ReadAllComments() ([]models.Comment, error) {
	rows, err := Database.Query(`SELECT * FROM comments`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []models.Comment{}

	for rows.Next() {
		c := models.Comment{}
		var createdAt string

		err := rows.Scan(&c.ID, &c.UserID, &c.PostID, &c.Content, &createdAt)
		if err != nil {
			log.Printf("error reading c: %v", err)
			continue
		}

		c.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			log.Printf("error parsing created_at for comment ID %d: %v", c.ID, err)
			continue
		}

		comments = append(comments, c)
	}
	if len(comments) == 0 {
		return nil, models.ErrNoResultFound
	}

	return comments, nil
}

func UpdateComment(c models.Comment) error {
	currentTime := time.Now()

	if _, err := ReadUser(c.UserID); err != nil {
		return err
	}

	if _, err := ReadPost(c.PostID); err != nil {
		return err
	}
	if _, err := ReadComment(c.ID); err != nil {
		return err
	}

	statement, err := Database.Prepare(
		`UPDATE comments
        SET userid = ?, postid = ?, content = ?, created_at = ?
        WHERE id = ?`,
	)
	if err != nil {
		return err
	}

	_, err = statement.Exec(
		c.UserID,
		c.PostID,
		c.Content,
		currentTime.Format("2006-01-02 15:04:05"),
		c.ID,
	)
	if err != nil {
		return err
	}

	return err
}

func DeleteComment(id int) error {
	if _, err := ReadComment(id); err != nil {
		return err
	}
	statement, err := Database.Prepare(
		`DELETE FROM comments
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

func ReadCommentsByFunc(f func(models.Comment) bool) ([]models.Comment, error) {
	rows, err := Database.Query(`SELECT * FROM comments`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []models.Comment{}

	for rows.Next() {
		c := models.Comment{}
		var createdAt string

		err := rows.Scan(&c.ID, &c.UserID, &c.PostID, &c.Content, &createdAt)
		if err != nil {
			log.Printf("error reading c: %v", err)
			continue
		}

		c.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			log.Printf("error parsing created_at for comment ID %d: %v", c.ID, err)
			continue
		}

		if f(c) {
			comments = append(comments, c)
		}
	}
	if len(comments) == 0 {
		return nil, models.ErrNoResultFound
	}

	return comments, nil
}
