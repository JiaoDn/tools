package sessions

import (
	"container/list"
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	maxLifeTime int64
}

func GetManager(cookieName string, maxLifeTime int64) (*Manager, error) {
	return &Manager{cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

type SessionStore struct {
	sessionId       string                      // session id唯一标示
	recentlyVisited time.Time                   // 最后访问时间
	value           map[interface{}]interface{} // session里面存储的值
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value

	return nil
}

func (st *SessionStore) Get(key interface{}) interface{} {

	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
}

func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)

	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sessionId
}

type LocalStore struct {
	lock     sync.Mutex               // 用来锁
	sessions map[string]*list.Element // 用来存储在内存
	list     *list.List               // 用来做 gc
}

func (ls *LocalStore) SessionInit(sessionId string) (*SessionStore, error) {
	ls.lock.Lock()
	defer ls.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sessionId: sessionId, recentlyVisited: time.Now(), value: v}
	element := ls.list.PushFront(newsess)
	ls.sessions[sessionId] = element
	return newsess, nil
}

func (ls *LocalStore) SessionRead(sessionId string) (*SessionStore, error) {
	if element, ok := ls.sessions[sessionId]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := ls.SessionInit(sessionId)
		return sess, err
	}
}
