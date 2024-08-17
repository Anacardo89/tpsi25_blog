package auth

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/Anacardo89/tpsi25_blog/internal/handlers/data/orm"
	"github.com/Anacardo89/tpsi25_blog/internal/model/database"
	"github.com/Anacardo89/tpsi25_blog/internal/model/mapper"
	"github.com/Anacardo89/tpsi25_blog/internal/model/presentation"
	"github.com/Anacardo89/tpsi25_blog/pkg/logger"
	"github.com/gorilla/sessions"
)

type Config struct {
	Pass string `yaml:"session_pass"`
}

var (
	SessionStore *sessions.CookieStore
)

func CreateSession(w http.ResponseWriter, r *http.Request) presentation.Session {
	usrSession := presentation.Session{}
	session, err := SessionStore.Get(r, "tpsi25blog")
	if err != nil {
		logger.Error.Println(err)
	}
	newSID := generateSessionId()
	session.Values["sid"] = newSID
	err = session.Save(r, w)
	if err != nil {
		logger.Error.Println(err)
	}
	usrSession.SessionId = newSID
	usrSession.Authenticated = true
	return usrSession
}

func ValidateSession(r *http.Request) presentation.Session {
	usrSession := presentation.Session{}
	session, err := SessionStore.Get(r, "tpsi25blog")
	if err != nil {
		logger.Error.Println(err)
	}
	if sid, valid := session.Values["sid"]; valid {
		dbuser := orm.GetUserBySessionID(sid.(string))
		u := mapper.User(&dbuser)
		usrSession.User = *u
		UpdateSession(sid.(string), usrSession.User.Id)
		usrSession.SessionId = sid.(string)
		usrSession.Authenticated = true
	} else {
		usrSession.Authenticated = false
	}
	return usrSession
}

func UpdateSession(sid string, uid int) {
	s := &database.Session{
		SessionId: sid,
		UserId:    uid,
		Active:    1,
	}
	if err := orm.Da.CreateSession(s); err != nil {
		logger.Error.Println(err)
	}
}

func generateSessionId() string {
	sid := make([]byte, 24)
	_, err := io.ReadFull(rand.Reader, sid)
	if err != nil {
		logger.Error.Println(err)
	}
	return base64.URLEncoding.EncodeToString(sid)
}