package models

const (
	DB_NAME      = "./data.db"
	DEFAULT_PORT = ":8080"
	AUTH_COOKIE_TITLE = "AuthCookie"
)

type Vote int

// post
const (
	UP_VOTE   Vote = 0
	DOWN_VOTE Vote = 1
	NO_VOTE   Vote = 2
)

type Interaction int
// comment
const (
	LIKE    Interaction = 0
	DISLIKE Interaction = 1
	NEUTRAL Interaction = 2
)
