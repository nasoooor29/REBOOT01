package utils

import (
	"db-test/db"
	"db-test/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateCookie(w http.ResponseWriter, userId int) error {
	cookieValue, err := uuid.NewV4()
	if err != nil {
		return err
	}
	_, err = db.CreateCookie(userId, cookieValue.String())
	if err != nil {
		err = db.DeleteCookieByUserID(userId)
		if err != models.ErrCookieNotFound && err != nil {
			fmt.Printf("err: %v\n", err)
			return err
		}
		fmt.Println("deleted the old cookie creating new one")
		db.CreateCookie(userId, cookieValue.String())
	}

	http.SetCookie(w, &http.Cookie{
		Name:     models.AUTH_COOKIE_TITLE,
		Value:    cookieValue.String(),
		Path:     "/",
		MaxAge:   int(time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	return nil
}

func CheckIfAuth(r *http.Request) (*models.Cookie, error) {
	reqCookie, err := r.Cookie(models.AUTH_COOKIE_TITLE)
	if err != nil {
		return nil, err
	}
	cookie, err := db.ReadCookieByFunc(func(c models.Cookie) bool {
		return c.Cookie == reqCookie.Value
	})
	if err != nil {
		return nil, err
	}
	return cookie, nil
}
