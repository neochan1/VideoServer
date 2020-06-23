package session

import (
	"VideoServer/apis/models"
	"VideoServer/apis/utils"
	"sync"
	"time"
	"VideoServer/apis/database"
)

var sessionMap *sync.Map

func init()  {
	sessionMap = &sync.Map{}
}
func nowInMillion() int64 {
	return time.Now().UnixNano()/1000000
}
func deleteExpiredSession(sid string)  {
	sessionMap.Delete(sid)
	database.DeleteSession(sid)
}
func LoadSessionsFromDB() {
	r, err := database.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{}) bool{
		ss := v.(*models.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowInMillion()
	ttl := ct + 30 * 60 * 1000// Severside session valid time: 30 min

	ss := &models.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	database.InsertSession(id, ttl, un)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMillion()
		if ss.(*models.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*models.SimpleSession).Username, false
	}

	return "", true
}
