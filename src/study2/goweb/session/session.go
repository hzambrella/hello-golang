package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestory(sid string) error
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

var provides = make(map[string]Provider)

func NewSessionManager(provideName, cookieName String, MaxLifeTime int64) (*Manager, error) {
	provide, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session:unknown provider")
	}
	return &Manager{cookieName: cookieName, MaxLifeTime: MaxLifeTime, provider: provide}, nil
}

func Registion(provide Provider, name String) {
	if Provide == nil {
		Panic("Registion:no provide")
	}
	if _, dup := Provide[name]; dup {
		Panic("Registion,provide is exist ,can't Registion")
	}
	provides[name] = provide
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manger.cookieName, Value: url.QueryEscape(sid), Path: '/', HttpOnly: true, Expire: int(Manager.MaxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnEscape(cookie.Value)
		session, _ := manager.provider.SessionRead(sid)
	}
	return session
}

func (manager *Manager) SessionDestory(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(Manager.cookieName)
	if err != nil || cookie.Value == "" {
		log.Fatal("SessionDestory:value of cookie is nil")
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestory(cookie.Value)
		expiresion := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiresion, MaxAge: -1}
		http.SetCookie(w, cookie)
	}
}

func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() { manager.GC })
}
