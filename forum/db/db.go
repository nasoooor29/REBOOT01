package db

import (
	"database/sql"
	"db-test/models"
	"fmt"
	"log"
)

var Database *sql.DB

func CreateTables() error {
	if err := CreateUsersTable(); err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	if err := CreateCookieTable(); err != nil {
		return fmt.Errorf("failed to create cookie table: %v", err)
	}

	if err := CreatePostTable(); err != nil {
		return fmt.Errorf("failed to create post table: %v", err)
	}

	if err := CreateCommentTable(); err != nil {
		return fmt.Errorf("failed to create comment table: %v", err)
	}

	if err := CreatePostInteractionTable(); err != nil {
		return fmt.Errorf("failed to create post interaction table: %v", err)
	}

	if err := CreateCommentInteractionTable(); err != nil {
		return fmt.Errorf("failed to create comment interaction table: %v", err)
	}

	if err := CreateCategoryTable(); err != nil {
		return fmt.Errorf("failed to create category table: %v", err)
	}
	if err := CreateCataPostTable(); err != nil {
		return fmt.Errorf("failed to create cata_post table: %v", err)
	}

	return nil
}

// NOTE: ONLY FOR TESTING
func DropTables(db *sql.DB, tables []string) error {
	for _, table := range tables {
		query := fmt.Sprintf("DROP TABLE IF EXISTS %s;", table)
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to drop table %s: %w", table, err)
		}
	}
	return nil
}

func UpdateTest() {
	err := UpdateUser("updatetestuser", "update@email.com", 1)
	if err != nil {
		log.Fatalf("could not update user: %v", err)
	}

	updatePost := models.Post{
		Title:   "updated-post",
		Content: "updated-content",
		UserID:  1,
		ID:      1,
	}
	err = UpdatePost(updatePost)
	if err != nil {
		log.Fatalf("could not update post: %v", err)
	}

	err = UpdateCookie(1, 1, "newcookie")
	if err != nil {
		log.Fatalf("could not update cookie: %v", err)
	}

	updateComment := models.Comment{
		UserID:  1,
		PostID:  1,
		Content: "new-content-stuff",
		ID:      1,
	}
	err = UpdateComment(updateComment)
	if err != nil {
		log.Fatalf("could not update comment: %v", err)
	}

	updateCategory := models.Category{
		UserID:      1,
		Title:       "new-awesome-title",
		Description: "new-awesome-description",
		ID:          1,
	}
	err = UpdateCategory(updateCategory)
	if err != nil {
		log.Fatalf("could not update category: %v", err)
	}

	err = UpdatePostInteraction(1, 1, 1, 1)
	if err != nil {
		log.Fatalf("could not update PI: %v", err)
	}

	err = UpdateCommentInteraction(1, 1, 1, 1)
	if err != nil {
		log.Fatalf("could not update PI: %v", err)
	}
}

func Init() error {
	d, err := sql.Open("sqlite3", models.DB_NAME)
	if err != nil {
		return err
	}
	Database = d

	if err := Database.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	err = CreateTables()
	if err != nil {
		return err
	}
	log.Println("created all tables successfully")

	return nil
}
