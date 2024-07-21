package session

import (
	"github.com/gorilla/sessions"
	"net/http"
)

func SessionAuthMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		store := sessions.NewCookieStore([]byte("your-secret-key"))
		session, _ := store.Get(request, "session-name")

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(rw, "Forbidden", http.StatusForbidden)
			return
		}

		handlerFunc.ServeHTTP(rw, request)
	}
}
