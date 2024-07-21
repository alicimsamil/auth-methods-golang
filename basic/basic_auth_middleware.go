package basic

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

// BasicAuthMiddleWare is an HTTP middleware function. This middleware checks the validity
// of Basic Auth credentials in incoming HTTP requests. If the username and password are correct,
// the request is forwarded to the specified handler function. Otherwise, it responds with a
// 401 Unauthorized status and prompts the client for authentication.
func BasicAuthMiddleWare(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		name, password, isOk := request.BasicAuth()
		if isOk {
			// Hash the username and password using SHA-256.
			usernameHash := sha256.Sum256([]byte(name))
			passwordHash := sha256.Sum256([]byte(password))

			// Predefined expected hashes for the username and password.
			expectedUsernameHash := sha256.Sum256([]byte("ali"))
			expectedPasswordHash := sha256.Sum256([]byte("kucuk"))

			// Compare the hashed username and password using constant-time comparison.
			// Constant-time comparison ensures that the time taken to compare the two hashed values is the same,
			// regardless of where the first difference between them is found. This is crucial for security purposes
			// because it prevents attackers from gaining information about the actual values based on how long the
			// comparison takes.
			usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
			passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

			if usernameMatch && passwordMatch {
				handlerFunc.ServeHTTP(writer, request)
				return
			}
		}

		writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		writer.WriteHeader(http.StatusUnauthorized)
	}
}
