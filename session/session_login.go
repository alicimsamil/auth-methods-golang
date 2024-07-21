package session

import (
	"auth-methods/common"
	"github.com/gorilla/sessions"
	"net/http"
)

func UserLoginWithSession(rw http.ResponseWriter, r *http.Request) {
	store := sessions.NewCookieStore([]byte("your-secret-key"))
	session, _ := store.Get(r, "session-name")

	common.DecodeUser(r)
	// Perform login actions with user here.

	session.Values["authenticated"] = true

	if err := session.Save(r, rw); err != nil {
		return
	}

	http.Redirect(rw, r, "/", http.StatusFound)
}
