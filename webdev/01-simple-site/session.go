package main

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	// get cookie:
	c, err := r.Cookie("session")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			log.Fatalln(err)
		}
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure:   true,
			HttpOnly: true,
		}
		//http.SetCookie(w, c)
	}

	// if user exists, get user:
	var u user
	if username, ok := dbSessions[c.Value]; ok {
		u = dbUsers[username]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
