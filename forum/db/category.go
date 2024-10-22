package db

import (
	"db-test/models"
	"fmt"
	"log"
)

func CreateCategoryTable() error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS category (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        userid INTEGER NOT NULL,
        title TEXT NOT NULL UNIQUE,
        description TEXT NOT NULL,
        FOREIGN KEY (userid) REFERENCES users(id)
    );`

	_, err := Database.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating category table: %v", err)
	}

	return nil
}

func CreateCategory(userID int, title, description string) (int, error) {
	o, err := Database.Exec(
		"INSERT INTO category (userid, title, description) VALUES (?, ?, ?)",
		userID,
		title,
		description,
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

func ReadCategory(id int) (*models.Category, error) {
	category := &models.Category{}
	rows, err := Database.Query(
		`SELECT id, userid, title, description FROM category WHERE id = ?`,
		id,
	)
	if err != nil {
		return category, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, models.ErrNoResultFound
	}
	err = rows.Scan(&category.ID, &category.UserID, &category.Title, &category.Description)
	if err != nil {
		return category, fmt.Errorf("error reading category: %v", err)
	}
	return category, nil
}

func ReadAllCategory() ([]models.Category, error) {
	rows, err := Database.Query(`SELECT * FROM category`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []models.Category{}

	for rows.Next() {
		c := models.Category{}
		err := rows.Scan(&c.ID, &c.UserID, &c.Title, &c.Description)
		if err != nil {
			log.Printf("error reading c: %v", err)
			continue
		}
		categories = append(categories, c)
	}
	if len(categories) == 0 {
		return nil, models.ErrNoResultFound
	}
	return categories, nil
}

func UpdateCategory(c models.Category) error {
	if _, err := ReadCategory(c.ID); err != nil {
		return err
	}
	if _, err := ReadUser(c.UserID); err != nil {
		return err
	}
	statement, err := Database.Prepare(
		`UPDATE category
        SET userid = ?, title = ?, description = ?
        WHERE id = ?`,
	)
	if err != nil {
		return err
	}

	_, err = statement.Exec(c.UserID, c.Title, c.Description, c.ID)
	if err != nil {
		return err
	}

	return err
}

func DeleteCategory(id int) error {
	if _, err := ReadCategory(id); err != nil {
		return err
	}
	statement, err := Database.Prepare(
		`DELETE FROM category
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
