package db

import (
	"db-test/models"
	"fmt"
	"log"
)

func CreatePostInteractionTable() error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS post_interactions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        userid INTEGER NOT NULL,
        postid INTEGER NOT NULL,
        vote INTEGER NOT NULL,
        FOREIGN KEY (userid) REFERENCES users(id),
        FOREIGN KEY (postid) REFERENCES posts(id)
        );`
	_, err := Database.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating post_interactions table: %v", err)
	}

	return nil
}

func CreatePostInteraction(vote models.Vote, userID, postID int) error {
	var check int 
	Database.QueryRow("SELECT COUNT(*) FROM posts WHERE id = ?", postID).Scan(&check)
	if check == 0 {
		return fmt.Errorf("ID Not there")
		
	}
	_, err := Database.Exec(
		"INSERT INTO post_interactions (userid, postid, vote) VALUES (?, ?, ?)",
		userID,
		postID,
		vote,
	)
	return err
}

func ReadPostInteraction(id int) (*models.PostInteraction, error) {
	interaction := &models.PostInteraction{}
	rows, err := Database.Query(
		`SELECT id, userid, postid, vote FROM post_interactions WHERE id = ?`,
		id,
	)
	if err != nil {
		return interaction, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, models.ErrNoResultFound
	}

	var voteInt int
	err = rows.Scan(
		&interaction.ID,
		&interaction.UserID,
		&interaction.PostID,
		&voteInt,
	)
	if err != nil {
		return interaction, fmt.Errorf("error reading interaction: %v", err)
	}
	interaction.Vote = models.Vote(voteInt)

	return interaction, nil
}

func ReadAllPostInteractions() ([]models.PostInteraction, error) {
	rows, err := Database.Query(`SELECT * FROM post_interactions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	interactions := []models.PostInteraction{}

	for rows.Next() {
		p := models.PostInteraction{}
		var voteInt int
		err := rows.Scan(&p.ID, &p.UserID, &p.PostID, &voteInt)
		if err != nil {
			log.Printf("error reading p: %v", err)
			continue
		}
		p.Vote = models.Vote(voteInt)
		interactions = append(interactions, p)
	}

	if len(interactions) == 0 {
		return nil, models.ErrNoResultFound
	}

	return interactions, nil
}

func UpdatePostInteraction(vote models.Vote, userID, postID, piID int) error {
	if _, err := ReadPostInteraction(piID); err != nil {
		return err
	}
	if _, err := ReadUser(userID); err != nil {
		return err
	}
	if _, err := ReadPost(postID); err != nil {
		return err
	}

	statement, err := Database.Prepare(
		`UPDATE post_interactions
        SET userid = ?, postid = ?, vote = ?
        WHERE id = ?`,
	)
	if err != nil {
		return err
	}

	_, err = statement.Exec(userID, postID, vote, piID)
	if err != nil {
		return err
	}

	return err
}

func DeletePostInteraction(id int) error {
	if _, err := ReadPostInteraction(id); err != nil {
		return err
	}

	statement, err := Database.Prepare(
		`DELETE FROM post_interactions
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

func ReadPostInteractionByFunc(f func(c models.PostInteraction) bool) (*models.PostInteraction, error) {
	postInteractions, err := ReadAllPostInteractions()
	if err != nil {
		return nil, err
	}

	for _, interaction := range postInteractions {
		if f(interaction) {
			return &interaction, nil
		}
	}
	return nil, models.ErrNoPostFound
}
