package handlers

import (
	"db-test/db"
	"db-test/utils"
	//"fmt"
	"net/http"
)

func HandleLogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := utils.CheckIfAuth(r)
	if err != nil {
	//	fmt.Println("VAT U DO??")
		return
	}
	db.DeleteCookieByUserID(cookie.UserID)
	http.Redirect(w, r, "/signIn", http.StatusSeeOther)

}
