package models

import "time"

var STATUS_TEMPLATES = map[int]string{
	400: `templates/pages/error.html`,
	404: `templates/pages/error.html`,
	500: `templates/pages/error.html`,
	200: `templates/pages/error.html`,
}

type User struct {
	ID       int // primary key
	Username string
	Email    string
	Password string
}

type Cookie struct {
	ID     int // primary key
	UserID int // foreign key
	Cookie string
	// NASER: add created at
}

type Post struct {
	ID        int // primary key
	Title     string
	Content   string
	UserID    int // foreign key
	CreatedAt time.Time

	Categories []Category
}

type Comment struct {
	ID        int // primary key
	UserID    int // foreign key
	PostID    int // foreign key
	Content   string
	CreatedAt time.Time
}


// upvote downvote
type PostInteraction struct {
	ID     int // primary key
	UserID int // foreign key
	PostID int // foreign key
	Vote   Vote
}

// like disklike
type CommentInteraction struct {
	ID          int // primary key
	UserID      int // foreign key
	CommentID   int // foreign key
	Interaction Interaction
}

type Category struct {
	ID          int // primary key
	UserID      int // foreign key
	Title       string
	Description string
}

type PostCategory struct {
	ID         int // primary key
	PostID     int // foreign key
	CategoryID int // foreign key

}
