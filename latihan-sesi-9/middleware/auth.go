package middleware

import (
	"encoding/json"
	"latihan-rest-api/models"
	"net/http"
)

func Auth(rw http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		outputJson(rw, "something went wrong")
		return false
	}

	isValid := false
	for _, user := range models.Users {
		isValid = (user.Username == username) && (user.Password == password)
		if isValid{
			break
		}
	}
	if !isValid {
		outputJson(rw, "username / password is wrong")
		return isValid
	}
	return isValid
}

func outputJson(rw http.ResponseWriter, payload interface{}) {
	response := map[string]interface{}{
		"error": payload,
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(401)
	json.NewEncoder(rw).Encode(response)
}
