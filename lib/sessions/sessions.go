package sessions

import (
	"strings"
	"time"

	set "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type SessionID string

type Session struct {
	createdAt     time.Time
	lastUsageAt   time.Time
	subscriptions set.Set[string]
}

var sessionsMap map[SessionID]*Session

func init() {
	sessionsMap = make(map[SessionID]*Session)
	go loadBackup()
	go storeBackup()
}

func loadBackup() {
}

func storeBackup() {
	for {
		time.Sleep(5 * time.Minute)
		func() {
		}()
	}
}

func isSessionIDValid(sid string) bool {
	_, err := uuid.FromBytes([]byte(sid))
	return err == nil
}

func createSessionId() SessionID {
	return SessionID(
		strings.Replace(
			uuid.New().String(), "-", "", -1),
	)
}

func GetSession(uuidStr string) (*Session, SessionID) {
	var uuid SessionID
	if isSessionIDValid(uuidStr) {
		uuid = SessionID(uuidStr)
	} else {
		uuid = createSessionId()
	}
	session, ok := sessionsMap[uuid]
	if !ok {
		session = &Session{
			createdAt:     time.Now(),
			lastUsageAt:   time.Now(),
			subscriptions: set.NewSet[string](),
		}
		sessionsMap[uuid] = session
	}
	return session, uuid
}

func (s *Session) HeartBeat() {
	s.lastUsageAt = time.Now()
}

func (s *Session) Get() *[]string {
	subscriptionsList := make([]string, s.subscriptions.Cardinality())
	return &subscriptionsList
}

func (s *Session) Subscribe(topics *[]string) {
	for _, t := range *topics {
		s.subscriptions.Add(t)
	}
}

func (s *Session) UnSubscribe(topics *[]string) {
	for _, t := range *topics {
		s.subscriptions.Remove(t)
	}
}
