package sessions

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
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
