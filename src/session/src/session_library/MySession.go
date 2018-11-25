package geekori_session

import (
	"fmt"
	"time"
	"container/list"
	"sync"
	"encoding/base64"
	"math/rand"
	"net/http"
	"net/url"
)

var pder = &Provider{list: list.New()}
func GetSession() *Provider  {
	return pder
}
type SessionStore struct {
	sid          string                      //session id唯一标示
	timeAccessed time.Time                   //最后访问时间
	value        map[interface{}]interface{} //session里面存储的值
}

type Provider struct {
	lock     sync.Mutex               //用来锁
	sessions map[string]*list.Element //用来存储在内存
	list     *list.List               //用来做gc
}

func (this *SessionStore) Set(key, value interface{}) error {
	this.value[key] = value
	pder.SessionUpdate(this.sid)
	return nil
}
func (this *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(this.sid)
	if v, ok := this.value[key]; ok {
		return v
	} else {
		return nil
	}
}

func (this *SessionStore) Delete(key interface{}) error {
	delete(this.value, key)
	pder.SessionUpdate(this.sid)
	return nil
}



func (pder *Provider) SessionInit(sid string) (*SessionStore, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

func (pder *Provider) SessionRead(sid string) (*SessionStore, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

func (pder *Provider) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (this *Provider) SessionUpdate(sid string) error {
	this.lock.Lock()
	defer this.lock.Unlock()
	if element, ok := this.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		this.list.MoveToFront(element)
		return nil
	}
	return nil
}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	Register("memory", pder)
}

var providers = make(map[string]*Provider)


type Manager struct {
	cookieName  string     // Cookie名称
	lock        sync.Mutex // 同步锁
	provider    *Provider
	maxLifeTime int64
}

type Session interface {
	Set(key, value interface{}) error // set session value
	Get(key interface{}) interface{}  // get session value
	Delete(key interface{}) error     // delete session value
	SessionID() string                // back current sessionID
}

func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := providers[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}

func Register(name string, provider *Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := providers[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	providers[name] = provider
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	//  第一次访问服务端，还没有生成Session
	if err != nil || cookie.Value == "" {
		//  获取Session ID
		sid := manager.sessionId()
		//  初始化Session
		session, _ = manager.provider.SessionInit(sid)
		//  创建Cookie对象，并将session id保存到Cookie对象中
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		//  向客户端写入包含Session ID的
		http.SetCookie(w, &cookie)
	} else {  //  Session已经找到
        //  获取Session ID
		sid, _ := url.QueryUnescape(cookie.Value)
		//  找到Session对象
		session, _ = manager.provider.SessionRead(sid)

	}
	return
}
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}
func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() { manager.GC() })
}