package db

import (
	"db-test/models"
	"fmt"
)

func CreateCataPostTable() error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS Cata_post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        Cata_ID INTEGER NOT NULL,
        Post_ID INTEGER NOT NULL,
        FOREIGN KEY (Post_ID) REFERENCES posts(id),
        FOREIGN KEY (Cata_ID) REFERENCES category(id)
    );`

	_, err := Database.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating category table: %v", err)
	}

	return nil
}

func CreateCataPost(pId, cataId int) error {
	if _, err := ReadPost(pId); err != nil {
		return err
	}
	if _, err := ReadCategory(cataId); err != nil {
		return err
	}

	_, err := Database.Exec(
		"INSERT INTO Cata_post (cata_id, post_id) VALUES (?, ?)",
		cataId,
		pId,
	)
	return err
}

func ReadCataPost() ([]models.PostCategory, error) {
	rows, err := Database.Query("SELECT * FROM Cata_post")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cataPosts := []models.PostCategory{}
	for rows.Next() {
		var cp models.PostCategory
		err := rows.Scan(&cp.ID, &cp.CategoryID, &cp.PostID)
		if err != nil {
			return nil, err
		}
		cataPosts = append(cataPosts, cp)
	}
	return cataPosts, nil
}

func ReadAllCataPost() ([]models.PostCategory, error) {
	rows, err := Database.Query("SELECT * FROM Cata_post")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cataPosts := []models.PostCategory{}
	for rows.Next() {
		var cp models.PostCategory
		err := rows.Scan(&cp.ID, &cp.CategoryID, &cp.PostID)
		if err != nil {
			return nil, err
		}
		cataPosts = append(cataPosts, cp)
	}
	return cataPosts, nil
}

func UpdateCataPost(pId, cataId int) error {

	if _, err := ReadPost(pId); err != nil {
		return err
	}
	if _, err := ReadCategory(cataId); err != nil {
		return err
	}
	stmt, err := Database.Prepare("UPDATE Cata_post SET cata_id = ? WHERE post_id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cataId, pId)
	return err
}
func DeleteCataPost(pId, cataId int) error {
	if _, err := ReadPost(pId); err != nil {
		return err
	}
	if _, err := ReadCategory(cataId); err != nil {
		return err
	}
	_, err := Database.Exec("DELETE FROM Cata_post WHERE post_id = ? AND cata_id = ?", pId, cataId)
	if err != nil {
		return err
	}
	return nil
}
