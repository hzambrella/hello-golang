package memory

import (
	"container/list"
	"fmt"
	"session"
	"sync"
	"time"
)

var pder = &Provider{list: list.New()}

type SessionStore struct {
	sid         string
	timeAccessed time.Time
	value       map[interface{}]interface{}
}

type Provider struct {
	lock sync.Mutex
	sessions map[string]*list.Element //use as save seg
	list     *list.List               // use as gc,what is gc
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	fmt.Println("memory.go:GET:st.sid",st.sid)
	fmt.Println("memory.go:GET:st.timeAccessed",st.timeAccessed)
	fmt.Println("memory.go:GET:st.value",st.value)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
	return nil
}

func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}

func (provider *Provider) SessionInit(sid string) (session.Session, error) {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := provider.list.PushBack(newsess)
	provider.sessions[sid] = element
	//provider.sessions[sid]=newsess
	return newsess, nil
}

func (provider *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := provider.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := provider.SessionInit(sid)
		return sess, err
	}
	//	return nil,nil
}

func (provider *Provider) SessionDestory(sid string) error {
	if element,ok := provider.sessions[sid]; ok {
		delete(provider.sessions, sid)
		provider.list.Remove(element)
		return nil
	} else {
		return fmt.Errorf("SessionDestory:sid not found")
	}
}

func (provider *Provider) SessionGC(maxLifeTime int64) {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	for {
		element := provider.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLifeTime) < time.Now().Unix() {
			provider.list.Remove(element)
			delete(provider.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (provider *Provider) SessionUpdate(sid string) error {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	if element, ok := provider.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		provider.list.MoveToFront(element)
		return nil
	}
	return nil
}


func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	session.Registor(pder,"memory")
	fmt.Println("memory is doing")
}
