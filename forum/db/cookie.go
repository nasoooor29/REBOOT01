package db

import (
	"db-test/models"
	"fmt"
	"log"
)

func CreateCookieTable() error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS cookie (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        userid INTEGER NOT NULL UNIQUE,
        cookie TEXT NOT NULL UNIQUE,
        FOREIGN KEY (userid) REFERENCES users(id)
        );`
	_, err := Database.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating cookie table: %v", err)
	}
	return nil
}

func CreateCookie(userID int, cookie string) (int, error) {
	o, err := Database.Exec(
		"INSERT INTO cookie (userid, cookie) VALUES (?, ?)",
		userID,
		cookie,
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

func ReadCookie(id int) (*models.Cookie, error) {
	c := &models.Cookie{}
	rows, err := Database.Query(`SELECT id, userid, cookie FROM cookie WHERE id = ?`, id)
	if err != nil {
		return c, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, models.ErrNoResultFound
	}

	err = rows.Scan(&c.ID, &c.UserID, &c.Cookie)
	if err != nil {
		return c, fmt.Errorf("error reading c: %v", err)
	}
	return c, nil
}

func ReadAllCookies() ([]models.Cookie, error) {
	rows, err := Database.Query(`SELECT * FROM cookie`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cookies := []models.Cookie{}

	for rows.Next() {
		c := models.Cookie{}
		err := rows.Scan(&c.ID, &c.UserID, &c.Cookie)
		if err != nil {
			log.Printf("error reading c: %v", err)
			continue
		}
		cookies = append(cookies, c)
	}
	if len(cookies) == 0 {
		return nil, models.ErrNoResultFound
	}
	return cookies, nil
}

func UpdateCookie(userID, cookieID int, cookie string) error {
	statement, err := Database.Prepare(
		`UPDATE cookie
        SET userid = ?, cookie = ?
        WHERE id = ?`,
	)
	if err != nil {
		return err
	}
	_, err = statement.Exec(userID, cookie, cookieID)
	if err != nil {
		return err
	}

	return err
}

func DeleteCookie(id int) error {
	if _, err := ReadCookie(id); err != nil {
		return err
	}

	statement, err := Database.Prepare(
		`DELETE FROM cookie
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

func DeleteCookieByUserID(uid int) error {
	cookies, err := ReadAllCookies()
	if err != nil {
		return err
	}

	for _, cookie := range cookies {
		if cookie.UserID == uid {
			DeleteCookie(cookie.ID)
			return nil
		}
	}
	return models.ErrCookieNotFound
}

func ReadCookieByFunc(f func(c models.Cookie) bool) (*models.Cookie, error) {
	cookies, err := ReadAllCookies()
	if err != nil {
		return nil, err
	}

	for _, cookie := range cookies {
		if f(cookie) {
			return &cookie, nil
		}
	}
	return nil, models.ErrCookieNotFound
}
