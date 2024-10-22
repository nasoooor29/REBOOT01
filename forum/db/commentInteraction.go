package db

import (
	"db-test/models"
	"fmt"
	"log"
)

func CreateCommentInteractionTable() error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS comment_interactions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        userid INTEGER NOT NULL,
        commentid INTEGER NOT NULL,
        interaction INTEGER NOT NULL,
        FOREIGN KEY (userid) REFERENCES users(id),
        FOREIGN KEY (commentid) REFERENCES comment(id)
        );`
	_, err := Database.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating comment_interactions table: %v", err)
	}

	return nil
}

func CreateCommentInteraction(interaction models.Interaction, userID, commentID int) (int, error) {
	if _, err := ReadUser(userID); err != nil {
		return -1, err
	}
	if _, err := ReadComment(commentID); err != nil {
		return -1, err
	}

	var check int
	Database.QueryRow("SELECT COUNT(*) FROM comments WHERE id = ?", commentID).Scan(&check)
	if check == 0 {
		return -1, fmt.Errorf("Id not present")
	}

	o, err := Database.Exec(
		"INSERT INTO comment_interactions (userid, commentid, interaction) VALUES (?, ?, ?)",
		userID,
		commentID,
		interaction,
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

// TODO: need to test this
func ReadCommentInteraction(id int) (*models.CommentInteraction, error) {
	ci := &models.CommentInteraction{}
	rows, err := Database.Query(
		`SELECT id, userid, commentid, interaction FROM comment_interactions WHERE id = ?`,
		id,
	)
	if err != nil {
		return ci, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var interactionInt int
	if !rows.Next() {
		return nil, models.ErrNoResultFound
	}
	err = rows.Scan(
		&ci.ID,
		&ci.UserID,
		&ci.CommentID,
		&interactionInt,
	)
	if err != nil {
		return ci, fmt.Errorf("error reading comment interaction: %v", err)
	}

	ci.Interaction = models.Interaction(interactionInt)
	return ci, nil
}

// TODO: need to test this
func ReadAllCommentInteractions() ([]models.CommentInteraction, error) {
	rows, err := Database.Query(`SELECT * FROM comment_interactions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	interactions := []models.CommentInteraction{}

	for rows.Next() {
		ci := models.CommentInteraction{}
		var interactionInt int
		err := rows.Scan(&ci.ID, &ci.UserID, &ci.CommentID, &interactionInt)
		if err != nil {
			log.Printf("error reading ci: %v", err)
			continue
		}
		ci.Interaction = models.Interaction(interactionInt)
		interactions = append(interactions, ci)
	}
	return interactions, nil
}

func UpdateCommentInteraction(interaction models.Interaction, userID, commentID, ciID int) error {
	if _, err := ReadUser(userID); err != nil {
		return err
	}
	if _, err := ReadComment(commentID); err != nil {
		return err
	}

	if _, err := ReadCommentInteraction(ciID); err != nil {
		return err
	}

	statement, err := Database.Prepare(
		`UPDATE comment_interactions
        SET userid = ?, commentid = ?, interaction = ?
        WHERE id = ?`,
	)
	if err != nil {
		return err
	}
	_, err = statement.Exec(userID, commentID, interaction, ciID)
	if err != nil {
		return err
	}

	return err
}

func DeleteCommentInteraction(id int) error {

	if _, err := ReadCommentInteraction(id); err != nil {
		return err
	}

	statement, err := Database.Prepare(
		`DELETE FROM comment_interactions
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
