package routes

import (
	"encoding/json"
	"engine/datastore"
	"errors"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

const COOKIE_MAX_MAX_AGE = time.Hour * 1 / time.Second // 单位：秒。
var (
	cookie string = "user_login"
	maxAge        = int(COOKIE_MAX_MAX_AGE)
)

type userInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"user_name"`
}

var userNotFound = errors.New("user not found in session")

type session map[string][]byte

func NewSession() session {
	sess := make(session, 0)
	return sess
}

var userSession session = NewSession()

func (s session) GetUser(key string) (*userInfo, error) {
	ub, ok := s[key]
	if !ok {
		return nil, userNotFound
	}
	user := &userInfo{}
	err := json.Unmarshal(ub, &user)
	if err != nil {
		return nil, err
	}

	s.Delete(key)
	return user, nil

}

func (s session) SetUser(key string, user *userInfo) error {
	ub, err := json.Marshal(user)
	if err != nil {
		return err
	}
	s[key] = ub
	return nil

}

var reqUrlStore datastore.Data = make(datastore.Data, 0)

func (s session) SetUser(key string, user *userInfo) error {
	ub, err := json.Marshal(user)
	if err != nil {
		return err
	}
	s[key] = ub
	return nil

}
func (s session) Delete(key string) {
	delete(s, key)
}

func (u *userInfo) setCookie(w http.ResponseWriter) {
	uuid := uuid.NewV4().String()
	userSession.SetUser(uuid, u)

	uid_cookie := &http.Cookie{
		Name:     cookie,
		Value:    uuid,
		Path:     "/",
		HttpOnly: false,
		MaxAge:   maxAge,
	}

	http.SetCookie(w, uid_cookie)
}

func auth(w http.ResponseWriter, r *http.Request) (*userInfo, bool) {
	ck, err := r.Cookie(cookie)
	if err != nil {
		logl.Error(err)
		return nil, false
	}

	u, err := userSession.GetUser(ck.Value)
	if err != nil {
		if err == userNotFound {
			login(w, r)
		} else {
			logl.Error(err)
			return nil, false
		}
	}

	u.setCookie(w)

	return u, true
}
