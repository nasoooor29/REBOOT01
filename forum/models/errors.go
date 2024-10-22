package models

import "fmt"

var (
	ErrCookieNotFound      = fmt.Errorf("cookie not found")
	ErrMustProvideCategory = fmt.Errorf("must provide a category")
	ErrNoResultFound       = fmt.Errorf("no results found")
	ErrNoPostFound         = fmt.Errorf("no posts found")
	ErrCommentNotFound     = fmt.Errorf("Comment not found")
)
