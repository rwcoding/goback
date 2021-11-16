package session

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/util"
	"strings"
	"time"
)

var db = models.GetDb()

func NewKVSession(v string, sessionId string) *models.Session {
	if sessionId != "" {
		var session models.Session
		db.Where("session_id=?", sessionId).Take(&session)
		if session.Id > 0 {
			session.SessionValue = v
			session.Expire = uint32(time.Now().Unix()) + 1800
			db.Save(&session)
			return &session
		}
	}

	k := util.RandString(32)
	s := &models.Session{
		Type:         models.SessionTypeCaptcha,
		SessionId:    k,
		SessionValue: v,
		Expire:       uint32(time.Now().Unix()) + 1800,
	}
	if db.Create(s).Error != nil {
		return nil
	}
	return s
}

func NewAuthSession(adminerId uint32) *models.Session {
	sessionId := util.RandString(32)
	sessionValue := util.RandString(32)
	s := &models.Session{
		AdminerId:    adminerId,
		Type:         models.SessionTypeAuth,
		SessionId:    sessionId,
		SessionValue: sessionValue,
		Expire:       uint32(time.Now().Unix()) + 86400*30,
	}
	if db.Create(s).Error != nil {
		return nil
	}
	return s
}

func VerifySession(sessionId, v string) bool {
	s := &models.Session{}
	db.Where("session_id=?", sessionId).Take(s)
	if s.Id == 0 {
		return false
	}
	return strings.ToLower(s.SessionValue) == strings.ToLower(v)
}

func QuerySession(sessionId string) *models.Session {
	s := &models.Session{}
	if db.Where("session_id=?", sessionId).Take(s).Error != nil {
		return nil
	}
	return s
}
