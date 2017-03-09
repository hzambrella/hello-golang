package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
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

// different user have different session,distinguish by sid 
type Provider interface {
	SessionInit(sid string) (Session, error)// init
	SessionRead(sid string) (Session, error)// read
	SessionDestory(sid string) error // destory
	SessionGC(maxLifeTime int64)  // Gabarge destory
}

// session 
type Session interface {
	Set(key, value interface{}) error//  add modify
	Get(key interface{}) interface{}// view
	Delete(key interface{}) error// delete 
	SessionID() string  // aquire sid
}

// provider map [provider name]Provider
var provides = make(map[string]Provider)

func NewSessionManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session:unknown provider %q",provideName)
	}
	return &Manager{provider:provider,cookieName:cookieName,maxLifeTime:maxLifeTime}, nil
}

func Registor(provide Provider, name string) {
	if provide == nil {
		panic("Registion:no provide")
	}
	if _, dup := provides[name]; dup {
		panic("Registion,provide is exist")
	}
	provides[name] = provide
}

// create sid
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// start deal with cookie and session
// cookie is in the http.cookie.
// return session
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session){
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w,&cookie)
		fmt.Println("session.go:SessonStart:new session and cookie are created")
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
		fmt.Println("session.go:SessonStart:new session and cookie are readed")
	}
	return session
}

func (manager *Manager) SessionDestory(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		log.Fatal("session.go:SessionDestory:SessionDestory:value of cookie is nil")
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestory(cookie.Value)
		expiresion := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiresion, MaxAge: -1}
		http.SetCookie(w, &cookie)
		fmt.Println("session.go:SessionDestory")
	}
}

// Garbage collection
func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() { manager.GC() })
}
