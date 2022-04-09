package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var provides = make(map[string]Provider)

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider // Manager的能力存储在这个接口中，核心，可以认为这就是session本体
	maxLifeTime int64
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// 核心方法，将这一堆接口、结构体与http请求场景结合
// 从cookie中读取session的唯一标识（sessionId），若无则新建，若有，从注册的providers中拿出
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" { // 无历史session，新建
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

// 挂在Manager上的SessionDestroy，实际上使用的Provider的SessionDestroy
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration}
		http.SetCookie(w, &cookie)
	}
}

// 挂在Manager上的GC，实际上使用Provider的SessionGC
func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() { manager.provider.SessionGC(manager.maxLifeTime) })
}

// Session管理器功能接口
type Provider interface {
	SessionInit(sid string) (Session, error) // 初始化
	SessionRead(sid string) (Session, error) // 读，若不存在，新建
	SessionDestroy(sid string) error         // 摧毁
	SessionGC(maxLifeTime int64)             // GC
}

// Session的功能接口
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

// 在全局manager中注册一个命名Provider的方法
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknow provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}
