package memory
import(
"fmt"
"container/list"
"session"
"sync"
"time"
)

type sessionStore struct {
	sid string
	timeAcessed time.Time
	value map[interface{}]interface{}
}

type provider struct{
	lock *sync.Mutex

	sessions map[string]*list.Element //use as save seg
	list *list.List // use as gc,what is gc
}

func (provider *Provider)SessionInit(sid string)(session.Session,error){
	provider.lock.Lock()
	defer provider.lock.Unlock()
	v:=make(map[interface{}]interface{},0)
	newsess:=&SessionStore{sid:sid,timeAccessed:time.Now(),value:v}
	element:=provider.list.PushBack(newsess)
	provider.sessions[sid]=element
	//provider.sessions[sid]=newsess
	return newsess,nil
}

func (provider *Provide)SessionRead(sid string)(session.Session,error){
	if element,ok:=provider.session[sid];ok{
		return element.Value.(*SessionStore),nil
	}else{
		sess,err:=provider.SessionInit(sid)
		return sess,err
	}
//	return nil,nil
}

func(provider *Provide)SessionDestory(sid string)(error){
	if element:=provider.sessions[sid];ok{
		delete(provider.sessions,sid)
		return nil
	}else{
		return fmt.Errorf("SessionDestory:sid not found")
	}
}

func (provider *Provider)SessionGC(maxLifeTime int64){
}
